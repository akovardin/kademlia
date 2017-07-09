package kademlia

import (
	"testing"
)

func TestNewNodeId(t *testing.T) {
	id := NewNode("48656c6c6f20476f706865722148656c6c6f2047")

	t.Log(id)
}

func TestNewRandomNodeId(t *testing.T) {
	id := NewRandomNode()

	t.Log(id)
}

func TestNodeId_String(t *testing.T) {
	id := NewNode("48656c6c6f20476f706865722148656c6c6f2047")

	t.Log(id)
}

func TestNodeId_Equals(t *testing.T) {
	right := NewNode("48656c6c6f20476f706865722148656c6c6f2047")
	left := NewNode("48656c6c6f20476f706865722148656c6c6f2047")

	t.Log(right.Equals(left))
}

func TestNodeId_Less(t *testing.T) {
	right := NewNode("48656c6c6f20476f706865722148656c6c6f2047")
	left := NewNode("48656c6c6f20476f706865722148656c6c6f2046")

	t.Log(right.Less(left))
}

func TestNode_PrefixLen(t *testing.T) {
	right := NewNode("48656c6c6f20476f706865722148656c6c6f2047")
	left := NewNode("48656c6c6f20476f706865722148656c6c6f2046")

	t.Log(right.PrefixLen())
	t.Log(left.PrefixLen())
}

func TestNode_Xor(t *testing.T) {
	right := NewNode("48656c6c6f20476f706865722148656c6c6f2047")
	left := NewNode("48656c6c6f20476f706865722148656c6c6f2047")

	t.Log(right.Xor(left).PrefixLen())
}
