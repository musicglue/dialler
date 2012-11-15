Dialler
=======

Dialler works by opening a connection (tcp or udp) to an arbitrary host/port with an arbitrary timeout.

### Usage
To install and use Dialler, you need to do the following steps:

```shell
go get github.com/musicglue/dialler
go install github.com/musicglue/dialler
dialler -address google.com:80 -timeout 1000 -protocol tcp 
```

### Options
The following flags are supported:

```
-address  = A string of url/ip and a port, such as "localhost:80" or "127.0.0.1:80"
-timeout  = The time a connection must be established within (milliseconds), e.g. 1000
-protocol = Whether to test connecting over TCP or UDP
```

Enjoy!
