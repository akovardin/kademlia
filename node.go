package kademlia

import (
	"encoding/hex"
	"math/rand"
	"time"
)

const NodeLength = 20

type Node [NodeLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}

func NewNode(data string) Node {
	decoded, _ := hex.DecodeString(data)
	id := Node{}
	for i := 0; i < NodeLength; i++ {
		id[i] = decoded[i]
	}

	return id
}

func NewRandomNode() Node {
	id := Node{}
	for i := 0; i < NodeLength; i++ {
		id[i] = uint8(rand.Intn(256))
	}

	return id
}

func (node Node) String() string {
	return hex.EncodeToString(node[0:NodeLength])
}

func (node Node) Equals(other Node) bool {
	for i := 0; i < NodeLength; i++ {
		if node[i] != other[i] {
			return false
		}
	}
	return true
}

func (node Node) Less(other interface{}) bool {
	for i := 0; i < NodeLength; i++ {
		if node[i] != other.(Node)[i] {
			return node[i] < other.(Node)[i]
		}
	}
	return false
}

func (node Node) Xor(other Node) Node {
	ret := Node{}
	for i := 0; i < NodeLength; i++ {
		ret[i] = node[i] ^ other[i]
	}
	return ret
}

// разобраться что тут происходит
func (node Node) PrefixLen() (ret int) {
	for i := 0; i < NodeLength; i++ {
		for j := 0; j < 8; j++ {
			if (node[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return NodeLength*8 - 1
}
