package topological

import (
	"fmt"
	"go_algorithms/data-structures/graph"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	h := graph.NewDirected()

	_ = h.AddVertex(graph.VertexId(2))
	_ = h.AddVertex(graph.VertexId(3))
	_ = h.AddVertex(graph.VertexId(5))
	_ = h.AddVertex(graph.VertexId(7))
	_ = h.AddVertex(graph.VertexId(8))
	_ = h.AddVertex(graph.VertexId(9))
	_ = h.AddVertex(graph.VertexId(10))
	_ = h.AddVertex(graph.VertexId(11))

	_ = h.AddEdge(graph.VertexId(7), graph.VertexId(8), 1)
	_ = h.AddEdge(graph.VertexId(7), graph.VertexId(11), 1)
	_ = h.AddEdge(graph.VertexId(5), graph.VertexId(11), 1)
	_ = h.AddEdge(graph.VertexId(3), graph.VertexId(8), 1)
	_ = h.AddEdge(graph.VertexId(3), graph.VertexId(10), 1)
	_ = h.AddEdge(graph.VertexId(8), graph.VertexId(9), 1)
	_ = h.AddEdge(graph.VertexId(11), graph.VertexId(10), 1)
	_ = h.AddEdge(graph.VertexId(11), graph.VertexId(2), 1)
	_ = h.AddEdge(graph.VertexId(11), graph.VertexId(9), 1)

	s := Sort(h)

	first := int(s.Pop().(graph.VertexId))
	if first != 7 &&
		first != 5 &&
		first != 3 {
		fmt.Print(first, first != 7)
		t.Error()
	}
}
