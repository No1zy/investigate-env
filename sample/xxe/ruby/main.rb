require 'rexml/document'

xml = File.open('{{ .FILENAME }}').read
doc = REXML::Document.new(xml)

entities = doc.doctype.entities

p entities["file"]
