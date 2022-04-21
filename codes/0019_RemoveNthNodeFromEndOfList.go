package algorithmn

type ListNode struct {
     Val int
     Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    i, left, right := 1, head, head
    for ; right.Next != nil; i++ {
        right = right.Next
        if i > n {
            left = left.Next
        }
    }
    if i <= n {
        for n > 1 {
            left = left.Next
            n--
        }
    }
    if left == right {
        if i == 1 {
            return nil
        } else {
            return head.Next
        }
    }
    if left.Next == nil {
        return right
    } else {
        left.Next = left.Next.Next
    }
    return head
}