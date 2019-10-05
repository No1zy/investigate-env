import xml.etree.ElementTree as ET

filename = '{{ .VARIABLE }}'

tree = ET.parse(filename)
root = tree.getroot()

for child in root:
    print(child.tag, child.attrib)
