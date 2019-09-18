# Informer
Go library for exposing program metadata, version and build information.

## TL;DR

See the [example](./example) directory. That directory is an example of how a package that uses Informer might look.

```bash
cd example/
ls
# cmd  doc.go  example  info.go  Makefile
make
# CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -X "github.com/cisco-cx/informer/example.Program=example" -X "github.com/cisco-cx/informer/example.License=Apache-2.0" -X "github.com/cisco-cx/informer/example.URL=https://github.com/cisco-cx/informer/example" -X "github.com/cisco-cx/informer/example.BuildUser=vagrant" -X "github.com/cisco-cx/informer/example.BuildDate=2019-09-18T19:50:30+0000" -X "github.com/cisco-cx/informer/example.Version=f2211ec" -X "github.com/cisco-cx/informer/example.Revision=f2211ec" -X "github.com/cisco-cx/informer/example.Branch=master"' ./cmd/example
ls -l example
# -rwxr-xr-x 1 vagrant vagrant 1423424 Sep 18 19:50 example
./example
# (metadata=(program=example, license=Apache-2.0, url=https://github.com/cisco-cx/informer/example), versionInfo=(version=f2211ec, branch=master, revision=f2211ec), buildInfo=(go=go1.12.7, user=vagrant, date=2019-09-18T19:50:30+0000))
```

## Usage

* Step 1: Define global variables at the root (or other) package of your project.
* Step 2: Import the package you used in Step 1 in a `cmd` (or other) package.
* Step 3: Pass in values for your global variables at build time using `ldflags`.
* Step 4: Call `v1.NewInfoService` and pass in your global vars as args to it.
* Step 5: Use the various methods on the `informer.InfoService` interface as you wish.

## Metrics

Exceprt of an HTTP GET response from a program that implements Informer:

```
# TYPE informer_go gauge
informer_go{branch="master",build_date="2019-09-18T19:50:30+0000",build_user="vagrant",go_version="go1.12.7",license="Apache-2.0",program="example",revision="f2211ec",url="https://github.com/cisco-cx/informer/example",version="f2211ec"} 1
```

## Roadmap

* Demonstrate the optional InfoCollector in the example directory.
* Use the shared `mock` package in `v1_test` package.

## Credit

This project is inspired by:

- https://github.com/benbjohnson/peapod
- https://github.com/prometheus/common/blob/master/version/info.go
- https://golang.org/src/encoding
- https://golang.org/src/encoding/base64/base64.go
- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- https://povilasv.me/exposing-go-modules-to-prometheus/
- https://travix.io/encapsulating-dependencies-in-go-b0fd74021f5a
