package linearLinkedListCApproach

import "fmt"

func singleLinkedList() {

	type node struct {
		data int
		link *node
	}

	var root *node

	length := func() int {
		var count int

		temp := root

		for temp != nil {
			count++
			temp = temp.link
		}

		return count
	}

	createNode := func() (*node, bool) {
		var data int

		fmt.Print("Enter a data to insert: ")

		_, err := fmt.Scan(&data)

		if err != nil {
			panic(err)
		}

		newNode := &node{data: data, link: nil}

		if root != nil {
			return newNode, false
		}

		root = newNode

		return newNode, true
	}

	insertAtStart := func() {
		newNode, isInitialAppend := createNode()

		if isInitialAppend {
			return
		}

		newNode.link = root

		root = newNode
	}

	insertAtEnd := func() {
		newNode, isInitialAppend := createNode()

		if isInitialAppend {
			return
		}

		temp := root

		for temp.link != nil {
			temp = temp.link
		}

		temp.link = newNode
	}

	insertAtSpecific := func() {
		var (
			pos        int
			prev, temp *node
		)

		fmt.Print("Enter the position to insert: ")

		_, err := fmt.Scan(&pos)

		if err != nil {
			panic(err)
		}

		if pos == 1 {
			insertAtStart()
			return
		}

		if pos == length()+1 {
			insertAtEnd()
			return
		}

		if pos > length()+1 {
			fmt.Println("Position exceeds the list length")

			return
		}

		newNode, isInitialAppend := createNode()

		if isInitialAppend {
			return
		}

		temp = root

		for i := 1; i < pos; i++ {
			prev = temp

			temp = temp.link
		}

		newNode.link = temp

		prev.link = newNode
	}

	display := func() {

		temp := root

		for temp != nil {
			fmt.Printf("%d\t", temp.data)

			temp = temp.link
		}
	}

	deleteAtStart := func() {
		if root == nil {
			fmt.Println("List is empty")
			return
		}
		root = root.link
	}

	deleteAtEnd := func() {
		if root == nil {
			fmt.Println("List is empty")
			return
		}

		var prev *node

		temp := root

		for temp.link != nil {
			prev = temp
			temp = temp.link
		}

		prev.link = nil
	}

	deleteAtSpecific := func() {
		var (
			pos        int
			prev, temp *node
		)

		if root == nil {
			fmt.Println("List is empty")
			return
		}

		fmt.Print("Enter the position to delete: ")

		_, err := fmt.Scan(&pos)

		if err != nil {
			panic(err)
		}

		if pos == 1 {
			deleteAtStart()
			return
		}

		if length() == pos {
			deleteAtEnd()
			return
		}

		if pos > length() {
			fmt.Println("Position exceeds the list length")
			return
		}

		temp = root

		for i := 1; i < pos; i++ {
			prev = temp
			temp = temp.link
		}

		prev.link = temp.link

		temp.link = nil
	}

	search := func() {

		var data int

		fmt.Print("Enter data to search: ")

		_, err := fmt.Scan(&data)

		if err != nil {
			panic(err)
		}

		temp := root

		for temp != nil {
			if temp.data == data {
				fmt.Println("Data found")
				return
			}

			temp = temp.link
		}

		fmt.Println("Data not found")
	}

	var choice int

	for true {

		fmt.Print("\n1.Insert at beginning\n2.Insert at specific position\n3.Insert at end\n4.Display\n5.Delete at beginning\n6.Delete at specific position\n7.Delete at end\n8.Search\n9.Length of the list")

		fmt.Print("\nChoose an operation: ")

		_, err := fmt.Scan(&choice)

		if err != nil {
			panic(err)
		}

		switch choice {
		case 1:
			insertAtStart()
		case 2:
			insertAtSpecific()
		case 3:
			insertAtEnd()
		case 4:
			display()
		case 5:
			deleteAtStart()
		case 6:
			deleteAtSpecific()
		case 7:
			deleteAtEnd()
		case 8:
			search()
		case 9:
			fmt.Println("Length of the list is", length())
		}
	}
}
