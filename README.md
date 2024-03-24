![logo](./img/logo.png)
# Gee API Framework

Gee API Framework encapsulates the net/http framework, making it easier and faster for us to write backend APIs

**Gee has the following characteristics**
- Fast
- Zero allocation router
- Middleware support
- Routes grouping
- Extendable

## Get Started

### Prerequisites

- **[Go](https://go.dev/)**: any one of the **three latest major** [releases](https://go.dev/doc/devel/release) (we test it with these).

### Getting Gee
```
import "github.com/gitlayzer/gee_api_framework"
```

to your code, and then go [build|run|test] will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `gin` package:

```sh
$ go get -u github.com/gitlayzer/gee_api_framework
```

### Running Gee
First you need to import Gee package for using Gee, one simplest example likes the follow `main.go`:

```go
package main

import gee "github.com/gitlayzer/gee_api_framework"

func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"message": "Hello Gee",
		})
	})

	_ = r.Run(":8888")
}
```

And use the Go command to run the demo:
```
❯ go run .\main.go
2024/03/24 17:23:12 [Gee]  GET - /
```
