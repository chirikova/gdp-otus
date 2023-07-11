package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	start *ListItem
	end   *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.start
}

func (l *list) Back() *ListItem {
	return l.end
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{v, nil, nil}

	// кейс если стартовый уже определен
	if l.start != nil {
		l.start.Prev, newItem.Next = &newItem, l.start
	}

	// кейс определен только стартовый
	if l.start != nil && l.end == nil {
		l.end = l.start
	}

	// кейс если стартовый не определен, но задан конечный
	if l.start == nil && l.end != nil {
		l.end.Prev = &newItem
		newItem.Next = l.end
	}

	l.start = &newItem

	l.len++

	return l.Front()
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{v, nil, nil}

	// кейс если конечный уже определен
	if l.end != nil {
		l.end.Next, newItem.Prev = &newItem, l.end
	}

	// кейс определен только конечный
	if l.end != nil && l.start == nil {
		l.start = l.end
	}

	// кейс если конечный не определен, но задан стартовый
	if l.start != nil && l.end == nil {
		l.start.Next = &newItem
		newItem.Prev = l.start
	}

	l.end = &newItem

	l.len++

	return l.Back()
}

func (l *list) Remove(i *ListItem) {
	l.len--

	// проверяем кейс если удаляется стартовый элемент и у него есть следующий элемент
	if i == l.start && i.Next != nil {
		i.Next.Prev = nil
		l.start = i.Next
		return
	}

	// проверяем кейс если удаляется стартовый элемент без следующего элемента
	if i == l.start {
		l.start = nil
		return
	}

	// проверяем кейс если удаляется конечный элемент и у него есть предыдущий элемент
	if i == l.end && i.Prev != nil {
		i.Prev.Next = nil
		l.end = i.Prev
		return
	}

	// проверяем кейс если удаляется конечный элемент без предыдущего элемента
	if i == l.end {
		l.end = nil
		return
	}

	// кейс если удаляется не стартовый и не конечный элемент
	i.Prev.Next, i.Next.Prev = i.Next, i.Prev
}

func (l *list) MoveToFront(i *ListItem) {
	// кейс если уже впереди
	if i.Prev == l.start {
		i.Prev.Prev = i
	}
	if i == l.start {
		return
	}

	// кейс если меняем местами с начальным
	if i.Prev == l.start {
		i.Prev.Prev = i
	}

	// кейс если двигаем с самого конца
	if i == l.end {
		i.Prev.Next = nil
		l.end = i.Prev
	} else {
		i.Prev.Next, i.Next.Prev = i.Next, i.Prev
	}

	i.Prev = nil
	i.Next = l.start
	l.start = i
}

func NewList() List {
	return new(list)
}
