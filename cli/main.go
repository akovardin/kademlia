package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/horechek/kademlia"
)

var (
	address  = flag.String("address", "127.0.0.1:3000", "address")
	neighbor = flag.String("neighbor", "", "id/address:port")
)

func main() {
	flag.Parse()

	node := kademlia.NewRandomNode()
	log.Println("node id:", node)

	k := kademlia.NewKademlia(kademlia.Contact{
		Node:    node,
		Address: *address,
	})

	parts := strings.Split(*neighbor, "/")

	// @todo: move to flag parsing
	if len(parts) != 2 {
		log.Fatalln("undefined neighbor")
	}

	id := parts[0]
	addr := parts[1]

	k.Routes.Update(kademlia.Contact{
		Node:    kademlia.NewNode(id),
		Address: addr,
	})

	go k.Handle()

	loop(k)
}

func loop(k *kademlia.Kademlia) {
	var input string

	for {
		fmt.Print("> ")
		if _, err := fmt.Scan(&input); err != nil {
			log.Println("scan error:", err)
			continue
		}

		fmt.Println(input)

		// parse commands, find in examples
	}
}
