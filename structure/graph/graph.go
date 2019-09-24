package graph

import (
	"sync"
)

type Graph struct {
	// количество нод
	NodeCount int64

	// все ноды графа
	Nodes map[int64]*Node

	l sync.Mutex
}

type Node struct {
	// значение, хранимое нодой
	Value interface{}

	// идентификатор ноды
	ID int64

	Edges []*Edge
}

type Edge struct {
	SourceID *Node
	TargetID *Node
	Weight   float64
}

func NewGraph() *Graph {
	return &Graph{
		NodeCount: 0,
		Nodes:     make(map[int64]*Node),
	}
}

func (g *Graph) NewNode(id int64, value interface{}) *Node {
	g.l.Lock()
	defer g.l.Unlock()

	g.NodeCount++
	g.Nodes[id] = &Node{ID: id, Value: value}
	return g.Nodes[id]
}

func (g *Graph) NewEdge(sourceID, targetID int64, weight float64) *Edge {
	g.l.Lock()
	defer g.l.Unlock()

	// проверка существования нод
	source, ok := g.Nodes[sourceID]
	if !ok {
		return nil
	}
	target, ok := g.Nodes[targetID]
	if !ok {
		return nil
	}

	e  := &Edge{SourceID: source, TargetID: target, Weight: weight}
	te := &Edge{SourceID: target, TargetID: source, Weight: weight}

	source.Edges = append(source.Edges, e)
	target.Edges = append(target.Edges, te)

	return e
}
