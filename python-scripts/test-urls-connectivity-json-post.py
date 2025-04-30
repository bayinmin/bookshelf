import requests
import json
import argparse
import urllib3

# Suppress warnings
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# Default proxy settings
DEFAULT_PROXY = "http://127.0.0.1:8080"
USE_PROXY = True  # Proxy enabled by default

# Dummy JSON post body
DUMMY_DATA = {
    "name": "Test User"
}

def send_post_request(url, use_proxy):
    """Send a POST request to the given URL."""
    proxies = {"http": DEFAULT_PROXY, "https": DEFAULT_PROXY} if use_proxy else None

    try:
        response = requests.post(url, json=DUMMY_DATA, proxies=proxies, verify=False)
        print(f"Sent POST to {url}, Status Code: {response.status_code}")
    except requests.exceptions.RequestException as e:
        print(f"Error sending request to {url}: {e}")

def main():
    parser = argparse.ArgumentParser(description="Python script to send POST requests from a file list of URLs.")
    parser.add_argument("file", help="File containing URLs (one per line)")
    parser.add_argument("--no-proxy", action="store_true", help="Disable proxy")

    args = parser.parse_args()
    use_proxy = not args.no_proxy  # Proxy is enabled unless --no-proxy is provided

    try:
        with open(args.file, "r") as f:
            urls = [line.strip() for line in f if line.strip()]
        
        for url in urls:
            send_post_request(url, use_proxy)
    
    except FileNotFoundError:
        print(f"Error: File '{args.file}' not found.")

if __name__ == "__main__":
    main()