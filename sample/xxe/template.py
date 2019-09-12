import xml.etree.ElementTree as ET

filename = '{{ .FILENAME }}'

tree = ET.parse(filename)
root = tree.getroot()

for child in root:
    print(child.tag, child.attrib)
