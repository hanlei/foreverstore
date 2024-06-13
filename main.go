package main

import (
	"fmt"
	"github.com/hanlei/foreverstore/p2p"
	"log"
)

func main() {
	fmt.Println("We Start Listen!")

	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
