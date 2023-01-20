package main

import (
	"bufio"
	"fmt"
	"os"
	"serdestres/tree"
	"strconv"
	"strings"
)

func main() {
	//ReadInput
	list := readLineFromTerminal()
	root := deserialize(list)
	fmt.Println(root)

	//Traverse(tree)
	fmt.Println(serialize(root))
}

func readLineFromTerminal() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Tree: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	fmt.Println(text)
	strs := strings.Split(text, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}

	fmt.Println(ary)
	return ary
}

func serialize(root *tree.TreeNode) (serializedTree []int) {
	if root == nil {
		return nil
	}
	serializedTree = append(serializedTree, root.Data)
	if root.Left != nil {
		serialize(root.Left)
	}
	if root.Right != nil {
		serialize(root.Right)
	}
	return serializedTree
}

func deserialize(serializedTree []int) (root *tree.TreeNode) {
	len := len(serializedTree)
	if len == 0 || serializedTree[0] == 0 {
		return nil
	}
	root = tree.NewTreeNode(serializedTree[0])
	root.Left = deserialize(serializedTree[1:])
	root.Right = deserialize(serializedTree[1:])
	return root
}
