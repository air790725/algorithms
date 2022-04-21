package algorithmn

 type ListNode struct {
     Val int
     Next *ListNode
 }
 
 func middleNode(head *ListNode) *ListNode {
    left := head
    for i, right := 0, left; right.Next != nil; {
        right = right.Next
        if i % 2 == 0 {
            left = left.Next
        }
        i++
    }
    return left
}