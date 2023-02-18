package rewritedatastructure

import "errors"

type Node struct {
	val  interface{}
	next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{
		val:  val,
		next: nil,
	}
}

// Node.insert(val interface{}) insert the val in the tail of the node
func (n *Node) insert(val interface{}) error {
	if n.val == nil && n.next == nil {
		return errors.New("Empty node")
	}
	// Find tail
	temp := n
	for temp.next != nil {
		temp = temp.next
	}
	// Insert
	temp.next = NewNode(val)
	return nil
}

// Node.output() return the values of the node in order
func (n *Node) output() ([]interface{}, error) {
	if n.val == nil && n.next == nil {
		return nil, errors.New("Empty node")
	}
	temp := n
	logs := []interface{}{}
	for temp != nil {
		logs = append(logs, temp.val)
		temp = temp.next
	}
	return logs, nil
}

// deleteIndex delete the indexth position of the node. For example, if index is 0, delete the first element of the node
func (n *Node) deleteIndex(index int) (*Node, error) {
	// empty node
	if n.val == nil && n.next == nil {
		return nil, errors.New("empty node")
	}
	if n.next == nil {
		return nil, errors.New("Too short nodelist to be deleted")
	}
	temp := n
	i := 0
	for i < index {
		if temp.next != nil {
			i += 1
			temp = temp.next
		} else {
			return nil, errors.New("The index is larger than the lenght of nodelist")
		}
	}
	// head
	if temp == n {
		return n.next, nil
	}
	// tail and ordinary
	new_temp := n
	// find the node before temp
	for new_temp.next != temp {
		new_temp = new_temp.next
	}
	if temp.next != nil {
		temp_next := temp.next
		new_temp.next = temp_next
	} else {
		new_temp.next = nil
	}
	return n, nil
}

// reverse reverse the nodelist and return the new head of the nodelist
func (head *Node) reverse() *Node {
	// reverse from the head node
	var pre *Node = nil
	cur := head
	for cur != nil {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}
	return pre
}
