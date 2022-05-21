package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := make(map[int]*Node)

	sorted := records[:]
	sort.Slice(sorted, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	root := &Node{
		ID:       sorted[0].ID,
		Children: make([]*Node, 0),
	}

	if root.ID != 0 || sorted[0].Parent != 0 {
		return nil, errors.New("invalid root node")
	}

	nodes[root.ID] = root

	for i, record := range records {
		if i == 0 {
			continue
		}

		if nodes[record.ID] != nil {
			return nil, errors.New("duplicate node")
		}

		if record.ID <= record.Parent {
			return nil, errors.New("invalid node")
		}

		if record.ID-1 != records[i-1].ID {
			return nil, errors.New("non continous")
		}

		node := &Node{
			ID:       record.ID,
			Children: make([]*Node, 0),
		}

		parent := nodes[record.Parent]
		if parent != nil {
			parent.Children = append(parent.Children, node)
		}

		nodes[record.ID] = node
	}

	return root, nil
}
