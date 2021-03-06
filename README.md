# Another Memory ID

[![Go Report Card](https://goreportcard.com/badge/github.com/anothermemory/id)](https://goreportcard.com/report/github.com/anothermemory/id)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/anothermemory/id)
[![Coveralls github](https://img.shields.io/coveralls/github/anothermemory/id.svg?style=flat-square)](https://coveralls.io/github/anothermemory/id)
[![Release](https://img.shields.io/github/release/anothermemory/id.svg?style=flat-square)](https://github.com/anothermemory/id/releases/latest)
[![license](https://img.shields.io/github/license/anothermemory/id.svg?style=flat-square)](LICENSE.md)
[![Codacy grade](https://img.shields.io/codacy/grade/b37d4ed0bbab42a29a05d81829ca1dbb.svg?style=flat-square)](https://www.codacy.com/app/anothermemory/id)

This library contains interface and set of implementations for generating different identifiers. There are multiple
implementations available. For production usage identifiers must be mostly unique. For testing purposes it's often much 
easier to use mocked generator so it will generate the same id each time.

Table of Contents
=================

* [Another Memory ID](#another-memory-id)
  * [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installing](#installing)
    * [See it in action](#see-it-in-action)
      * [Mocked ID for testing](#mocked-id-for-testing)
      * [Real ID for production](#real-id-for-production)
  * [Built With](#built-with)
  * [Contributing](#contributing)
  * [Versioning](#versioning)
  * [Authors](#authors)
  * [License](#license)

## Getting Started

### Prerequisites

The project is tested with go 1.9 but probably will work with earlier versions.

### Installing

Simple go get it

```bash
go get github.com/anothermemory/id
```

### See it in action

#### Mocked ID for testing

```go
package id_test

import (
    "fmt"
    "github.com/anothermemory/id"
)

func Example() {
    g := id.NewMock("123")

    fmt.Println(g.Generate())
    fmt.Println(g.Generate())

    // Output:
    // 123
    // 123
}
```

#### Real ID for production

```go
package id_test

import (
    "fmt"
    "github.com/anothermemory/id"
)

func Example() {
    g := id.NewUUID()

    fmt.Println(g.Generate())
    fmt.Println(g.Generate())

    // Output:
    // 2118679c-d6f6-4848-8084-aaab790fd1ae
    // 394d79cb-803f-4807-a9bc-00d84d5f9ad3
}
```

## Built With

* [dep](https://github.com/golang/dep) - The dependency management tool for Go

## Contributing

Please read [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for details on our code of conduct, and [CONTRIBUTING.md](CONTRIBUTING.md) for details on the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/anothermemory/id/tags). 

## Authors

* **Vyacheslav Enis**

See also the list of [contributors](https://github.com/anothermemory/id/contributors) who participated in this project.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md) file for details
