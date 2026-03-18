package repository

import (
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	g := New()
	id1 := g.AddNode("wallet1", 100)
	if id1 != 0 {
		t.Errorf("Expected id 0, got %d", id1)
	}
	if len(g.Nodes) != 1 {
		t.Errorf("Expected 1 node, got %d", len(g.Nodes))
	}
	if g.Nodes[0].WalletId != "wallet1" {
		t.Errorf("Expected wallet1, got %s", g.Nodes[0].WalletId)
	}
	if g.Nodes[0].value != 100 {
		t.Errorf("Expected value 100, got %d", g.Nodes[0].value)
	}

	id2 := g.AddNode("wallet2", 200)
	if id2 != 1 {
		t.Errorf("Expected id 1, got %d", id2)
	}
}

func TestGraph_AddEdge(t *testing.T) {
	g := New()
	n1 := g.AddNode("wallet1", 100)
	n2 := g.AddNode("wallet2", 200)
	g.AddEdge(n1, n2, 10)

	if len(g.Nodes[n1].Edges) != 1 {
		t.Errorf("Expected 1 edge, got %d", len(g.Nodes[n1].Edges))
	}
	if g.Nodes[n1].Edges[n2] != 10 {
		t.Errorf("Expected weight 10, got %d", g.Nodes[n1].Edges[n2])
	}
}

func TestGraph_Neighbors(t *testing.T) {
	g := New()
	n0 := g.AddNode("w0", 0)
	n1 := g.AddNode("w1", 1)
	n2 := g.AddNode("w2", 2)

	g.AddEdge(n0, n1, 1)
	g.AddEdge(n0, n2, 1)

	neighbors := g.Neighbors(n0)
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors, got %d", len(neighbors))
	}

	// Neighbors should include n1 and n2
	foundN1 := false
	foundN2 := false
	for _, n := range neighbors {
		if n == n1 {
			foundN1 = true
		}
		if n == n2 {
			foundN2 = true
		}
	}
	if !foundN1 || !foundN2 {
		t.Errorf("Neighbors did not contain expected nodes")
	}
}
