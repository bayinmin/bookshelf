import requests
import time
import warnings
import csv
import argparse

# Suppress warnings
warnings.filterwarnings("ignore")

port = 8080

# Proxy configuration
proxy = {
    "http": f"http://127.0.0.1:{port}",
    "https": f"http://127.0.0.1:{port}"
}

# search string
search_string = ""
search_string_found = ""

row_processed = 0


# Function to make HTTP requests

def fetch_url(url, writer, use_proxy):
    global row_processed
    row_processed += 1
    try:
        if use_proxy:
            response = requests.get(url, proxies=proxy, verify=False)  # Ignore SSL errors with verify=False
        else:
            response = requests.get(url, verify=False)  # Ignore SSL errors with verify=False
            
        if search_string in response.text:
            search_string_found = "Y"
        else:
            search_string_found = "N"
        
        print(f"String found: {search_string_found} | Status Code: {response.status_code} | URL: {url} | response size: {len(response.content)}")
        writer.writerow([search_string_found, response.status_code, url, len(response.content)])  # Write to CSV file
    except requests.exceptions.RequestException as e:
        print(f"URL: {url} - Error: {e}")
        writer.writerow([url, "Error"])  # Write error to CSV file

# Reading URLs from the text file
def read_urls_from_file(file_name, output_file, use_proxy,path=""):
    try:
        with open(file_name, "r") as file, open(output_file, "w", newline="") as csvfile:
            writer = csv.writer(csvfile)
            writer.writerow(["String Found", "Status Code","URL", "Response Length",])  # Write header
            for line in file:
                url = (line.strip() + path).strip()# Remove any whitespace
                if url:  # Ensure the line is not empty
                    time.sleep(0.5)  # Pause for 0.5 seconds before each request
                    fetch_url(url, writer, use_proxy)
    except FileNotFoundError:
        print(f"File not found: {file_name}")
        
        
# Reading URLs from the text file
def read_paths_from_file(file_name, output_file, use_proxy):
    try:
        with open(file_name, "r") as file:
            for path in file:
               read_urls_from_file(input_file, output_file, use_proxy,path)
    except FileNotFoundError:
        print(f"File not found: {file_name}")

# Main execution
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Fetch URLs from a file and log their statuses. Proxy is enabled by default.")
    parser.add_argument("--input", type=str, help="Path to the input file containing URLs", required=True)
    parser.add_argument("--output", type=str, help="Path to the output CSV file (optional; defaults to 'test-connectivity-output.csv')", default="test-connectivity-output.csv")
    parser.add_argument("--paths", type=str, help="Paths to the output CSV file.")
    parser.add_argument("--no-proxy", action="store_true", help="Disable proxy usage")
    args = parser.parse_args()

    input_file = args.input
    output_file = args.output
    paths_file = args.paths
    use_proxy = not args.no_proxy  # Proxy is enabled by default; --no-proxy disables it
    
    print(f"[i] Processing begins....\n\n")

    if paths_file:
        read_paths_from_file(paths_file, output_file, use_proxy)
    else:
        read_urls_from_file(input_file, output_file, use_proxy)
        
    print(f"\n[i] Total row(s) processed: {row_processed}")
    print(f" [i] Processing is complete!")