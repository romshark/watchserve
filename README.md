# watchserve
[watchserve](https://github.com/romshark/watchserve) is an HTTP file server that's watching the served file for updates and automatically reloads the page in any modern JavaScript-capable browser.

## Install

Download one of the packaged executables from the latest [release version](https://github.com/romshark/watchserve/releases).

Alternatively, you can use the Go toolchain to install watchserve.

### Through the Go toolchain

1. Download the latest version of [Go](https://golang.org/).
2. Run the following command:
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
- `-debounce` specifies how much time needs to pass after the last change was detected before a reload is triggered.
- `-no-redirect` disables automatic redirect to the browser on start.
