vcl 4.0;

backend web {
    .host = "node";
    .port = "3000";
    .connect_timeout = 15s;
}

backend websocket {
    .host = "flask";
    .port = "5000";
    .connect_timeout = 15s;
}

sub vcl_recv {
    if (req.url == "/") {
            set req.backend_hint = web;
    }
    if (req.url ~ "^/socket.io/") {
            set req.backend_hint = websocket;
    }
    if (req.http.upgrade ~ "(?i)websocket") {
        return (pipe);
    }
}

sub vcl_pipe {
    if (req.http.upgrade) {
        set bereq.http.upgrade = req.http.upgrade;
        set bereq.http.connection = req.http.connection;
    }
}
