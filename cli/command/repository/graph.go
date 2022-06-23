package repository

// Graph : represents a Graph
type Graph struct {
	Nodes []*GraphNode
}

// GraphNode : represents a Graph node
type GraphNode struct {
	id       int
	WalletId string
	value int
	Edges map[int]int
}

// New : returns a new instance of a Graph
func New() *Graph {
	return &Graph{
		Nodes: []*GraphNode{},
	}
}

// AddNode : adds a new node to the Graph
func (g *Graph) AddNode(walletId string, value int) (id int) {
	id = len(g.Nodes)
	g.Nodes = append(g.Nodes, &GraphNode{
		id:       id,
		WalletId: walletId,
		value:    value,
		Edges:    make(map[int]int),
	})
	return
}

// AddEdge : adds a directional edge together with a weight
func (g *Graph) AddEdge(n1, n2 int, w int) {
	g.Nodes[n1].Edges[n2] = w
}

// Neighbors : returns a list of node IDs that are linked to this node
func (g *Graph) Neighbors(id int) []int {
	var neighbors []int
	for _, node := range g.Nodes {
		for edge := range node.Edges {
			if node.id == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.id)
			}
		}
	}
	return neighbors
}

// Edges : returns a list of edges with weights
func (g *Graph) Edges() [][3]int {
	edges := make([][3]int, 0, len(g.Nodes))
	for i := range g.Nodes {
		for k, v := range g.Nodes[i].Edges {
			edges = append(edges, [3]int{i, k, v})
		}
	}
	return edges
}
