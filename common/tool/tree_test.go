package tool

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	//nodeList := []Node{
	//	{ID: 1, ParentID: 0, Value: "Node A"},
	//	{ID: 2, ParentID: 1, Value: "Node B"},
	//	{ID: 3, ParentID: 1, Value: "Node C"},
	//	{ID: 4, ParentID: 0, Value: "Node D"},
	//	{ID: 5, ParentID: 4, Value: "Node E"},
	//}
	//
	//tree := BuildTree(nodeList)
	//rootNodes := tree.GetRootNodes()
	//
	//fmt.Println("Root nodes:")
	//for _, node := range rootNodes {
	//	fmt.Printf("%d: %s\n", node.ID, node.Value)
	//}
	//
	//fmt.Println("Tree structure:", tree)

	nodelist := []TreeNode{
		{Id: "1", ParentId: "0", Name: "Node A"},
		{Id: "2", ParentId: "1", Name: "Node B"},
		{Id: "3", ParentId: "2", Name: "Node c"},
		{Id: "4", ParentId: "0", Name: "Node c"},
		{Id: "5", ParentId: "4", Name: "Node c"},
	}
	trees := ArrayToTrees(nodelist, "0")
	fmt.Println(trees)
}

// Node 定义节点结构
type Node struct {
	ID       int
	ParentID int
	Value    string
	Children []*Node
}

// Tree 定义树结构
type Tree struct {
	Nodes map[int]*Node
}

// BuildTree 从节点列表构建树
func BuildTree(nodeList []Node) *Tree {
	tree := &Tree{Nodes: make(map[int]*Node)}

	for _, node := range nodeList {
		newNode := &Node{
			ID:       node.ID,
			ParentID: node.ParentID,
			Value:    node.Value,
		}
		tree.Nodes[node.ID] = newNode

		// 如果父节点已存在于树中，则将其添加为子节点
		if parent, ok := tree.Nodes[node.ParentID]; ok {
			parent.Children = append(parent.Children, newNode)
		}
	}

	return tree
}

// GetRootNodes 获取顶级节点（parentID为0）
func (t *Tree) GetRootNodes() []*Node {
	var rootNodes []*Node

	for _, node := range t.Nodes {
		if node.ParentID == 0 {
			rootNodes = append(rootNodes, node)
		}
	}

	return rootNodes
}

func ArrayToTrees(arr []TreeNode, pid string) []TreeNode {
	var newArr []TreeNode
	for _, item := range arr {
		if item.ParentId == pid {
			item.Children = ArrayToTrees(arr, item.Id)
			newArr = append(newArr, item)
		}
	}
	return newArr
}

type TreeNode struct {
	Id       string
	ParentId string
	Name     string
	Children []TreeNode
}
