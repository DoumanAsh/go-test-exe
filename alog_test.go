package algo

import "testing"

func TestEmptyChildren(t *testing.T) {
    node := Node{"alone", []Node{}}
    elements := []Node{}

    for elem := range node.iter() {
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

    for elem := range node.iter() {
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
    deep1_child1 := Node{"1deep1", []Node{}}
    child1_children1 := Node{"1sub1", []Node{deep1_child1}}
    child1_children2 := Node{"1sub2", []Node{}}
    child1_children3 := Node{"1sub3", []Node{}}
    child2_children1 := Node{"2sub1", []Node{}}
    deep2_child2 := Node{"2deep2", []Node{}}
    child2_children2 := Node{"2sub2", []Node{deep2_child2}}

    child1 := Node{"chil1", []Node{
        child1_children1,
        child1_children2,
        child1_children3,
    }}
    child2 := Node{"chil2", []Node{
        child2_children1,
        child2_children2,
    }}

    node := Node{"alone", []Node{child1, child2}}
    elements := []Node{}
    expect_elements := []Node{
        node,
        child1,
        child1_children1,
        deep1_child1,
        child1_children2,
        child1_children3,
        child2,
        child2_children1,
        child2_children2,
        deep2_child2,
    }

    for elem := range node.iter() {
        elements = append(elements, *elem)
    }

    if len(elements) != len(expect_elements) {
        t.Error("Expected %d element, but got", len(expect_elements), len(elements))
    }

    for idx, element := range elements {
        if (element.name != expect_elements[idx].name) {
            t.Error("%d: Expected element '%s' but got '%s", idx, element.name, expect_elements[idx].name)
        }
    }
}
