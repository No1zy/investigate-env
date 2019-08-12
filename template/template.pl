#!/usr/bin/perl
use URI;

$url = '{{ .URL }}';
$u = URI->new($url);
print 'scheme = ' . $u->scheme . "\n";
print 'userinfo = ' . $u->userinfo . "\n";
print 'host = ' . $u->host . "\n";
print 'port = ' . $u->port . "\n";
print 'opaque = ' . $u->opaque . "\n";
print 'path = ' . $u->path . "\n";
print 'query = ' . $u->query . "\n";
print 'fragment = ' . $u->fragment . "\n";
print 'canonical = ' . $u->canonical . "\n";
