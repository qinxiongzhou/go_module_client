package main

import (
	"fmt"
	go_module_server "github.com/qinxiongzhou/go_module_server"
	tour "github.com/qinxiongzhou/go_module_server/module/tour"
)

func main() {
	fmt.Println(go_module_server.Calculate(2))
	fmt.Print(tour.Calculate(2))
}
