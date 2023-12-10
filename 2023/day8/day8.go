package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/felixz92/aoc/2023/input"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("No input file")
	}

	i := input.FromFile(args[0])

	res := Day8(i)
	fmt.Println(res)

	res2 := Day8Part2(i)
	fmt.Println(res2)
}

type Direction rune

type Node struct {
	Name  string
	Left  string
	Right string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v = (left = %v, right = %v)", n.Name, n.Left, n.Right)
}

func Day8(lines []string) int {
	instructions := []Direction(lines[0])

	nodes := ParseNodes(lines[2:])

	return Walk(nodes, "AAA", instructions, func(n Node) bool {
		return n.Name == "ZZZ"
	})
}

func Walk(nodes map[string]Node, startingNode string, instructions []Direction, endCond func(Node) bool) int {
	steps := 0
	current := nodes[startingNode]
	for i := 0; i < len(instructions); i++ {
		if endCond(current) {
			break
		}

		direction := instructions[i]
		if direction == 'L' {
			current = nodes[current.Left]
		} else {
			current = nodes[current.Right]
		}

		steps++
		if i == len(instructions)-1 {
			i = -1
		}

	}
	return steps
}

func StartingNodes(nodes map[string]Node) []string {
	startingNodes := make([]string, 0)
	for _, node := range nodes {
		if node.Name[2] == 'A' {
			startingNodes = append(startingNodes, node.Name)
		}
	}
	return startingNodes
}

func AllEnding(nodes []string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func Day8Part2(lines []string) int {
	instructions := []Direction(lines[0])
	nodes := ParseNodes(lines[2:])
	startingNodes := StartingNodes(nodes)

	steps := make([]int, len(startingNodes))
	for i := range startingNodes {
		steps[i] = Walk(nodes, startingNodes[i], instructions, func(n Node) bool {
			return n.Name[2] == 'Z'
		})
	}

	res := 1
	for _, step := range steps {
		res = lcm(res, step)
	}

	return res
}

func ParseNodes(lines []string) map[string]Node {
	nodes := make(map[string]Node)
	for _, line := range lines {
		parts := strings.Split(line, "=")
		name := strings.TrimSpace(parts[0])
		pointers := strings.Split(parts[1], ",")
		left := strings.TrimPrefix(strings.TrimSpace(pointers[0]), "(")
		right := strings.TrimSuffix(strings.TrimSpace(pointers[1]), ")")
		nodes[name] = Node{
			Name:  name,
			Left:  left,
			Right: right,
		}
	}

	return nodes
}
