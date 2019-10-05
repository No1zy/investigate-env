require 'rexml/document'

xml = File.open('{{ .VARIABLE }}').read
doc = REXML::Document.new(xml)

entities = doc.doctype.entities

p entities["file"]
