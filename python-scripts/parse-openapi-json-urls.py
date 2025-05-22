import json
import csv
import requests

# Define input and output files
input_file = "urls.txt"  # File containing URLs line by line
output_file = "output.csv"

# Open output CSV file for writing
with open(output_file, mode="w", newline="") as csv_file:
    writer = csv.writer(csv_file)
    writer.writerow(["URL", "Paths"])  # CSV header

    # Read URLs from the file
    with open(input_file, "r") as f:
        urls = f.read().splitlines()

    for url in urls:
        try:
            # Make GET request
            response = requests.get(url)
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