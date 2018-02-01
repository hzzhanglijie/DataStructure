package tree

import "testing"

var bst = BSTree{}

func init() {
	bst.Insert(53)
	bst.Insert(47)
	bst.Insert(58)
	bst.Insert(26)
	bst.Insert(32)
	bst.Insert(21)
	bst.Insert(34)
	bst.Insert(55)
	bst.Insert(14)
}

func TestBSTree_Contains(t *testing.T) {

	nodes := make([]*BSTreeNode, 0)
	nodes = append(nodes, bst.Contains(53))
	nodes = append(nodes, bst.Contains(47))
	nodes = append(nodes, bst.Contains(58))
	nodes = append(nodes, bst.Contains(26))
	nodes = append(nodes, bst.Contains(32))
	nodes = append(nodes, bst.Contains(21))
	nodes = append(nodes, bst.Contains(34))
	nodes = append(nodes, bst.Contains(55))
	nodes = append(nodes, bst.Contains(14))

	for _, v := range nodes {
		if v == nil {
			t.Fail()
		}
	}

}

func TestBSTree_InOrder(t *testing.T) {
	array := []int{14, 21, 26, 32, 34, 47, 53, 55, 58}
	order := bst.InOrder()

	for i := 0; i < 9; i++ {
		if array[i] != order[i] {
			t.Fail()
		}
	}
}

func TestBSTree_Max(t *testing.T) {
	max := bst.Max()
	if max != 58 {
		t.Fail()
	}
}

func TestBSTree_Min(t *testing.T) {
	min := bst.Min()
	if min != 14 {
		t.Fail()
	}
}

func TestBSTree_Insert(t *testing.T) {
	bst.Insert(60)
	array := []int{14, 21, 26, 32, 34, 47, 53, 55, 58, 60}
	order := bst.InOrder()

	for i := 0; i < 10; i++ {
		if array[i] != order[i] {
			t.Fail()
		}
	}
}

func TestBSTree_Remove(t *testing.T) {
	bst.Remove(32)

	array := []int{14, 21, 26, 34, 47, 53, 55, 58}
	order := bst.InOrder()

	for i := 0; i < 8; i++ {
		if array[i] != order[i] {
			t.Fail()
		}
	}

}
