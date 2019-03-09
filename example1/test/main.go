package main

import (
	"fmt"
	"time"
)

func main() {
	for {

		fmt.Println(time.Now().Format(time.RFC850) + " Hello Skaffold!")

		time.Sleep(time.Second * 1)
	}
}
