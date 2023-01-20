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
	//NewTree(list)
	tree.NewTree(list)
	//Traverse(tree)
}

func readLineFromTerminal() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Tree: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	strs := strings.Split(text, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}

	fmt.Println(ary)
	return ary
}
