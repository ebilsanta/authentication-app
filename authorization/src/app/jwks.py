import requests

from config import Settings
from functools import lru_cache

@lru_cache()
def get_settings():
    return Settings()

def update_authN_key():
    key = requests.get(get_settings().pub_key_url).text
    Settings.pub_key = key
    return key