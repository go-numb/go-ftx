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

	lev, err := client.Leverage(5)
	if err != nil {
		client.Logger.Error(err)
    }

    fmt.Printf("%v\n", lev)
    // -> success nil, false err

    o, err := client.Order(&RequestForOrder{
    Market: "ETH-PERP",
    Type:   LIMIT,
    Side:   BUY,
    Price:  1,
    Size:   1,
    // and more options
	})
	if err != nil {
		client.Logger.Error(err)
    }
    

	ok, err := client.CancelByID(o.ID)
	if err != nil {
		client.Logger.Error(err)
    }

    fmt.Println(ok)
    // ok is status comment

    
}
```


## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-ftx/blob/master/LICENSE)