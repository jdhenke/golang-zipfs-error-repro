# golang `zipfs` bug repro

This repo illustrates **[a fix](https://github.com/golang/tools/compare/master...jdhenke:fix/zipfs-ls-root?expand=1&ts=4)** for `zipfs` so it can `ReadDir("/")` per [golang/go#12743](https://github.com/golang/go/issues/12743).

## Installation

Requires [go](https://golang.org/doc/install) 1.5.

```
go get github.com/jdhenke/golang-zipfs-error-repro
cd $GOPATH/src/github.com/jdhenke/golang-zipfs-error-repro
git submodule init
git submodule update
go get
```

## Usage

This repo uses a [vendored clone](vendor/golang.org/x) of the [tools repo](https://github.com/golang/tools) with **[a fix](https://github.com/golang/tools/compare/master...jdhenke:fix/zipfs-ls-root?expand=1&ts=4)** for [this bug](https://github.com/golang/go/issues/12743), so...

`GO15VENDOREXPERIMENT=0` uses the buggy, public code and *cannot* list `/`

```
$ GO15VENDOREXPERIMENT=0 go run main.go
2015/09/24 20:06:36 file not found: /
exit status 1
```

`GO15VENDOREXPERIMENT=1` uses the fixed, vendored code and *can* list `/`

```
$ GO15VENDOREXPERIMENT=1 go run main.go
2015/09/24 20:06:38 Found: bar
2015/09/24 20:06:38 Found: foo
```
