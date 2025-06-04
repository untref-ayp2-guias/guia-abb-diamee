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
	for n.GetRight() != nil {
		x = n.GetData()
		n = n.GetRight()
	}
	return x, nil
}
