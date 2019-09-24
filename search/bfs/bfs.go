package bfs

import (
	"github.com/dbzer0/algorithm/queue"
	"github.com/dbzer0/algorithm/structure/graph"
)

type BFS struct {
	Graph *graph.Graph
	Queue *queue.FIFO
}

func NewBFS(graph *graph.Graph) *BFS {
	return &BFS{Graph: graph, Queue: queue.NewFIFO()}
}

func (bfs *BFS) Find(rootID, findID int64) *graph.Node {
	viewedIDs := []int64{rootID} // массив того, что мы уже посетили

	bfs.Queue.Push(rootID)

	for bfs.Queue.Len() != 0 {
		current := bfs.Queue.Pop()

		node := *bfs.Graph.Nodes[current.(int64)]

		// нашли искомую ноду
		if current == findID {
			return &node
		}

		viewedIDs = append(viewedIDs, current.(int64))

		for _, edge := range node.Edges {
			// если соседняя нода искомая, то возвращаем ее
			if edge.TargetID.ID == findID {
				return edge.TargetID
			}

			var contains bool
			contains = false
			// проверяем, был ли этот id в очереди ранее
			for _, id := range viewedIDs {
				if id == edge.TargetID.ID {
					contains = true
				}
			}

			// если id не было в очереди, то добавляем id
			// следующей ноды в очередь
			if !contains {
				bfs.Queue.Push(edge.TargetID.ID)
			}
		}
	}

	return nil
}
