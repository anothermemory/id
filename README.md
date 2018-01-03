# Another Memory ID

This library contains interface and set of implementations for generating different identifiers. There are multiple
implementations available. For production usage identifiers must be mostly unique. For testing purposes it's often much 
easier to use mocked generator so it will generate the same id each time.

## Getting Started

### Prerequisites

The project is tested with go 1.9 but probably will work with earlier versions.

### Installing

Simple go get it

```
go get github.com/anothermemory/id
```

### See it in action

#### Mocked ID for testing

```
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

```
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

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/anothermemory/id/tags). 

## Authors

* **Vyacheslav Enis**

See also the list of [contributors](https://github.com/anothermemory/id/contributors) who participated in this project.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md) file for details
