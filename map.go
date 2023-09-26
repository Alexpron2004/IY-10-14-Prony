package main

import (
	"fmt"
	"math/rand"
)

type Map struct {
	id int64
	mp map[int64]int64
}

func NewMap() *Map {
	mp := &Map{
		mp: make(map[int64]int64),
		id: 0,
	}
	return mp
}

func (s Map) Add(data int64) (id int64) {
	id = rand.Int63()
	s.mp[id] = data
	return
}

func (s Map) Get(data int64) (id int64, ok bool) {
	id, ok = s.mp[data]
	fmt.Println(id, ok)
	return
}

func (s *Map) Len() {
	fmt.Println(len(s.mp))
	return
}

// RemoveByIndex удаляет элемент из списка по индексу
func (s *Map) RemoveByIndex(id int64) {
	for id := range s.mp {
		if s.id == id {
			delete(s.mp, id)
			break
		}
		return
	}
}

// RemoveByValue удаляет элемент из списка по значению
func (s Map) RemoveByValue(value int64) (id int64) {
	for id, data := range s.mp {
		if value == data {
			delete(s.mp, id)
			break
		}
	}
	return
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (s Map) RemoveAllByValue(value int64) (id []int64) {
	for id, data := range s.mp {
		if value == data {
			delete(s.mp, id)
		}
	}
	return
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (s Map) GetByIndex(id int64) (data int64, ok bool) {
	for id, data := range s.mp {
		if s.id == id {
			fmt.Println("ID:", id, "Data:", data)
			break
		} else {
			ok = false
			fmt.Println(nil, ok)
			break
		}
	}
	return
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (s Map) GetByValue(value int64) (id int64, ok bool) {
	if len(s.mp) == 0 {
		ok = false
		fmt.Println(nil, ok)
	}
	for id, data := range s.mp {
		if value == data {
			fmt.Println("ID:", id, "Data:", data)
			break
		}
	}
	return
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (s Map) GetAllByValue(value int64) (ids []int64, ok bool) {
	if len(s.mp) == 0 {
		ok = false
		fmt.Println(nil, ok)
	}
	for id, data := range s.mp {
		if value == data {
			fmt.Println("ID:", id, "Data:", data)
		}
	}
	return
}

// GetAll возвращает все элементы map
//
// Если список пуст, то возвращается nil и false.
func (s *Map) GetAll() (data []int64, ok bool) {
	for id, data := range s.mp {
		fmt.Println("ID:", id, "Data:", data)
	}
	if len(s.mp) == 0 {
		ok = false
		fmt.Println(nil, ok)
	}
	return
}

// Clear очищает список
func (s *Map) mapClear() {
	clear(s.mp)
	return
}

// Print выводит список в консоль
func (s Map) Print() {
	fmt.Println(s)
	return
}

// --------------------------------------------------------------------
func main() {
	l := NewMap()
	l.Add(5)
	l.Add(2)
	l.Add(5)
	l.Add(4)
	l.Add(5)
	fmt.Println("Initial Map: ")
	l.Print()
	fmt.Println("Len Map: ")
	l.Len()
	fmt.Println("Get all: ")
	l.GetAll()
	fmt.Println("Get by index: ")
	l.GetByIndex(1000)
	fmt.Println("Get by value: ")
	l.GetByValue(5)
	fmt.Println("Get All by value: ")
	l.GetAllByValue(5)
	fmt.Println("RemoveByValue: ")
	l.RemoveByValue(5)
	l.Print()
	fmt.Println("RemoveByIndex: ")
	l.RemoveByIndex(500000)
	l.Print()
	fmt.Println("RemoveAllByValue: ")
	l.RemoveAllByValue(5)
	l.Print()
	fmt.Println("Clear Map: ")
	l.mapClear()
	l.Print()
	fmt.Println("Clear Map len: ")
	l.Len()
}

//---------------------------------------------------------------------------------------------------------
