from functools import lru_cache
import os
from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    allowed_client_pub_key: str
    allowed_client_pvk_key: str
    allowed_client: str
    allowed_issuer: str
    allowed_redirect: str

    pub_key: str
    pvt_key: str

    audience: str

    db_name: str
    db_collection_authcodes: str
    db_collection_users: str
    db_region_name: str
    db_access_key_id: str
    db_secret_access_key: str
    
    dpop_htm: str
    dpop_htu: str

    model_config = SettingsConfigDict(env_file=os.getenv('env_file') if os.getenv('env_file') else '.env')

@lru_cache()
def get_settings():
    return Settings()