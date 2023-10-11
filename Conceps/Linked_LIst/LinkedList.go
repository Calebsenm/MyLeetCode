package main

import (
	"fmt"
)

type Node struct{
	name string 
	age  int 
	next  *Node
}

type LinkedList struct{
	head *Node 
	size  int 
}

//Add a new element to the end
func (list *LinkedList) Append(name string , age int ){
	newNode := &Node{name: name, age:  age }

	if list.head == nil{
		list.head = newNode
	}	else {
		current := list.head 

		for current.next != nil{
			current = current.next
		}
		current.next = newNode
	}

	list.size ++
}

// delete a element of a list by name 

func (list *LinkedList) Remove(name string ) bool{
	if list.head == nil{
		return false
	}

	if list.head.name == name {
		list.head = list.head.next
		list.size --
		return true
	}

	current := list.head

	for current.next != nil{
		if current.next.name == name{
			current.next = current.next.next
			list.size--
			return true
		}
		current = current.next
	}
	return false
}

// Search by name 

func ( list *LinkedList) Search (name string ) *Node{
	current := list.head

	for current != nil{
		if current.name == name {
			return current
		}
		current = current.next
	}

	return nil
}


// Print all the data 

func (list *LinkedList ) Print(){
	current := list.head

	for current != nil {
		fmt.Println("%d", current )
		current = current.next
	}
	fmt.Println()
}

func main(){
	list := LinkedList{}

	list.Append("Caleb", 18)
	list.Append("Belac", 19)

	fmt.Println("Lista Enlazada: ")
	list.Print()

	nodeSearc := list.Search("Caleb")
	
	if nodeSearc != nil{
		fmt.Println("Elemento encontrado:", nodeSearc)
	}	else {
		fmt.Println("Elemento no encontrado")
	}


}


/*
Linked List is:

[]->[]->[]->[]->[]->nil

*/
