[![TravisBuildStatus](https://api.travis-ci.org/sinlovgo/timezone.svg?branch=master)](https://travis-ci.org/sinlovgo/timezone)
[![GoDoc](https://godoc.org/github.com/sinlovgo/timezone?status.png)](https://godoc.org/github.com/sinlovgo/timezone/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlovgo/timezone)](https://goreportcard.com/report/github.com/sinlovgo/timezone)
[![codecov](https://codecov.io/gh/sinlovgo/timezone/branch/master/graph/badge.svg)](https://codecov.io/gh/sinlovgo/timezone)

## for what

- this project used to golang timezone

## depends

in go mod project

```bash
# test version info
$ git ls-remote -q http://github.com/sinlovgo/timezone.git

# test depends see full version
$ go list -v -m -versions github.com/sinlovgo/timezone
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/sinlovgo/timezone | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## usage

### set timezone

```go
import "github.com/sinlovgo/timezone"
func main() {
	// use set of second time string
	timezone.Second()
	// use set of micro time string
	timezone.Micro()
	// timestamp unix
	timezone.TimestampSecond()
	// timestamp unix nano
	timezone.TimestampUnixNano()
	// UTC string
	timezone.UTCSecond()
	// UTC string as Micro
	timezone.UTCMicro()
	// use Asia/Shanghai
	timezone.SetZoneByName(timezone.ZoneAsiaShanghai)

	timezone.ZoneUTC
	timezone.ZoneGMT
	timezone.ZoneCET
	timezone.ZoneCST

	// this location
	timezone.SetZoneLocation()

	timezone.ZoneCEST
}
```

### layout

format for golang

```go
import "github.com/sinlovgo/timezone"
	timezone.LayoutSecond
	timezone.LayoutMicro
	timezone.LayoutISO8601TimeSecond
	timezone.LayoutISO8601TimeMicro
	timezone.LayoutISO8601TimeUTC
```

### now

```go
import "github.com/sinlovgo/timezone"
func main() {
	// use set of second time string
	timezone.Second()
	// use set of micro time string
	timezone.Micro()
	// timestamp unix
	timezone.TimestampSecond()
	// timestamp unix nano
	timezone.TimestampUnixNano()
	// UTC string
	timezone.UTCSecond()
	// UTC string as Micro
	timezone.UTCMicro()
	// use Asia/Shanghai
	timezone.SetZoneByName(timezone.ZoneAsiaShanghai)
	// this location
	timezone.SetZoneLocation()
}
```

### Parse

```go
import "github.com/sinlovgo/timezone"
func main() {
	timestamp, err := timezone.ParseTimestamp("2006-01-02 15:04:05", "2018-07-11 15:07:51")
	timestampSecond, err := timestamp.ParseTimestampSecond("2018-07-11 15:07:51")
}
```

### timeparse

```go
import "github.com/sinlovgo/timezone/timeparse"
func main() {
	testTime := "2020-02-25T16:52:18Z"
	parseLocation, err := Location(timezone.LayoutISO8601TimeSecond, timezone.LayoutISO8601TimeMicro,
		testTime,
		timezone.ZoneUTC, timezone.ZoneAsiaShanghai)
}
```

### Compare

```go
import "github.com/sinlovgo/timezone"
func main() {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocationEQ, err := timestamp.CompareEQ(timeFrom, timeTo, "2006-01-02 15:04:05")
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocation, err := timestamp.CompareGT(timeFrom, timeTo, timestamp.LayoutSecond)
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compare, err := timestamp.CompareSecondLT(timeFrom, timeTo)
}
```

### Calc

```go
import "github.com/sinlovgo/timezone"
func main() {
	calcDurationSecond := time.Duration(-30) * time.Minute
	calcDuration := timestamp.CalcDuration(calcDurationSecond, timestamp.LayoutMicro)

	calcDateSecond := timestamp.CalcDateSecond(2, 3, 10)
	calcDateMicro := timestamp.CalcDateMicro(2, 2, 2)

	calcDay := timestamp.CalcDay(10, LayoutSecond)
	calcDaySecond := timestamp.CalcDaySecond(1)
	calcMonthSecond := timestamp.CalcMonthSecond(-1)
	calcYearSecond := timestamp.CalcYearSecond(-1)
}
```

### start time

```go
import "github.com/sinlovgo/timezone"
func main() {
	yesterdaySecond := StartDaySecond(-1)
	lastMonthMicro := StartMonthMicro(-1)
	lastYearSecond := StartYearSecond(-1)
}
```

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
