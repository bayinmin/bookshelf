import json

def validate_jwks(file_path):
    with open(file_path, 'r') as f:
        jwks = json.load(f)
    
    for key in jwks.get("keys", []):
        if "d" in key:  # Private key component for RSA/EC
            print("Private key exposed: ", key)
            return False
    print("No private keys detected. JWKS is safe.")
    return True

# Usage
validate_jwks("jwks.json")