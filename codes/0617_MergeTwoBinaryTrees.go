package algorithmn

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }
    merge(root1, root2)
    return root1
}

func merge(node1 *TreeNode, node2 *TreeNode) {
    node1.Val += node2.Val
    if node1.Left != nil && node2.Left != nil {
        merge(node1.Left, node2.Left)
    } else if node1.Left == nil {
        node1.Left = node2.Left
    }
    if node1.Right != nil && node2.Right != nil {
        merge(node1.Right, node2.Right)
    } else if node1.Right == nil {
        node1.Right = node2.Right
    }
}