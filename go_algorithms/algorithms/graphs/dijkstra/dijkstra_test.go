package dijkstra

import (
	"fmt"
	"go_algorithms/data-structures/graph"
	"testing"
)

func TestDijkstra(t *testing.T) {
	h := graph.NewUndirected()

	for i := 0; i < 5; i++ {
		_ = h.AddVertex(graph.VertexId(i))
	}

	_ = h.AddEdge(graph.VertexId(0), graph.VertexId(1), 10)
	_ = h.AddEdge(graph.VertexId(1), graph.VertexId(2), 20)
	_ = h.AddEdge(graph.VertexId(2), graph.VertexId(3), 40)
	_ = h.AddEdge(graph.VertexId(0), graph.VertexId(2), 50)
	_ = h.AddEdge(graph.VertexId(0), graph.VertexId(3), 80)
	_ = h.AddEdge(graph.VertexId(0), graph.VertexId(4), 10)
	_ = h.AddEdge(graph.VertexId(4), graph.VertexId(3), 10)

	prev := ShortestPath(h, graph.VertexId(0))

	if prev[1] != graph.VertexId(0) ||
		prev[2] != graph.VertexId(1) ||
		prev[3] != graph.VertexId(4) ||
		prev[4] != graph.VertexId(0) {

		fmt.Println(prev)
		t.Error()
	}
}
