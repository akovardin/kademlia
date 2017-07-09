package main

import (
	"flag"
	"log"

	"github.com/horechek/kademlia"
)

var (
	address   = flag.String("address", "127.0.0.1:3000", "address")
	find      = flag.String("find", "", "find")
	closenode = flag.String("closenode", "", "closenode")
	closeaddr = flag.String("closeaddr", "", "closeaddr")
)

func main() {
	flag.Parse()

	node := kademlia.NewRandomNode()

	log.Println("node id:", node)

	k := kademlia.NewKademlia(kademlia.Contact{
		Node:    node,
		Address: *address,
	})

	if *closenode != "" && *closeaddr != "" {
		k.Routes.Update(kademlia.Contact{
			Node:    kademlia.NewNode(*closenode),
			Address: *closeaddr,
		})
	}

	if *find != "" {
		contact := k.FindNode(kademlia.NewNode(*find))
		log.Println("result", contact)
	}

	k.Handle()
}
