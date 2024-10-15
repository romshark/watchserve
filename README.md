<a href="https://github.com/romshark/watchserve/actions?query=workflow%3ACI">
  <img src="https://github.com/romshark/watchserve/workflows/CI/badge.svg" alt="GitHub Actions: CI">
</a>
<a href="https://goreportcard.com/report/github.com/romshark/watchserve">
  <img src="https://goreportcard.com/badge/github.com/romshark/watchserve" alt="GoReportCard">
</a>

# watchserve
[watchserve](https://github.com/romshark/watchserve) is an HTTP file server that's watching the served file for updates and automatically reloads the page in any modern JavaScript-capable browser.

## Install

Download one of the packaged executables from the latest [release version](https://github.com/romshark/watchserve/releases).

Alternatively, you can use the Go toolchain to install watchserve.

### Through the Go toolchain

1. Download the latest version of [Go](https://go.dev/).
2. Run the following command:
```
go install github.com/romshark/watchserve@latest
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
