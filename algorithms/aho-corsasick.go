package algorithms

import (
	"fmt"
)

const charSpace = 256 //ascii
const rootChar = byte('_')

type Node struct {
	children    [charSpace]*Node
	failureLink *Node
	char        byte
	isLeaf      bool
	str         string
}

func printTree(node *Node, level int) {
	// fmt.Printf("%c(%d, %s)", node.char, level, node.str)
	if node.char != rootChar {
		// fmt.Printf("%c(%s)", node.char, node.str)
		fmt.Printf("%c", node.char)
	}
	if node.failureLink != nil {
		fmt.Printf("@%c(%s)", node.failureLink.char, node.failureLink.str)
	}
	if node.isLeaf {
		fmt.Printf(", (leaf)\n")
		return
	}
	for i := 0; i < charSpace; i++ {
		if node.children[i] != nil {
			fmt.Print(" => ")
			printTree(node.children[i], level+1)
		}
	}
}

func buildFailureLink(root *Node, node *Node) {
	node.failureLink = findFailureLink(root, node)
	for i := 0; i < len(node.children); i++ {
		if node.children[i] != nil {
			buildFailureLink(root, node.children[i])
		}
	}
}

func findFailureLink(root *Node, node *Node) *Node {
	if node == root || len(node.str) == 1 {
		return root
	}

	for i := 1; i < len(node.str); i++ {
		suffix := node.str[i:]
		// fmt.Printf("node: %c, suffix: %s\n", node.char, suffix)
		ptr := root
		for j := 0; j < len(suffix); j++ {
			char := suffix[j]
			if ptr.children[char] != nil {
				ptr = ptr.children[char]
				if j == len(suffix)-1 {
					return ptr
				}
			} else {
				break
			}
		}
	}
	return root
}

func addPattern(node *Node, pattern string, i int) *Node {
	if i >= len(pattern) {
		return node
	}
	char := pattern[i]
	if node.children[char] == nil {
		node.children[char] = &Node{[charSpace]*Node{}, nil, char, false, pattern[0 : i+1]}
	}
	return addPattern(node.children[char], pattern, i+1)
}

func buildTree(patterns []string) *Node {
	root := &Node{}
	root.char = rootChar
	for _, pattern := range patterns {
		var leaf *Node = addPattern(root, pattern, 0)
		leaf.isLeaf = true
	}
	return root
}

func AhoCorsasick(text string, patterns []string) (int, int, map[string]int) {
	root := buildTree(patterns)
	buildFailureLink(root, root)
	// printTree(root, 0)

	foundPatternMap := make(map[string]int)
	lenText := len(text)
	ptr := root
	i := 0
	totalFound := 0

	for i < lenText {
		char := text[i]
		if ptr.children[char] == nil {
			if ptr == root && ptr.failureLink == root {
				i++
			}
			ptr = ptr.failureLink
			continue
		}
		ptr = ptr.children[char]
		if ptr.isLeaf {
			_, match := foundPatternMap[ptr.str]
			if !match {
				totalFound++
				charIndex := i - len(ptr.str) + 1
				foundPatternMap[ptr.str] = charIndex
				// fmt.Printf("'%s' found at %d\n", ptr.str, charIndex)
			}
		}
		i++
	}

	return totalFound, lenText, foundPatternMap
}
