package kademlia

import "testing"

func TestNewKademlia(t *testing.T) {
	k := NewKademlia(Contact{Node: NewRandomNode()})

	t.Log(k)
}
