require 'uri'

url = "{{ .URL }}"

p uri = URI.parse(url)

p 'scheme: ' + uri.scheme    # => "http"
p 'userinfo: ' + (uri.userinfo.nil? ? 'nil' : uri.userinfo)
p 'host: ' + (uri.host.nil? ? 'nil' : uri.host)       # => "www.ruby-lang.org"
p 'port: ' + (uri.port.nil? ? 'nil' : uri.port.to_s)      # => 80
p 'registry: ' + (uri.registry.nil? ? 'nil' : uri.registry)
p 'path: ' + (uri.path.nil? ? 'nil' : uri.path)
p 'opaque: ' + (uri.opaque.nil? ? 'nil' : uri.opaque)
p 'query: ' + (uri.query.nil? ? 'nil' : uri.query)
p 'fragment: ' + (uri.fragment.nil? ? 'nil' : uri.fragment)
#p 'arg_check: ' + (uri.arg_check.nil? ? 'nil' : uri.arg_check)

