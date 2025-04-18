import json
import sys

def validate_jwks(file_path):
    try:
        with open(file_path, 'r') as f:
            jwks = json.load(f)
            
            
        print("Finding d, p, q entries in the jwks file....")
        for key in jwks.get("keys", []):
            sensitive_fields = ["d", "p", "q"]  # Private key components
            exposed_fields = [field for field in sensitive_fields if field in key]

            if exposed_fields:
                print(f"Private key components exposed: {exposed_fields}")
                print("Exposed Key: ", key)
                return False
        print("No private keys detected. JWKS is safe.")
        return True
    except FileNotFoundError:
        print(f"Error: File '{file_path}' not found.")
    except json.JSONDecodeError:
        print(f"Error: File '{file_path}' is not a valid JSON file.")

# Usage: Pass the file name as a command-line argument
if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python validate_jwks.py <file_path>")
    else:
        validate_jwks(sys.argv[1])