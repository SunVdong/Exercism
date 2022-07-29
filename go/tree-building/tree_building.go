package tree

import (
	"fmt"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	//sort.Slice(records, func(i, j int) bool {
	//	return records[i].ID < records[j].ID
	//})

	root := Node{ID: 0, Children: nil}

	res,_ := root.findChildren(records)

	fmt.Println(res)

	return &res, nil
}

func (r Node) findChildren(records []Record) (Node, error) {
	var res []*Node
	for _, item := range records {
		if r.ID == item.Parent {
			node := Node{ID: item.ID}
			newNode, _ := node.findChildren(records)
			res = append(res, &newNode)
		}
	}
	r.Children = res

	return r, nil
}
