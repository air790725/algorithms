package algorithmn

import "math"

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

func connect(root *Node) *Node {
    var prev *Node
    var queue []*Node
    queue = append(queue, root)
    count, level, lastEnd, currentEnd := 1, 1, 1, 1
    for len(queue) > 0 {
        currentEnd = int(math.Pow(2, float64(level)))
        if lastEnd < currentEnd - 1 {
            lastEnd++
        }
        node := queue[0]
        if node != nil && node.Left != nil && node.Right != nil {
            node.Left.Next = node.Right
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
        }
        if prev != nil && count > lastEnd && count < currentEnd {
            prev.Next = node
        } else if count == currentEnd {
            level++
        }
        count++
        prev = node
        queue = queue[1:]
    }
    return root
}