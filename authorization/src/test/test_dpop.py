from app.dpop_proof import create_dpop_jwt, verify_dpop_jwt
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization

def test_create_dpop():
    private_key = rsa.generate_private_key(
        public_exponent=65537,
        key_size=2048,
    )
    pvk = private_key.private_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PrivateFormat.TraditionalOpenSSL,
        encryption_algorithm=serialization.NoEncryption()
    )

    public_key = private_key.public_key().public_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PublicFormat.SubjectPublicKeyInfo
    )

    dpop_jwt = create_dpop_jwt(pvk, public_key, "www.google.com")[1:-1]
    verify_dpop_jwt(dpop_jwt, "www.google.com", 'POST')