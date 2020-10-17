/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//import "fmt"
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := new (ListNode)
    tail := result
    var x1 int
    var x2 int
    var carry int
    var last *ListNode
    carry = 0
    for ;(l1 != nil || l2 != nil); {
        if l1 == nil {
            x1 = 0
        } else {
            x1 = l1.Val
        }
        if l2 == nil {
            x2 = 0
        } else {
            x2 = l2.Val
        }
        sum := x1 + x2 + carry
        //fmt.Println(x1, x2 , sum, carry)
        carry = int(sum/ 10)
        //fmt.Println(x1, x2 , sum, carry)
        sum = sum % 10
        //fmt.Println(x1, x2 , sum, carry)
        //fmt.Println()
        last = tail
        tail.Val = sum
        tail.Next = new(ListNode)
        tail = tail.Next
        if l1 == nil{
        }else{
            l1 = l1.Next    
        }
        if l2 == nil{
        }else{
            l2 = l2.Next    
        }
        
    }
    if carry > 0 {
        tail.Val = carry
    } else {
        last.Next = nil
    }
    return result
}
