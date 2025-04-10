import requests

# Proxy configuration
use_proxy = False  # Set this to False to turn off the proxy
proxy = {
    "http": "http://127.0.0.1:8080",
    "https": "http://127.0.0.1:8080"
}

# Function to make HTTP requests
def fetch_url(url):
    try:
        if use_proxy:
            response = requests.get(url, proxies=proxy)
        else:
            response = requests.get(url)
        
        print(f"URL: {url} - Status Code: {response.status_code}")
    except requests.exceptions.RequestException as e:
        print(f"URL: {url} - Error: {e}")

# Reading URLs from the text file
def read_urls_from_file(file_name):
    try:
        with open(file_name, "r") as file:
            for line in file:
                url = line.strip()  # Remove any whitespace
                if url:  # Ensure the line is not empty
                    fetch_url(url)
    except FileNotFoundError:
        print(f"File not found: {file_name}")

# Main execution
if __name__ == "__main__":
    file_name = "urls.txt"  # Replace with your text file containing URLs
    read_urls_from_file(file_name)