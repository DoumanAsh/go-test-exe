package algo

import "testing"

func TestStack(t *testing.T) {
    child1 := Node{"chil1", []Node{}}
    child2 := Node{"chil2", []Node{}}

    node := Node{"alone", []Node{child1, child2}}

    stack := new(Stack)
    stack.push([]*Node{&node})

    if stack.len() != 1 {
        t.Error("Expected stack size 1 but got", stack.len())
    }

    stack_elem := stack.last()
    child_elem := stack_elem.pop()

    if child_elem.name != node.name {
        t.Errorf("Expected child node with name='%s' but got '%s'", node.name, child_elem.name)
    }

    children_len := len(child_elem.children)
    if (children_len != len(node.children)) {
        t.Error("Top node children are incorrect on stack!")
    }

    stack.push_owned(&child_elem.children)

    if stack.len() != 2 {
        t.Error("Expected stack size 2 but got", stack.len())
    }

    stack_elem = stack.last()
    if len(stack_elem.nodes) != 2 {
        t.Error("Expected to see 2 nodes on stack but got", len(stack_elem.nodes))
    }

    child_elem = stack_elem.pop()
    if child_elem.name != child1.name {
        t.Errorf("Expected child node with name='%s' but got '%s'", child1.name, child_elem.name)
    }

    child_elem = stack_elem.pop()
    if child_elem.name != child2.name {
        t.Errorf("Expected child node with name='%s' but got '%s'", child2.name, child_elem.name)
    }

    child_elem = stack_elem.pop()
    if child_elem != nil {
        t.Error("Children should end, but there is still some  more")
    }

    stack.pop()
    stack.pop()

    if stack.len() != 0 {
        t.Error("Expected stack size 0 but got", stack.len())
    }
}

func TestEmptyChildren(t *testing.T) {
    node := Node{"alone", []Node{}}
    elements := []Node{}

    iter := node.iter()
    for elem := iter(); elem != nil; elem = iter() {
        elements = append(elements, *elem)
    }

    if len(elements) != 1 {
        t.Error("Expected 1 element, but got", len(elements))
    }

    if node.name != elements[0].name {
        t.Error("Expected node with name", node.name, "but got", elements[0].name)
    }
}

func TestOneLevelChildren(t *testing.T) {
    child1 := Node{"chil1", []Node{}}
    child2 := Node{"chil2", []Node{}}

    node := Node{"alone", []Node{child1, child2}}
    elements := []Node{}

    iter := node.iter()
    for elem := iter(); elem != nil; elem = iter() {
        elements = append(elements, *elem)
    }

    if len(elements) != 3 {
        t.Error("Expected 3 element, but got", len(elements))
    }

    if node.name != elements[0].name {
        t.Error("Expected node with name", node.name, "but got", elements[0].name)
    }

    if child1.name != elements[1].name {
        t.Error("Expected node with name", child1.name, "but got", elements[1].name)
    }

    if child2.name != elements[2].name {
        t.Error("Expected node with name", child2.name, "but got", elements[2].name)
    }
}

func TestMultiLevelChildren(t *testing.T) {
    node := Node{"alone", []Node{
        Node{"child1", []Node{
            Node{"1sub1", []Node{
                Node{"1deep1", []Node{}},
            }},
            Node{"1sub2", []Node{}},
            Node{"1sub3", []Node{}},
        }},
        Node{"child2", []Node{
            Node{"2sub1", []Node{}},
            Node{"2sub2", []Node{
                Node{"2deep2", []Node{}},
            }},
        }},
    }}

    elements := []Node{}
    expect_elements := []string{
        "alone",
        "child1",
        "1sub1",
        "1deep1",
        "1sub2",
        "1sub3",
        "child2",
        "2sub1",
        "2sub2",
        "2deep2",
    }

    iter := node.iter()
    for elem := iter(); elem != nil; elem = iter() {
        elements = append(elements, *elem)
    }

    if len(elements) != len(expect_elements) {
        t.Error("Expected %d element, but got", len(expect_elements), len(elements))
    }

    for idx, element := range elements {
        if (element.name != expect_elements[idx]) {
            t.Error("%d: Expected element '%s' but got '%s", idx, element.name, expect_elements[idx])
        }
    }
}
