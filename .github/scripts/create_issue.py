import json
import os
from github import Github

GITHUB_TOKEN = os.environ.get("GITHUB_TOKEN")
REPO_NAME = os.environ.get("GITHUB_REPOSITORY")

g = Github(GITHUB_TOKEN)
repo = g.get_repo(REPO_NAME)

with open('results.json', 'r') as f:
    results = json.load(f)


for finding in results['results']:
    if 'severity' in finding['extra'] and finding['extra']['is_ignored'] == False:
        title = f"Semgrep Error: {finding['check_id']}"
        body = f"""
Severity: {finding['extra']['severity']}
Path: {finding['path']}
Line: {finding['start']['line']}
Code: {finding['extra']['lines']}
Message: {finding['extra']['message']}"""
        repo.create_issue(title=title, body=body)