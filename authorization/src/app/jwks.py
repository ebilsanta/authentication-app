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

def update_authZ_key():
    key = requests.get(get_settings().authz_pub_key_url).text
    Settings.authz_pub_key = key
    return key

def update_client_pub_key():
    key = requests.get(get_settings().allowed_client_pub_key_url).text
    Settings.allowed_client_pub_key = key
    return key