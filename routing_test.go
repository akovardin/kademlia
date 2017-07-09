package kademlia

import "testing"

func TestNewRoutingTable(t *testing.T) {
	rt := NewRoutingTable(Contact{Node: NewRandomNode()})

	t.Log(rt)
}

func TestRoutingTable_Update(t *testing.T) {
	rt := NewRoutingTable(Contact{Node: NewRandomNode()})

	t.Log(rt.buckets[2].Front())

	rt.Update(Contact{
		Node: NewRandomNode(),
	})

	t.Log(rt.buckets[2].Front())
}

func TestRoutingTable_FindClosest(t *testing.T) {
	rt := NewRoutingTable(Contact{Node: NewRandomNode()})

	a := NewNode("48656c6c6f20476f706865722148656c6c6f2041")
	b := NewNode("48656c6c6f20476f706865722148656c6c6f2042")
	c := NewNode("48656c6c6f20476f706865722148656c6c6f2043")
	d := NewNode("48656c6c6f20476f706865722148656c6c6f2045")

	x := NewNode("48756c6c6f20476f706865722148656c6c6f2044")

	rt.Update(Contact{
		Node:    a,
		Address: "a",
	})
	rt.Update(Contact{
		Node:    b,
		Address: "b",
	})
	rt.Update(Contact{
		Node:    c,
		Address: "c",
	})
	rt.Update(Contact{
		Node:    d,
		Address: "d",
	})

	contacts := rt.FindClosest(x, 2)
	for _, v := range contacts {
		t.Log(v.Address)
	}
}
