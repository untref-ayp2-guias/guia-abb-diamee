package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func SecondLargestElement(bt *binarytree.BinarySearchTree[int]) (int, error) {
	var x int
	if bt.Size() < 2 {
		return x, errors.New("No hay valores")
	}
	if bt.GetRoot() == nil {
		return 0, errors.New("No hay valores")
	}
	return findSecLargeElem(bt.GetRoot())
}

func findSecLargeElem(n *binarytree.BinaryNode[int]) (int, error) {

	var x int
	x = n.GetLeft().GetData()

	if n == nil {
		return 0, errors.New("no hay valores")
	}

	if n.GetRight() == nil {
		n = n.GetLeft()
		for n.GetRight() != nil {
			n = n.GetRight()
			x = n.GetData()
		}
	}

	for n.GetRight() != nil {
		x = n.GetData()
		n = n.GetRight()
	}

	return x, nil
}
