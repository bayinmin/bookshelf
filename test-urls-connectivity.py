import requests
import time
import warnings
import csv

# Suppress warnings
warnings.filterwarnings("ignore")

# Proxy configuration
use_proxy = True  # Set this to False to turn off the proxy
proxy = {
    "http": "http://127.0.0.1:8080",
    "https": "http://127.0.0.1:8080"
}

# Function to make HTTP requests
def fetch_url(url, writer):
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
def read_urls_from_file(file_name, output_file):
    try:
        with open(file_name, "r") as file, open(output_file, "w", newline="") as csvfile:
            writer = csv.writer(csvfile)
            writer.writerow(["URL", "Status Code"])  # Write header
            for line in file:
                url = line.strip()  # Remove any whitespace
                if url:  # Ensure the line is not empty
                    time.sleep(0.5)  # Pause for 0.5 seconds before each request
                    fetch_url(url, writer)
    except FileNotFoundError:
        print(f"File not found: {file_name}")

# Main execution
if __name__ == "__main__":
    input_file = "urls.txt"  # Replace with your text file containing URLs
    output_file = "output.csv"  # Replace with your desired output CSV file name
    read_urls_from_file(input_file, output_file)