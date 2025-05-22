import json
import csv
import requests
import argparse

# Define input and output files
input_file = "urls.txt"  # File containing URLs line by line
output_file = "output.csv"

# Parse command-line arguments
parser = argparse.ArgumentParser(description="Extract paths from OpenAPI documents.")
parser.add_argument("--no-proxy", action="store_true", help="Disable proxy settings")
args = parser.parse_args()

# Define default proxy settings
proxies = {"http": "http://localhost:8080", "https": "http://localhost:8080"} if not args.no_proxy else {}

# Open output CSV file for writing
with open(output_file, mode="w", newline="") as csv_file:
    writer = csv.writer(csv_file)
    writer.writerow(["URL", "Paths"])  # CSV header

    # Read URLs from the file
    with open(input_file, "r") as f:
        urls = f.read().splitlines()

    for url in urls:
        try:
            # Make GET request with proxy settings and SSL verification disabled
            response = requests.get(url, proxies=proxies, verify=False)
            response.raise_for_status()  # Raise error for bad responses

            # Parse JSON response
            data = response.json()
            paths = list(data.get("paths", {}).keys())
            paths_str = "; ".join(paths)  # Separate paths with semicolon

            # Print to console
            print(f"Processed: {url} -> Paths: {paths_str}")

            # Write to CSV
            writer.writerow([url, paths_str])

        except requests.RequestException:
            print(f"Error fetching: {url}")
        except json.JSONDecodeError:
            print(f"Invalid JSON at: {url}")

print(f"CSV saved: {output_file}")