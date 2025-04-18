import requests
import time
import warnings
import csv
import argparse

# Suppress warnings
warnings.filterwarnings("ignore")

# Proxy configuration
proxy = {
    "http": "http://127.0.0.1:8080",
    "https": "http://127.0.0.1:8080"
}

# Function to make HTTP requests
def fetch_url(url, writer, use_proxy):
    try:
        if use_proxy:
            response = requests.get(url, proxies=proxy, verify=False)  # Ignore SSL errors with verify=False
        else:
            response = requests.get(url, verify=False)  # Ignore SSL errors with verify=False
        
        print(f"URL: {url} - Status Code: {response.status_code}")
        writer.writerow([url, response.status_code])  # Write to CSV file
    except requests.exceptions.RequestException as e:
        print(f"URL: {url} - Error: {e}")
        writer.writerow([url, "Error"])  # Write error to CSV file

# Reading URLs from the text file
def read_urls_from_file(file_name, output_file, use_proxy):
    try:
        with open(file_name, "r") as file, open(output_file, "w", newline="") as csvfile:
            writer = csv.writer(csvfile)
            writer.writerow(["URL", "Status Code"])  # Write header
            for line in file:
                url = line.strip()  # Remove any whitespace
                if url:  # Ensure the line is not empty
                    time.sleep(0.5)  # Pause for 0.5 seconds before each request
                    fetch_url(url, writer, use_proxy)
    except FileNotFoundError:
        print(f"File not found: {file_name}")

# Main execution
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Fetch URLs from a file and log their statuses. Proxy is enabled by default.")
    parser.add_argument("--input", type=str, help="Path to the input file containing URLs", required=True)
    parser.add_argument("--output", type=str, help="Path to the output CSV file (optional; defaults to 'test-connectivity-output.csv')", default="test-connectivity-output.csv")
    parser.add_argument("--no-proxy", action="store_true", help="Disable proxy usage")
    args = parser.parse_args()

    input_file = args.input
    output_file = args.output
    use_proxy = not args.no_proxy  # Proxy is enabled by default; --no-proxy disables it

    read_urls_from_file(input_file, output_file, use_proxy)