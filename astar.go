package astar

import (
	"iter"
	"slices"

	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/heap"
)

// node is a node in the search space.
type node struct {
	xy     vector2.Int // Position
	parent *node       // Parent node
	g      int         // Cost from start node
	h      int         // Heuristic cost to end node
	f      int         // F = G + H
	weight int         // Weight of the node (-1 = impassable)
	open   bool        // In open list
	closed bool        // In closed list
}

func (n1 *node) Cmp(n2 *node) int {
	return n1.f - n2.f
}

// Pathfinder is a simple A* pathfinding algorithm implementation.
type Pathfinder struct {
	size        vector2.Int
	searchSpace []node
	//weights       Grid[int]
	//options       option
	//heuristic     heuristicFunc
	//getSuccessors getSuccessorsFunc
}

func NewPathfinder(size vector2.Int) Pathfinder {
	return Pathfinder{
		size:        size,
		searchSpace: make([]node, size.Product()),
	}
}

func (p *Pathfinder) get(pos vector2.Int) *node {
	return &p.searchSpace[pos.X+pos.Y*p.size.X]
}

func (p *Pathfinder) getSuccessors(n *node) iter.Seq[*node] {
	return func(yield func(*node) bool) {
		// ab := b.Sub(a)
		// ab_length := ab.LengthF()
		// direction := ab.Normalized()
		// distance := mathex.Clamp(ab_length/(float32(count)-1), mind, maxd)
		//
		// _, dv := justify.Justyfy(min(distance*(float32(count)-1), ab_length), ab_length)
		// position := a.Add(direction.ScaleF(dv))
		// delta := direction.ScaleF(distance)
		// for _, e := range s {
		// if !yield(position, e) {
		// return
		// }
		//
		// position = position.Add(delta)
		// }
	}
}

// Find returns a path from start to end. If no path is found, an empty slice
// is returned.
func (p *Pathfinder) Find(startPos, endPos vector2.Int) []vector2.Int {
	clear(p.searchSpace)
	open := &heap.Heap[*node, heap.Min]{} // prioritised queue of f

	start := p.get(startPos)
	start.f = 0
	start.open = true

	heap.PushOrderable(open, start)

	for {
		q, ok := heap.PopOrderable(open)
		if !ok {
			break
		}

		for n := range p.getSuccessors(q) {
			// not traversable
			if n.weight < 0 {
				continue
			}

			n.parent = q

			g := q.g + 0 //p.heuristic(qPos, succPos)
			//if p.options.punishChangeDirection {
			//	g += punishChangeDirection(q, succPos, endPos)
			//}

			n.g = g
			n.h = 0 //p.heuristic(succPos, endPos)
			n.f = n.g + n.h
			n.open = true

			// found
			if n.xy == endPos {
				path := make([]vector2.Int, 0)
				for n != nil {
					path = append(path, n.xy)
					n = n.parent
				}
				slices.Reverse(path)
				return path
			}

			// 		// check if more optimal path to successor was already encountered
			// 		existingSuccessor := searchSpace.Get(succPos)
			// 		if existingSuccessor.open && existingSuccessor.f < successor.f {
			// 			continue
			// 		}

			// 		if existingSuccessor.closed && existingSuccessor.f < successor.f {
			// 			continue
			// 		}

			// 		searchSpace.Set(succPos, successor)
			heap.PushOrderable(open, n)
		}
		q.closed = true
	}

	// not found
	return []vector2.Int{}
}
