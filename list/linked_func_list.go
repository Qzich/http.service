package list

import "container/list"

func NewLinkedFuncList() *LinkedFuncList {
	return &LinkedFuncList{List: list.New()}
}

type LinkedFuncList struct {
	List    *list.List
	element *list.Element
}

func (linkedFuncList *LinkedFuncList) InvokeAll() error {
	var err error

	for currentElement := linkedFuncList.List.Front(); currentElement != nil; currentElement = currentElement.Next() {

		handler := currentElement.Value.(func() error)
		if err = handler(); err != nil {
			return err
		}
	}

	return err
}

func (linkedFuncList *LinkedFuncList) Push(f func() error) *LinkedFuncList {
	if linkedFuncList.element == nil {
		linkedFuncList.element = linkedFuncList.List.PushFront(f)
	} else {
		linkedFuncList.element = linkedFuncList.List.InsertAfter(f, linkedFuncList.element)
	}

	return linkedFuncList
}
