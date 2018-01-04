# Another Memory Clock

[![Go Report Card](https://goreportcard.com/badge/github.com/anothermemory/clock)](https://goreportcard.com/report/github.com/anothermemory/clock)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/anothermemory/clock)
[![Coveralls github](https://img.shields.io/coveralls/github/anothermemory/clock.svg?style=flat-square)](https://coveralls.io/github/anothermemory/clock)
[![Release](https://img.shields.io/github/release/anothermemory/clock.svg?style=flat-square)](https://github.com/anothermemory/clock/releases/latest)
[![license](https://img.shields.io/github/license/anothermemory/clock.svg?style=flat-square)](LICENSE.md)

This library contains interface and set of implementations for getting current time. For production usage
time must be mostly real time. For testing purposes it's often much easier to use mocked time so it will
return required time each time.

## Getting Started

### Prerequisites

The project is tested with go 1.9 but probably will work with earlier versions.

### Installing

Simple go get it

```bash
go get github.com/anothermemory/clock
```

### See it in action

#### Mocked clock for testing

```go
package clock_test

import (
    "time"
    
    "github.com/anothermemory/clock"
)

func ExampleMock() {
	var dummyTime = time.Date(2017, 11, 24, 17, 0, 0, 0, time.Local)
	c := clock.NewMock(dummyTime)
	c.Now() // Will return dummyTime each time
}

func ExampleMockPartial() {
	firstTime := time.Date(2017, 11, 24, 17, 0, 0, 0, time.Local)
    secondTime := time.Date(2018, 11, 24, 17, 0, 0, 0, time.Local)

    c := clock.NewMockPartial(firstTime, secondTime)
    c.Now() // This will return firstTime
    c.Now() // This will return secondTime
    c.Now() // This will return real time
}
```

#### Real ID for production

```go
package clock_test

import (
    "time"
    
    "github.com/anothermemory/clock"
)

func ExampleReal() {
    c := clock.NewReal()
    c.Now() // Will return real current time 
}
```

## Built With

* [dep](https://github.com/golang/dep) - The dependency management tool for Go

## Contributing

Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct, and [CONTRIBUTING.md](CONTRIBUTING.md) for details on the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/anothermemory/clock/tags). 

## Authors

* **Vyacheslav Enis**

See also the list of [contributors](https://github.com/anothermemory/clock/contributors) who participated in this project.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md) file for details
