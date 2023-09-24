package main

import (
	"fmt"
	list "list/storage"
)
//импортировать файл node и list в main не получилось//

func main() {
	list := &List{}
	list.Add(5)
	list.Add(2)
	list.Add(3)
	list.Add(5)
	list.Add(4)
	list.Add(5)
	fmt.Println("Initial List: ")
	listPrint(list)
	fmt.Println("Len List: ")
	listLen(list)
	fmt.Println("Get all: ")
	list.GetAll()
	fmt.Println("Get by index: ")
	list.GetByIndex(500)
	fmt.Println("Get by value: ")
	list.GetByValue(5)
	//fmt.Println("Get All by value: ")
	//list.GetAllByValue(5)
	//fmt.Println("RemoveByValue: ") Почему-то не работает
	//list.RemoveByValue(3)
	//listPrint(list)
	//fmt.Println("RemoveByIndex: ")
	//list.RemoveByIndex(500000) Почему-то зависает
	//listPrint(list)
	//fmt.Println("RemoveAllByValue: ")
	//list.RemoveAllByValue(3)
	//listPrint(list)
	fmt.Println("Clear List: ")
	Clear(list)
	listPrint(list)
	fmt.Println("Clear List len: ")
	listLen(list)

}
