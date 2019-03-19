package main

import (
	"fmt"
	"time"
)

func main() {
	for {

		fmt.Println(time.Now().Format(time.RFC850) + " Hello Skaffold - do you like the demo???")

		time.Sleep(time.Second * 5)
	}
}
