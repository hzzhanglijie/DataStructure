package tree

import "github.com/syndtr/goleveldb/leveldb/errors"

const (
	RED   = true
	BLACK = false
)

// 红黑树
// 节点不是红色就是黑色
// 根节点是黑色
// 所有的空叶子节点都是黑色的
// 如果一个节点是红色，那么它的两个子节点都是黑色
// 从任意一个节点到这个节点后代的每一个空叶子节点的直接路径都要经过相同个数的黑色节点
// 左旋：以P为轴心左旋，N原来的父节点P作为N节点的左孩子，原N节点的左孩子变为P节点的右孩子
// 右旋：以P为轴心左旋，N原来的父节点P作为N节点的左孩子，原N节点的右孩子变为P节点的左孩子
type RBTree struct {
	Root *RBTreeNode
}

type RBTreeNode struct {
	Color  bool
	Value  int
	Left   *RBTreeNode
	Right  *RBTreeNode
	Parent *RBTreeNode
}

func NewRBNode(v int) *RBTreeNode {
	return &RBTreeNode{Color: RED, Value: v}
}

// 获取当前节点祖父节点
func (node *RBTreeNode) Grandfather() *RBTreeNode {
	if node.Parent == nil {
		return nil
	}
	return node.Parent.Parent
}

// 获取当前节点兄弟节点
func (node *RBTreeNode) Sibling() *RBTreeNode {
	if node.Parent == nil {
		return nil
	}
	if node.Parent.Left == node {
		return node.Parent.Right
	} else {
		return node.Parent.Left
	}
}

// 获取当前节点父亲节点的兄弟节点
func (node *RBTreeNode) Uncle() *RBTreeNode {
	if node.Grandfather() == nil {
		return nil
	}
	if node.Grandfather().Left == node.Parent {
		return node.Grandfather().Right
	} else {
		return node.Grandfather().Left
	}
}

// 左旋当前节点
// 当前节点的右孩子为旋转轴
// 当前节点变为旋转轴的左孩子，旋转轴的左孩子变为旋转节点的右孩子
func (node *RBTreeNode) LeftRotate() error {
	if node == nil {
		return errors.New("Current node can't be nil.")
	}
	if node.Right == nil {
		return errors.New("Right rotate can't be nil.")
	}
	// 得到旋转轴
	axle := node.Right
	// 将旋转轴的父节点变为旋转节点的父节点
	if node.Parent != nil {
		axle.Parent = node.Parent
		// 修改旋转节点父节点的孩子节点的指针
		if node.Parent.Left == node {
			node.Parent.Left = axle
		} else {
			node.Parent.Right = axle
		}
	}
	// 得到旋转轴的左孩子
	axleLeft := axle.Left
	// 当前节点变为旋转轴的左孩子
	axle.Left = node
	// 修改旋转节点的父节点
	node.Parent = axle
	// 旋转轴的左孩子变为旋转节点的右孩子
	node.Right = axleLeft

	return nil

}

// 右旋当前节点
// 当前节点的左孩子为旋转轴
// 当前节点变为旋转轴的右孩子，旋转轴的左孩子变为当前节点的右孩子
func (node *RBTreeNode) RightRotate() error {
	if node == nil {
		return errors.New("Current node can't be nil.")
	}
	if node.Left == nil {
		return errors.New("Left rotate can't be nil.")
	}
	// 得到旋转轴
	axle := node.Left
	// 将旋转轴的父节点变为旋转节点的父节点
	if node.Parent != nil {
		axle.Parent = node.Parent
		// 修改旋转节点父节点的孩子节点的指针
		if node.Parent.Left == node {
			node.Parent.Left = axle
		} else {
			node.Parent.Right = axle
		}
	}
	// 得到旋转轴的右孩子
	axleLeft := axle.Right
	// 当前节点变为旋转轴的右孩子
	axle.Right = node
	// 修改旋转节点的父节点
	node.Parent = axle
	// 旋转轴的右孩子变为旋转节点的左孩子
	node.Left = axleLeft

	return nil
}
