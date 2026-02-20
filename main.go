package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	society := society{members: []node{}, connections: map[connection]struct{}{}}
	society.populateMembers(10)
	society.connectNodes(100)

	v, _ := society.displayNodes()
	data, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	w.Write(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

type node struct {
	id uint
}

type connection struct {
	A *node
	B *node
}

type connectionDTO struct {
	Edges [][]uint `json:"edges"`
	Order int      `json:"order"`
	Size  int      `json:"size"`
}

func newConnection(a *node, b *node) connection {
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
		conn := newConnection(&s.members[indexOne], &s.members[indexTwo])
		fmt.Println(s.members[indexOne])
		fmt.Println(s.members[indexTwo])
		s.connections[conn] = struct{}{}
	}

	fmt.Println(s.connections)
}

func (s *society) displayNodes() (connectionDTO, error) {
	if len(s.connections) == 0 {
		return connectionDTO{}, errors.New("No Connections to display")
	}

	var sliceConnections [][]uint

	for k := range s.connections {
		newEdge := []uint{}
		newEdge = append(newEdge, k.A.id)
		newEdge = append(newEdge, k.B.id)
		sliceConnections = append(sliceConnections, newEdge)
	}

	returnConnectionData := connectionDTO{Edges: sliceConnections,
		Order: len(s.members),
		Size:  len(sliceConnections)}

	return returnConnectionData, nil
}
