# watchserve
[watchserve](https://github.com/romshark/watchserve) is an HTTP file server that's watching the served file for updates and automatically reloads the page in any modern JavaScript-capable browser.

## Install

1. download the latest version of [Go](https://golang.org/).
2. run the following command:
```
go get github.com/romshark/watchserve
```

## How to use

```
watchserve -host localhost:8080 -f myfile.txt
```

- `-help` prints help.
- `-f` specifies the file to watch.
- `-host` specifies the server address to listen on.
