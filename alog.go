package algo

import (
    "fmt"
    "bytes"
)

type Node struct {
    name string
    children []Node
}

//Holds elements of Stack.
type StackElement struct {
    nodes []*Node
}

func (self *StackElement) pop() *Node {
    size := len(self.nodes)
    if (size > 0) {
        size--
        result := self.nodes[size]
        self.nodes = self.nodes[:size]
        return result
    }
    return nil
}

type Stack struct {
    inner []StackElement
}

func (stack *Stack) last() *StackElement {
    return &stack.inner[len(stack.inner)-1]
}

func (stack *Stack) len() int {
    return len(stack.inner)
}

func (stack *Stack) push(nodes []*Node) {
    stack.inner = append(stack.inner, StackElement {
        nodes: nodes,
    })
}

//Specialized push which transforms slice with owned elements
//to slice of references
func (stack *Stack) push_owned(nodes *[]Node) {
    size := len(*nodes)
    children := make([]*Node, size)
    children_idx := size - 1
    for idx := 0; idx < size; idx++ {
        children[children_idx] = &(*nodes)[idx]
        children_idx--
    }
    stack.push(children)
}

func (stack *Stack) pop() *StackElement {
    size := stack.len()
    if (size > 0) {
        size--
        result := stack.inner[size]
        stack.inner = stack.inner[:size]
        return &result
    }
    return nil
}

func (node *Node) iter() func() *Node {
    stack := new(Stack)
    stack.push([]*Node{node})

    return func() *Node {
        for stack.len() > 0 {
            stack_elem := stack.last()
            child_elem := stack_elem.pop()

            if (child_elem == nil) {
                stack.pop()
            } else {
                if (len(child_elem.children) > 0) {
                    stack.push_owned(&child_elem.children)
                }
                return child_elem
            }
        }

        return nil;
    }
}

//Returns string with all elements name separated by \n
func (node *Node) inspect() string {
    var buffer bytes.Buffer

    iter := node.iter()
    for elem := iter(); elem != nil; elem = iter() {
        buffer.WriteString(elem.name)
        buffer.WriteString("\n")
    }

    return buffer.String()
}

//Prints all elements as in inspect method
func (node *Node) debug_print() {
    fmt.Println(node.inspect())
}
