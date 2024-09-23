# /// script
# dependencies = [
#   "lxml",
# ]
# ///

import os
from lxml import etree
import argparse


def get_components(directory):
    components = []
    for item in os.listdir(directory):
        # Ignore hidden directories and 'prompts' directory
        if item.startswith(".") or item == "prompts":
            continue

        item_path = os.path.join(directory, item)
        if os.path.isdir(item_path):
            components.append(item.replace(os.sep, "/"))
    return components


def update_xml(file_path, components):
    parser = etree.XMLParser(remove_blank_text=True, strip_cdata=False)
    tree = etree.parse(file_path, parser)
    root = tree.getroot()

    # Find or create the Input element
    input_elem = root.find("Input")
    if input_elem is None:
        input_elem = etree.SubElement(root, "Input")

    # Remove existing Component elements
    for component in input_elem.findall("Component"):
        input_elem.remove(component)

    # Add new Component elements
    for component in components:
        comp_elem = etree.SubElement(input_elem, "Component")
        comp_elem.text = component

    return tree


def main():
    parser = argparse.ArgumentParser(
        description="Update Release Please manifest for a multi-component repository."
    )
    parser.add_argument(
        "directory", help="Path to the parent directory of the components"
    )
    parser.add_argument("xml_file", help="Path to the existing XML file to update")

    args = parser.parse_args()

    components = get_components(args.directory)
    updated_tree = update_xml(args.xml_file, components)

    updated_tree.write(
        args.xml_file, encoding="utf-8", xml_declaration=True, pretty_print=True
    )

    print(f"Manifest file updated: {args.xml_file}")
    print(f"Components included: {', '.join(components)}")


if __name__ == "__main__":
    main()
