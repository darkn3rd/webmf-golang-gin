# Gin WebMF

This is brief tutorial on using Gin web microframework for the Go Language.

## Install and Run

With Go install, setup [GOPATH](https://golang.org/doc/code.html#GOPATH), and run these in bash:

```bash
# run server
export PORT=6000
go run server.go &

# run tests manually
export WEBSERVER=localhost
curl -i ${WEBSERVER}:${PORT}/
curl -i ${WEBSERVER}:${PORT}/hello/Simon

fg
# CTRL-C
```

## Resources

* https://gin-gonic.com
