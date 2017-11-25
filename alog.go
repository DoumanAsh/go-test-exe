package algo

import (
    "fmt"
    "bytes"
)

type Node struct {
    name string
    children []Node
}

//iter produces pointers to element and walks through all its children.
//Iteration is emulated through channels.
//
//Algorithm:
//1. Places element on stack with children index equal to 0
//2. Iterate through children using index.
//3. If child has own children, then place the child onto stack to repeat step 1.
//4. If there is no more children, then pop last element of stack and repeat step 2.
func (node *Node) iter() chan *Node {
    ch := make(chan *Node)

    type Stack struct {
        node *Node
        idx int
    }

    go func(ch chan *Node) {
        ch <- node

        stack := []Stack {
            Stack {
                node: node,
                idx: 0,
            },
        }

        stack_len := len(stack)
        for stack_len > 0 {
            //Reference to last stack element.
            stack_elem := &stack[stack_len-1]
            //stack's children index.
            children_idx := stack_elem.idx
            children_len := len(stack_elem.node.children)

            if (children_len > children_idx) {
                //Found child, increment index.
                stack_elem.idx = children_idx + 1
                //Reference to stack element's child.
                child_elem := &stack_elem.node.children[children_idx]

                ch <- child_elem
                if len(child_elem.children) > 0 {
                    //child has own children so place it on stack.
                    stack = append(stack, Stack {
                        node: child_elem,
                        idx: 0,
                    })
                    stack_len = len(stack)
                }
            } else {
                //No more children, pop.
                stack = stack[:stack_len-1]
                stack_len = len(stack)
            }
        }

        close(ch)

	}(ch)

	return ch
}

//Returns string with all elements name separated by \n
func (node *Node) inspect() string {
    var buffer bytes.Buffer

    for elem := range node.iter() {
        buffer.WriteString(elem.name)
        buffer.WriteString("\n")
    }

    return buffer.String()
}

//Prints all elements as in inspect method
func (node *Node) debug_print() {
    fmt.Println(node.inspect())
}
