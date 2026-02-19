package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	society := society{members: []node{}, connections: map[connection]struct{}{}}
	society.populateMembers(10)
	society.connectNodes(100)
	fmt.Println(society.members)
	fmt.Println(society.connections)

}

type node struct {
	id uint
}

type connection struct {
	A node
	B node
}

func newConnection(a node, b node) connection {
	if a.id > b.id {
		return connection{A: a, B: b}
	}

	return connection{A: b, B: a}
}

type society struct {
	members     []node
	connections map[connection]struct{}
}

func (s *society) populateMembers(n uint) {
	for i := 0; i < int(n); i++ {
		s.members = append(s.members, node{id: uint(i)})
	}
}

func (s *society) connectNodes(n uint) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(n); i++ {
		indexOne := r.Intn(len(s.members))
		indexTwo := r.Intn(len(s.members))
		for indexOne == indexTwo {
			indexTwo = r.Intn(len(s.members))
		}
		conn := newConnection(s.members[indexOne], s.members[indexTwo])
		s.connections[conn] = struct{}{}
	}
}
