package dfs

import (
	"go_algorithms/data-structures/graph"
	"go_algorithms/data-structures/stack"
)

// 深度优先搜索
func UndirectedDfs(g *graph.UnGraph, v graph.VertexId, fn func(graph.VertexId)) {
	s := stack.New()
	s.Push(v)
	visited := make(map[graph.VertexId]bool)

	for s.Len() > 0 {
		v := s.Pop().(graph.VertexId)

		if _, ok := visited[v]; !ok {
			visited[v] = true
			fn(v)
			neighbours := g.GetNeighbours(v).VerticesIter()
			for neighbour := range neighbours {
				s.Push(neighbour)
			}
		}
	}
}

func DirectedDfs(g *graph.DirGraph, v graph.VertexId, fn func(graph.VertexId)) {
	s := stack.New()
	s.Push(v)
	visited := make(map[graph.VertexId]bool)

	for s.Len() > 0 {
		v = s.Pop().(graph.VertexId)

		if _, ok := visited[v]; !ok {
			visited[v] = true
			fn(v)
			neighbours := g.GetSuccessors(v).VerticesIter()
			for neighbour := range neighbours {
				s.Push(neighbour)
			}
		}
	}
}
