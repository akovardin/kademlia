package main

import (
	"fmt"
	"strings"
	"errors"
	
	"github.com/horechek/kademlia"
)

type commandor interface {
	execute(k *kademlia.Kademlia) error
}

func parse(input string) (commandor, error) {
	// @todo: set stdout

	var c commandor

	switch {
	case strings.HasPrefix(input, "join"):
		c = &join{
			input: input,
		}
	case strings.HasPrefix(input, "list"):
		c = &list{}
	case strings.HasPrefix(input, "find"):
		c = &find{
			input: input,
		}

		// connect
		// send

	case strings.HasPrefix(input, "echo"):
		c = &echo{
			input: input,
		}
	case len(input) == 0:
		c = &br{}
	default:
		return nil, errors.New("undefined command '" + input + "'")
	}

	return c, nil
}

type br struct {
	input string
}

func (c *br) execute(k *kademlia.Kademlia) error {
	return nil
}

type echo struct {
	input string
}

func (c *echo) execute(k *kademlia.Kademlia) error {
	fmt.Println(c.input)

	return nil
}

type ping struct {
}

func (p *ping) execute(k *kademlia.Kademlia) error {
	return nil
}

type find struct {
	input string
}

func (f *find) execute(k *kademlia.Kademlia) error {
	var id string
	_, err := fmt.Sscanf(f.input, "find %s", &id)
	if err != nil {
		return err
	}

	c := k.FindNode(kademlia.NewNode(id))
	if c == nil {
		fmt.Println("not found")
		return nil
	}

	fmt.Printf("%s | %s\n", c.Node, c.Address)

	return nil
}

type join struct {
	input string
}

func (j *join) execute(k *kademlia.Kademlia) error {
	var (
		id   string
		addr string
	)

	_, err := fmt.Sscanf(j.input, "join %s %s", &id, &addr)
	if err != nil {
		return err
	}

	k.Routes.Update(kademlia.Contact{
		Node:    kademlia.NewNode(id),
		Address: addr,
	})

	return nil
}

type list struct{}

func (l *list) execute(k *kademlia.Kademlia) error {
	for _, c := range k.Routes.List() {
		fmt.Printf("%s | %s\n", c.Node, c.Address)
	}

	return nil
}
