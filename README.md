[![TravisBuildStatus](https://api.travis-ci.org/sinlovgo/timezone.svg?branch=master)](https://travis-ci.org/sinlovgo/timezone)
[![GoDoc](https://godoc.org/github.com/sinlovgo/timezone?status.png)](https://godoc.org/github.com/sinlovgo/timezone/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlovgo/timezone)](https://goreportcard.com/report/github.com/sinlovgo/timezone)
[![codecov](https://codecov.io/gh/sinlovgo/timezone/branch/master/graph/badge.svg)](https://codecov.io/gh/sinlovgo/timezone)

## for what

- this project used to golang timezone

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/sinlovgo/timezone.git

# test depends see full version
$ go list -v -m -versions github.com/sinlovgo/timezone
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/sinlovgo/timezone.git | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## usage

### set timezone

## dev

```bash
make init
```

- test code

```bash
make test
```

add main.go file and run

```bash
make run
```

# dev

## evn

- golang sdk 1.13+

## docker

base docker file can replace
