# go-ftx

FTX exchange API version1

## Description

go-ftx is a go client library for [FTX API Document](https://docs.ftx.com).

## Installation

```
$ go get -u github.com/go-numb/go-ftx
```

## Usage
``` golang
package main

import (
 "fmt"
 "github.com/go-numb/go-ftx"
)


func main() {
	client := ftx.New("<key>", "<secret>", nil)

	coins, err := client.Coins()
	if err != nil {
	   client.Logger.Error(err)
	}

	fmt.Printf("%v\n", coin)

	doSomething()
}
```


## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-ftx/blob/master/LICENSE)