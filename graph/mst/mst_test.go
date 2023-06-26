package mst

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMST(t *testing.T) {
	edges := FindMST[string](
		[]string{"A", "B", "C", "D", "E", "F"},
		edges[string]{
			{Node1: "A", Node2: "B", Weight: 5},
			{Node1: "A", Node2: "D", Weight: 9},
			{Node1: "A", Node2: "E", Weight: 5},
			{Node1: "A", Node2: "C", Weight: 4},
			{Node1: "A", Node2: "F", Weight: 11},
			{Node1: "B", Node2: "D", Weight: 3},
			{Node1: "B", Node2: "C", Weight: 10},
			{Node1: "B", Node2: "F", Weight: 2},
			{Node1: "B", Node2: "E", Weight: 7},
			{Node1: "C", Node2: "F", Weight: 5},
			{Node1: "C", Node2: "D", Weight: 7},
			{Node1: "C", Node2: "F", Weight: 4},
			{Node1: "C", Node2: "E", Weight: 1},
			{Node1: "D", Node2: "F", Weight: 8},
			{Node1: "E", Node2: "F", Weight: 5},
		},
	)
	log.Println("edges:", edges)
	assert.Equal(t, 5, len(edges))
}
