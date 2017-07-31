package kademlia

import (
	"container/list"
)

const BucketSize = 20

type Contact struct {
	Node    Node
	Address string
}

type RoutingTable struct {
	contact Contact
	buckets [NodeLength * 8]*list.List
}

func NewRoutingTable(contact Contact) *RoutingTable {
	table := &RoutingTable{}

	for i := 0; i < NodeLength*8; i++ {
		table.buckets[i] = list.New()
	}
	table.contact = contact
	return table
}

func (t *RoutingTable) Update(contact Contact) {
	prefix := contact.Node.Xor(t.contact.Node).PrefixLen()

	bucket := t.buckets[prefix]

	var element *list.Element
	for e := bucket.Front(); e != nil; e = e.Next() {
		if e.Value.(Contact).Node.Equals(t.contact.Node) {
			element = e
		}
	}

	if element == nil {
		if bucket.Len() <= BucketSize {
			bucket.PushFront(contact)
		}
		// todo: Handle insertion when the list is full by evicting old elements if
		// they don't respond to a ping.
	} else {
		bucket.MoveToFront(element)
	}
}

func (t *RoutingTable) FindClosest(target Node, count int) []Contact {
	prefix := target.Xor(t.contact.Node).PrefixLen()
	contacts := []Contact{}

	for i := prefix; i >= 0; i-- {
		elt := t.buckets[i].Front()
		for elt != nil {
			if len(contacts) < BucketSize {
				contacts = append(contacts, elt.Value.(Contact))
			} else {
				return contacts
			}
			elt = elt.Next()
		}
	}

	for i := prefix + 1; i < NodeLength; i++ {
		elt := t.buckets[i].Front()
		for elt != nil {
			if len(contacts) < BucketSize {
				contacts = append(contacts, elt.Value.(Contact))
				elt = elt.Next()
			} else {
				return contacts
			}
		}
	}

	if len(contacts) > count {
		contacts = contacts[:count]
	}

	return contacts
}

// special debug method
func (t *RoutingTable) List() []Contact {
	contacts := []Contact{}

	for _, bucket := range t.buckets {
		elt := bucket.Front()
		for elt != nil {
			contacts = append(contacts, elt.Value.(Contact))
			elt = elt.Next()
		}
	}

	return contacts
}
