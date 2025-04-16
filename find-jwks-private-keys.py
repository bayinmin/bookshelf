import json

def validate_jwks(file_path):
    with open(file_path, 'r') as f:
        jwks = json.load(f)
    
    for key in jwks.get("keys", []):
        sensitive_fields = ["d", "p", "q"]  # Private key components
        exposed_fields = [field for field in sensitive_fields if field in key]

        if exposed_fields:
            print(f"Private key components exposed: {exposed_fields}")
            print("Exposed Key: ", key)
            return False
    print("No private keys detected. JWKS is safe.")
    return True

# Usage
validate_jwks("jwks.json")