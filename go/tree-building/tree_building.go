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

	// Sort records by ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))
	for i, record := range records {
		if record.ID != i {
			return nil, errors.New("non-continuous IDs")
		}
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("root node has a parent")
		}
		if record.ID != 0 && record.ID <= record.Parent {
			return nil, errors.New("non-root node has ID less than or equal to parent ID")
		}
		if _, exists := nodes[record.ID]; exists {
			return nil, errors.New("duplicate node ID")
		}

		node := &Node{ID: record.ID}
		nodes[record.ID] = node

		if record.ID != 0 {
			parentNode, exists := nodes[record.Parent]
			if !exists {
				return nil, errors.New("parent node does not exist")
			}
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	return nodes[0], nil
}
