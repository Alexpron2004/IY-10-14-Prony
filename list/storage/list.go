package list

import (
	"fmt"
	"math/rand"
)

type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() *List {
	l := &List{
		len:       0,
		firstNode: nil,
	}
	return l
}

// Len возвращает длину списка
func listLen(l *List) (len int64) {
	fmt.Println(l.len)
	return
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	id = rand.Int63()
	n := &node{
		id:    id,
		value: value,
		next:  nil,
	}
	if l.firstNode == nil {
		l.firstNode = n
		return
	}
	l.len++
	ln := l.firstNode
	for ln.next != nil {
		ln = ln.next
	}
	ln.next = n
	return
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	var pn = l.firstNode
	var cn = l.firstNode
	for cn != nil {
		if l.firstNode == nil {
			fmt.Println("Список пуст")
			break
		}
		if cn.id == id {
			if pn == nil {
				l.firstNode = cn.next
				return
			}
			pn.next = cn.next
			return
		}
	}
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	var pn = l.firstNode
	var cn = l.firstNode
	for cn != nil {
		if l.firstNode == nil {
			fmt.Println("Список пуст")
			break
		}
		if cn.value == value {
			if pn == nil {
				l.firstNode = cn.next
				return
			}
			pn.next = cn.next
			return
		}
	}
	return
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	var pn = l.firstNode
	var cn = l.firstNode
	for cn != nil {
		if l.firstNode == nil {
			fmt.Println("Список пуст")
			break
		}
		if cn.value == value {
			if pn == nil {
				l.firstNode = cn.next
				return
			}
			pn.next = cn.next
			return
		}
	}
	return
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	var cn = l.firstNode
	for cn != nil {
		if cn.id == id {
			fmt.Println("ID =", cn.id, "Value =", cn.value)
			return
		} else {
			cn = cn.next
		}
	}
	ok = false
	fmt.Println(0, ok)
	return
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	var cn = l.firstNode
	var f int64 = 0
	for cn != nil {
		if cn.value != value {
			cn = cn.next
		} else {
			f++
			fmt.Println("Value =", cn.value, "ID =", cn.id)
			break
		}
	}
	if f == 0 {
		ok = false
		fmt.Println(0, ok)
	}
	return
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (id []int64, ok bool) {
	var cn = l.firstNode
	var f int64 = 0
	for cn != nil {
		if cn.value != value {
			cn = cn.next
		} else {
			f++
			fmt.Println("Value =", cn.value, "ID =", cn.id)
		}
	}
	if f == 0 {
		ok = false
		fmt.Println(0, ok)
	}
	return
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.firstNode == nil {
		ok = false
		fmt.Println(l.firstNode, ok)
	} else {
		curr := l.firstNode
		for curr != nil {
			fmt.Printf("%d ", curr.value)
			curr = curr.next
		}
	}
	fmt.Println()
	return
}

// Clear очищает список
func Clear(l *List) (len int64) {
	l.firstNode = nil
	l.len = 0
	return
}

// Print выводит список в консоль
func listPrint(l *List) {
	curr := l.firstNode
	for curr != nil {
		fmt.Printf("%d ", curr.value)
		fmt.Println(curr.id)
		curr = curr.next
	}
	fmt.Println()
}

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
