package list

type node struct {
	index int64 // уникальный индекс ноды. Необходим для того, чтобы можно было удалять ноды из списка
	value int64
	next  *node
}
