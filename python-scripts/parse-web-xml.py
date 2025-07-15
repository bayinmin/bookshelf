import argparse
import xml.etree.ElementTree as ET

def get_default_namespace(root):
    # Extract namespace from the root tag
    if root.tag.startswith('{'):
        return root.tag.split('}')[0][1:]  # strip { and }
    return ''

def extract_key_from_xml(file_path, key='url-pattern'):
    try:
        tree = ET.parse(file_path)
        root = tree.getroot()

        # Dynamically determine the namespace
        ns_uri = get_default_namespace(root)
        ns = {'ns': ns_uri} if ns_uri else {}

        print(f"Extracting '{key}' from '{file_path}':")

        search_expr = f'.//ns:{key}' if ns else f'.//{key}'
        elements = root.findall(search_expr, ns)

        if not elements:
            print(f"No elements found for tag: {key}")
        else:
            for elem in elements:
                print(f"{elem.text.strip() if elem.text else '(empty)'}")

    except ET.ParseError as e:
        print(f"Failed to parse XML: {e}")
    except FileNotFoundError:
        print(f"File not found: {file_path}")
    except Exception as e:
        print(f"Error: {e}")

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Extract specific tag values from an XML file.')
    parser.add_argument('--file', default='web.xml', help='Path to the XML file (default: web.xml)')
    parser.add_argument('--key', default='url-pattern', help='Tag name to extract (default: url-pattern)')

    args = parser.parse_args()
    extract_key_from_xml(args.file, args.key)