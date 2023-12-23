package main

import (
	"time"
	"sync"
	"fmt"
	// "math/rand"
)

// type Node struct {
// 	Data interface{}
// 	// Sleep time.Duration
// 	Left *Node
// 	Right *Node
// }

// func NewNode(data interface{}) *Node {

// 	node := new(Node)

// 	node.Data = data
// 	node.Left = nil
// 	node.Right = nil

// 	// rand.Seed(time.Now().UTC().UnixNano())
// 	// duration := int64(rand.Intn(100))
// 	// node.Sleep = time.Duration(duration) * time.Microsecond

// 	return node
// }

func (n *Node) ProcessNode() {

	var arr []int

	for i := 0; i < 10000; i++ {
		// time.Sleep(n.Sleep)
		arr = append(arr, i)
	}
    
	// fmt.Printf("Node %v done\n", n.Value)
}

func (n *Node) DFS() {

	if n == nil {
		return
	}

	n.Left.DFS()
	n.ProcessNode()
	n.Right.DFS()
}


func (n *Node) DFSParallel() {

	defer wg.Done()

	if n == nil {
		return
	}

	wg.Add(1)
	go n.Left.DFSParallel()

	wg.Add(1)
	go n.ProcessNodeParallel()

	wg.Add(1)
	go n.Right.DFSParallel()
}

func (n *Node) ProcessNodeParallel() {

	defer wg.Done()

	var arr []int
    
	for i := 0; i < 10000; i++ {
		// time.Sleep(n.Sleep)
		arr = append(arr, i)
	}
    
	// fmt.Printf("Node %v done\n", n.Value)
}

type Node struct {
    Value int
    Left  *Node
    Right *Node
	
}

func NewNode(value int) *Node {
    return &Node{Value: value}
}

func BuildTree(node *Node, n int) {
    if n == 0 {
        return
    }
	// fmt.Println("начинаем строить")
    node.Left = NewNode(2 * node.Value)
    node.Right = NewNode(2*node.Value + 1)
	// fmt.Println("строим")
    BuildTree(node.Left, n-1)
    BuildTree(node.Right, n-1)
	// fmt.Println("все")
}



var wg sync.WaitGroup

func main() {

	root := NewNode(1)
	BuildTree(root, 15)
	fmt.Println("Построили дерево")


	// Sequential
	start := time.Now()
	root.DFS()
	fmt.Printf("\nTime elapsed: %v\n\n", time.Since(start))

	// Parallel
	start = time.Now()
	wg.Add(1)
	go root.DFSParallel()
	wg.Wait()
	fmt.Printf("\nTime elapsed parallel: %v\n", time.Since(start))

}