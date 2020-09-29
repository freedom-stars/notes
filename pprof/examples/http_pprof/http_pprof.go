// http_pprof
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// 模拟有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:

		}
	}
}

func main() {
	for i := 0; i < 8; i++ {
		go logicCode()
	}

	log.Fatal(http.ListenAndServe(":6060", nil))
}
