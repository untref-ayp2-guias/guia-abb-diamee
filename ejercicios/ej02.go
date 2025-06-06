package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func PredecesorInOrder(bt *binarytree.BinarySearchTree[int], k int) (int, error) {

	if bt.Size() == 0 {
		return 0, errors.New("fsd")
	}

	n := bt.GetRoot()
	var predecesor *binarytree.BinaryNode[int] = nil

	for n != nil {
		if k > n.GetData() {
			predecesor = n
			n = n.GetRight()
		} else {
			n = n.GetLeft()
		}
	}

	if predecesor != nil {
		return predecesor.GetData(), nil
	}
	return 0, errors.New("No se encontró predecesor")

}
