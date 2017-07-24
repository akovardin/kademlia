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

func (table *RoutingTable) Update(contact Contact) {
	prefix := contact.Node.Xor(table.contact.Node).PrefixLen()

	bucket := table.buckets[prefix]

	var element *list.Element
	for e := bucket.Front(); e != nil; e = e.Next() {
		if e.Value.(Contact).Node.Equals(table.contact.Node) {
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

func (table *RoutingTable) FindClosest(target Node, count int) []Contact {
	prefix := target.Xor(table.contact.Node).PrefixLen()
	contacts := []Contact{}

	for i := prefix; i >= 0; i-- {
		elt := table.buckets[i].Front()
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
		elt := table.buckets[i].Front()
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
