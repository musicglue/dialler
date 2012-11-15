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

