import base64
import hashlib
import os
import random


def generate_pkce_code_verifier():
    length = random.randint(43, 128)
    random_bytes = os.urandom(length // 2)
    code_verifier = base64.urlsafe_b64encode(random_bytes).decode("ascii")[:length]
    return code_verifier


def generate_pkce_code_challenge(code_verifier):
    encoder = hashlib.sha256()
    encoder.update(bytes(code_verifier, "ascii"))
    hash = encoder.digest()

    code_challenge = base64.urlsafe_b64encode(hash).decode("ascii")
    return code_challenge
