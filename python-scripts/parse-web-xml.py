import argparse
import xml.etree.ElementTree as ET

def extract_key_from_xml(file_path, key='url-pattern'):
    try:
        tree = ET.parse(file_path)
        root = tree.getroot()

        ns = {'ns': 'http://xmlns.jcp.org/xml/ns/javaee'}
        print(f"Extracting '{key}' from '{file_path}':")

        elements = root.findall(f'.//ns:{key}', ns)
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