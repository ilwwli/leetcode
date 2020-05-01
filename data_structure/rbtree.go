package data_structure

import (
	"fmt"
	"strconv"
)

type RBTreeNode struct {
	Black       bool
	Left, Right *RBTreeNode
	P           *RBTreeNode
	Val         int
}

type RBTree struct {
	Root *RBTreeNode
}

func (t *RBTree) Insert(val int) {
	z := &RBTreeNode{Val: val}
	t.insert(z)
}

func (t *RBTree) Delete(val int) {
	x := t.Root
	for x != nil {
		if val == x.Val {
			break
		} else if val < x.Val {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	if x == nil {
		return
	}
	t.delete(x)
}

func (t *RBTree) leftRotate(x *RBTreeNode) {
	y := x.Right
	if y == nil {
		panic("leftRotate with nil right child")
	}

	x.Right = y.Left // handoff child
	if y.Left != nil {
		y.Left.P = x
	}
	y.Left = x

	if t.Root == x {
		t.Root = y
	} else if x == x.P.Left {
		x.P.Left = y
	} else {
		x.P.Right = y
	}

	y.P = x.P // link x'P to y
	x.P = y
}

func (t *RBTree) rightRotate(y *RBTreeNode) {
	x := y.Left
	if x == nil {
		panic("rightRotate with nil left child")
	}

	y.Left = x.Right // handoff child
	if x.Right != nil {
		x.Right.P = y
	}
	x.Right = y

	if t.Root == y {
		t.Root = x
	} else if y == y.P.Left {
		y.P.Left = x
	} else {
		y.P.Right = x
	}

	x.P = y.P // link x'P to y
	y.P = x
}

func (t *RBTree) insert(z *RBTreeNode) {
	y, x := (*RBTreeNode)(nil), t.Root
	for x != nil {
		y = x
		if z.Val < x.Val {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.P = y
	if y == nil {
		t.Root = z
	} else if z.Val < y.Val {
		y.Left = z
	} else {
		y.Right = z
	}
	z.Left, z.Right = nil, nil
	z.Black = false
	t.insertFixup(z)
}

func (t *RBTree) insertFixup(z *RBTreeNode) {
	// there are 3 loop invariants here
	// 1. z.Black must be false
	// 2. If z.P is root, then z.P.Black must be true
	// 3. At most only one property of RBTree can be broken, which should be one of Prop.2 or Prop.4
	for z != t.Root && z.P.Black == false {
		if z.P == z.P.P.Left { // z.P.P must exist
			y := z.P.P.Right                  // y is uncle of z
			if y != nil && y.Black == false { // y share the same color with z
				y.Black = true
				z.P.Black = true
				z.P.P.Black = false
				z = z.P.P
				continue
			}
			if z == z.P.Right {
				z = z.P
				t.leftRotate(z) // turn to case(z if left child to z.P which is left child to z.P.P)
			}
			z.P.Black = true
			z.P.P.Black = false
			t.rightRotate(z.P.P)
		} else {
			y := z.P.P.Left                   // y is uncle of z
			if y != nil && y.Black == false { // y share the same color with z
				y.Black = true
				z.P.Black = true
				z.P.P.Black = false
				z = z.P.P
				continue
			}
			if z == z.P.Left {
				z = z.P
				t.rightRotate(z) // turn to case(z if right child to z.P which is right child to z.P.P)
			}
			z.P.Black = true
			z.P.P.Black = false
			t.leftRotate(z.P.P)
		}
	}
	t.Root.Black = true
}

func (t *RBTree) predecessor(x *RBTreeNode) *RBTreeNode {
	if x.Left != nil {
		x = x.Left
		for x.Right != nil {
			x = x.Right
		}
		return x
	}
	y := x.P
	for y != nil && x == y.Left {
		y, x = y.P, y
	}
	return y
}

func (t *RBTree) successor(x *RBTreeNode) *RBTreeNode {
	if x.Right != nil {
		x = x.Right
		for x.Left != nil {
			x = x.Left
		}
		return x
	}
	y := x.P
	for y != nil && x == y.Right {
		y, x = y.P, y
	}
	return y
}

func (t *RBTree) delete(z *RBTreeNode) *RBTreeNode {
	var x, y *RBTreeNode
	// y is the node to delete
	if z.Left == nil || z.Right == nil {
		y = z
	} else {
		y = t.successor(z)
	}
	// x is the only child of y
	if y.Left != nil {
		x = y.Left
	} else {
		x = y.Right
	}
	// delete y
	if x != nil {
		x.P = y.P
	}
	if y == t.Root {
		t.Root = x
	} else if y == y.P.Left {
		y.P.Left = x
	} else {
		y.P.Right = x
	}
	if y != z {
		z.Val = y.Val
	}
	if y.Black == true {
		if x == nil {
			x = &RBTreeNode{Black: true, P: y.P} // make a temporary x
		}
		t.deleteFixup(x)
	}
	return y
}

func (t *RBTree) deleteFixup(x *RBTreeNode) {
	// x always points to a node which has an additional black
	// if x is red-black, then x is black
	// if x is root, then dispose the additional black
	for t.Root != nil && t.Root != x && x.Black {
		if x == x.P.Left || x.P.Left == nil { // x's brother cannot be nil , cause it will break Prop.4 (x provides 2 black while the other provides 1)
			w := x.P.Right                    // w is x's brother
			if w != nil && w.Black == false { // x.P must be black due to Prop.4
				w.Black = true
				x.P.Black = false
				t.leftRotate(x.P)
				w = x.P.Right // which is w.Left
				// fallthrough to the following cases
			}
			if (w.Left == nil || w.Left.Black == true) && (w.Right == nil || w.Right.Black == true) {
				// children of w are all black
				w.Black = false
				// x.P absorb a layer of black from its children, and turn to the node with additional black
				x = x.P
			} else {
				if w.Right == nil || w.Right.Black == true { // which means w.Left is red
					w.Left.Black = true
					w.Black = false
					t.rightRotate(w)
					w = x.P.Right
					// fallthrough
				}
				// 1. w pushes black to his child(left, right).
				// 2. w.left is with an additional black now, while w.right turnes to black(from red).
				// 3. w is red, we can cast a left rotate, which causes w and w.P exchanges its color. Notich that w.left is handed-off to x.P.right
				// 4. x.P absort a layer of black from his children, which both have an additional black.
				// 5. x.P turnes to black(from red)
				w.Black = x.P.Black
				x.P.Black = true
				if w.Right != nil {
					w.Right.Black = true
				}
				t.leftRotate(x.P)
				x = t.Root // if the old w becomes root, make sure it turns to black
			}
		} else {
			w := x.P.Left
			if w != nil && w.Black == false {
				w.Black = true
				x.P.Black = false
				t.rightRotate(x.P)
				w = x.P.Left
			}
			if (w.Left == nil || w.Left.Black == true) && (w.Right == nil || w.Right.Black == true) {
				w.Black = false
				x = x.P
			} else {
				if w.Left == nil || w.Left.Black == true {
					w.Right.Black = true
					w.Black = false
					t.leftRotate(w)
					w = x.P.Left
					// fallthrough
				}
				w.Black = x.P.Black
				x.P.Black = true
				if w.Left != nil {
					w.Left.Black = true
				}
				t.rightRotate(x.P)
				x = t.Root
			}
		}
	}
	x.Black = true
}

func PrintTree(tree *RBTree) {
	list := make([]*RBTreeNode, 0)
	list = append(list, tree.Root)
	ans := make([]string, 0)
	for len(list) > 0 {
		top := list[0]
		list = list[1:]
		if top != nil {
			s := strconv.Itoa(top.Val)
			if !top.Black {
				s += "r"
			}
			ans = append(ans, s)
			list = append(list, top.Left, top.Right)
		} else {
			ans = append(ans, strconv.Itoa(-1))
		}
	}
	fmt.Println(ans)
}
