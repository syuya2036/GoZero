package model

import "fmt"

type List struct {
	Fs []*FunctionBase
	Len int
}

func NewList() *List {
    s := new(List)
    return s
}

func (l *List) Push(f *FunctionBase) error {
	if l == nil {
		return fmt.Errorf("list is nil")
	}
	l.Fs = append(l.Fs, f)
	return nil
}

func (l *List) Pop() (*FunctionBase, error) {
	if l == nil {
	    return nil, fmt.Errorf("list is nil") 
	}
	elem := l.Fs[l.Len - 1]
	l.Fs = l.Fs[:l.Len - 1]
	l.Len--

	return elem, nil
}