import json
import os
from github import Github

GITHUB_TOKEN = os.environ.get("GITHUB_TOKEN")
REPO_NAME = os.environ.get("GITHUB_REPOSITORY")

g = Github(GITHUB_TOKEN)
repo = g.get_repo(REPO_NAME)

with open('semgrep-results.json', 'r') as f:
    results = json.load(f)

for finding in results['results']:
    if 'severity' in finding['extra'] and finding['extra']['severity'] == 'ERROR':
        title = f"Semgrep Error: {finding['check_id']}"
        body = f"""
Path: {finding['path']}
Line: {finding['start']['line']}
Message: {finding['extra']['message']}"""
    repo.create_issue(title=title, body=body)