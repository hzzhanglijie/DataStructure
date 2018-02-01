package tree

// 二叉排序树（Binary Sort Tree）
// 若它的左子树不空，则左子树上所有节点的值均小于它的根节点的值
// 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值
// 它的左右子树也分别为二叉排序树
type BSTree struct {
	Root *BSTreeNode
}

type BSTreeNode struct {
	Value  int
	Left   *BSTreeNode
	Right  *BSTreeNode
	Parent *BSTreeNode
}

// 查找节点
func (tree *BSTree) Contains(v int) *BSTreeNode {
	node := tree.Root
	for {
		if node == nil {
			return nil
		} else if node.Value < v {
			node = node.Right
		} else if node.Value > v {
			node = node.Left
		} else {
			return node
		}
	}
}

// 插入节点
func (tree *BSTree) Insert(v int) {
	if tree.Root == nil {
		tree.Root = &BSTreeNode{v, nil, nil, nil}
		return
	}
	tree.Root.insert(v)
}

func (node *BSTreeNode) insert(v int) {
	if v < node.Value {
		if node.Left != nil {
			node.Left.insert(v)
		} else {
			node.Left = &BSTreeNode{v, nil, nil, node}
		}
	} else {
		if node.Right != nil {
			node.Right.insert(v)
		} else {
			node.Right = &BSTreeNode{v, nil, nil, node}
		}
	}
}

// 中序遍历
func (tree *BSTree) InOrder() []int {
	order := make([]int, 0)
	tree.Root.inOrder(&order)
	return order
}

func (node *BSTreeNode) inOrder(order *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		node.Left.inOrder(order)
	}
	*order = append(*order, node.Value)
	if node.Right != nil {
		node.Right.inOrder(order)
	}
}

// 取最小值
func (tree *BSTree) Min() int {
	node := tree.Root
	for {
		if node.Left != nil {
			node = node.Left
		} else {
			return node.Value
		}
	}
}

// 取最大值
func (tree *BSTree) Max() int {
	node := tree.Root
	for {
		if node.Right != nil {
			node = node.Right
		} else {
			return node.Value
		}
	}
}

// 删除节点
// 若删除节点为叶子结点，修改其父节点的子节点为空即可
// 若删除节点只有左子树或者只有右子树，此时只要令要删除节点的父节点指向左子树或右子树即可
// 若删除节点既有左子树又有右子树，就用删除节点的左子树的最大节点取代删除节点
func (tree *BSTree) Remove(v int) {
	if tree.Root == nil {
		return
	}
	node := tree.Contains(v)
	if node == nil {
		return
	}

	removeNode := func(node *BSTreeNode) {
		// 左右子树都为空
		if node.Left == nil && node.Right == nil {
			if node.Parent.Right == node {
				node.Parent.Right = nil
			} else {
				node.Parent.Left = nil
			}
			node.Parent = nil
			// 左子树为空，右子树不为空
		} else if node.Left == nil && node.Right != nil {
			if node.Parent.Left == node {
				node.Parent.Left = node.Right
			} else {
				node.Parent.Right = node.Right
			}
			node.Right.Parent = node.Parent
			// 左子树不为空，右子树为空
		} else if node.Left != nil && node.Right == nil {
			if node.Parent.Left == node {
				node.Parent.Left = node.Left
			} else {
				node.Parent.Right = node.Left
			}
			node.Left.Parent = node.Parent
			// 左右子树都不为空
		} else {
			// 找到左子树的最大值
			mln := node.Left
			for {
				if mln.Right != nil {
					mln = mln.Right
					continue
				}

				// 删除节点的左子树是一个叶子结点
				if mln.Parent.Left == mln {
					mln.Parent.Left = nil
					// 删除节点的左子树非叶子结点
				} else {
					mln.Parent.Right = nil
				}

				if node.Parent.Left == node {
					node.Parent.Left = mln
				} else {
					node.Parent.Right = mln
				}
				mln.Parent = node.Parent
				mln.Left = node.Left
				mln.Right = node.Right
				node.Parent = nil
				node.Left = nil
				node.Right = nil
				break
			}
		}
	}

	removeNode(node)

}
