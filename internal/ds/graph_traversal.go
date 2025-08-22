package ds

func (g Graph) BFSTraversal(startVertex Vertex, activeEdges EdgeSet) []Vertex {
	q := NewQueue[Vertex]()
	q.Enqueue(startVertex)
	visited := NewSet[Vertex]()
	for !q.IsEmpty() {
		current := q.Dequeue()
		if visited.Contains(current) {
			continue
		}
		visited.Add(current)
		for _, neighbor := range g.activeNeighbors(current, activeEdges) {
			if !visited.Contains(neighbor) {
				q.Enqueue(neighbor)
			}
		}
	}
	return visited.Items()
}

func (g Graph) activeNeighbors(vertex Vertex, activeEdges EdgeSet) []Vertex {
	neighbors := NewSet[Vertex]()
	for _, edge := range g.connectedEdges[vertex] {
		if activeEdges == nil || activeEdges.Contains(edge) {
			for _, v := range edge {
				neighbors.Add(v)
			}
		}
	}
	neighbors.Delete(vertex)
	return neighbors.Items()
}
