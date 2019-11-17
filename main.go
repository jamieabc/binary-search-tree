package main

import "fmt"

func main() {
	bst := newBST([]int{})
	bst.Insert(100)
	bst.Insert(200)
	bst.Insert(150)
	bst.Insert(300)
	fmt.Printf("bst: %s\n", bst.String())

	search := 400
	location, _, err := bst.Search(search)
	if nil != err {
		fmt.Println("search not found for ", search)
	} else {
		fmt.Printf("%d at index %d\n", search, location)
	}

	remove := 200
	bst.Remove(remove)
	fmt.Printf("after remove 200, bst tree: %s\n", bst.String())
}
