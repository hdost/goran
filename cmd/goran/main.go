package main

import (
	"fmt"
        "github.com/hdost/goran/pkg/server"
)

func main(){
	fmt.Println("Here's the leader.");
        server.Loop();
}
