package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/horechek/kademlia"
)

var (
	address = flag.String("address", "127.0.0.1:3000", "address")
)

func main() {
	flag.Parse()

	node := kademlia.NewRandomNode()
	fmt.Println("node id:", node)

	k := kademlia.NewKademlia(kademlia.Contact{
		Node:    node,
		Address: *address,
	})

	go k.Handle()

	repl(k)
}

func repl(k *kademlia.Kademlia) {
	bio := bufio.NewReader(os.Stdin)

	for {
		// @todo: use one variant for all io
		fmt.Print("> ")
		input, _, err := bio.ReadLine()
		if err != nil {
			log.Println("scan error:", err)
			continue
		}

		c, err := parse(string(input))
		if err != nil {
			log.Println("execute err:", err)
			continue
		}

		if err := c.execute(k); err != nil {
			log.Println("execute err:", err)
		}
	}
}
