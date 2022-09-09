package tree

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	Used   bool
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func (node Node) IsEmpty() bool {
	return reflect.DeepEqual(node, Node{})
}

func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	root, rec, e := findRoot(records)
	records = *rec

	if e != nil {
		return nil, e
	}

	// return root, nil

	children, recs := findChildren(records, root.ID)
	if len(recs) > 0 {
		return nil, errors.New("wrong")
	}
	root.Children = children

	fmt.Println("records: ", records)
	fmt.Println("root: ", root)
	return root, nil
}

func findRoot(records []Record) (*Node, *[]Record, error) {
	var root Node
	for idx, record := range records {
		if record.Used {
			continue
		}

		if record.ID == 0 {
			if record.Parent != 0 {
				return nil, nil, errors.New("one root node and has parent")
			}

			if root.IsEmpty() {
				root.ID = record.ID
				record.Used = true
				records = append(records[:idx], records[idx+1:]...)
				return &root, &records, nil
			}
		}
	}

	return nil, nil, errors.New("can't find root node")
}

func findChildren(records []Record, rootId int) ([]*Node, []Record) {
	var res []*Node
	for idx, r := range records {
		if r.Used == false && r.Parent == rootId {
			var node Node
			node.ID = r.ID
			records = append(records[:idx], records[idx+1:]...)
			children, rec := findChildren(records, node.ID)
			node.Children = children
			records = rec

			res = append(res, &node)
		}
	}

	return res, records
}
