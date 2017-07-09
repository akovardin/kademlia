package kademlia

import (
	"fmt"
	"log"
	"net"
	//"time"
	"net/rpc"
)

type Kademlia struct {
	Routes *RoutingTable
	Self   Contact
}

func NewKademlia(self Contact) *Kademlia {
	ret := &Kademlia{
		Routes: NewRoutingTable(self),
		Self:   self,
	}

	return ret
}

func (k *Kademlia) FindNode(target Node) *Contact {
	if k.Self.Node.Equals(target) {
		return &k.Self
	}

	closes := k.Routes.FindClosest(target, 10)
	log.Println("closes", closes)

	for _, contact := range closes {
		if contact.Node.Equals(target) {
			log.Println("find contact in closes", contact)

			// update routing
			k.Routes.Update(contact)

			return &contact
		}

		remote, err := k.RemoteFindNode(contact.Address, target)
		if err != nil {
			log.Println("remote err", err)
		}

		if remote.Address != "" {

			// update routing
			k.Routes.Update(remote)

			return &remote
		}
	}

	return nil
}

func (k *Kademlia) RemoteFindNode(addr string, target Node) (Contact, error) {
	result := Contact{}

	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		return result, err
	}

	defer c.Close()

	err = c.Call("Server.FindNode", target, &result)
	log.Println("find contact in remote", result)

	return result, err
}

func (k *Kademlia) Handle() {
	rpc.Register(&Server{
		k: k,
	})

	ln, err := net.Listen("tcp", k.Self.Address)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

type Server struct {
	k *Kademlia
}

func (s *Server) FindNode(target Node, response *Contact) error {
	contact := s.k.FindNode(target)

	response.Node = contact.Node
	response.Address = contact.Address

	return nil
}

type FindValueResponse struct {
	Contacts []Contact
}

func (s *Server) FindValue(target Node, response *FindValueResponse) error {

	return nil
}
