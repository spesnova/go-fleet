# go-fleet
Go client library for fleet

[![Build Status](https://travis-ci.org/obukhov/go-fleet.svg?branch=master)](https://travis-ci.org/obukhov/go-fleet)

## Install

```
go get github.com/spesnova/go-fleet
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/spesnova/go-fleet/fleet"
)

func main() {
    client := fleet.NewClient("http://10.1.42.1:49153")

	opts := []*fleet.UnitOption{
		&fleet.UnitOption{
			Section: "Unit",
			Name:    "Description",
			Value:   "Useless infinite loop",
		},
		&fleet.UnitOption{
			Section: "Service",
			Name:    "ExecStart",
			Value:   "/bin/bash -c \"while true; do echo 'hello' && sleep 1; done\"",
		},
	}

	err := client.Start("hello.service", opts)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Unit have started")
	}
}
```

## License
See LICENSE file
