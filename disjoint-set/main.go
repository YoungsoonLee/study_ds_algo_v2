package main

import (
	"fmt"
)

type Element struct {
	Parent *Element
	Size   int
	Data   interface{}
}

func MakeSet(Data interface{}) *Element {
	s := &Element{}
	s.Parent = s
	s.Size = 1
	s.Data = Data
	return s
}

func Find(e *Element) *Element {
	for e.Parent != e {
		e = e.Parent
	}
	return e
}

func Union(e1, e2 *Element) {
	//e1.Parent = e2
	e1SetName := Find(e1)
	e2SetName := Find(e2)
	if e1SetName == e2SetName {
		return
	}

	if e1SetName.Size < e2SetName.Size {
		e1SetName.Parent = e2SetName
		e2SetName.Size += e1SetName.Size
	} else {
		e2SetName.Parent = e1SetName
		e1SetName.Size += e2SetName.Size
	}
}

func main() {
	aSet := MakeSet("a")
	bSet := MakeSet("b")
	oneSet := MakeSet(1)
	twoSet := MakeSet(2)

	Union(aSet, bSet)
	Union(oneSet, twoSet)

	result := Find(aSet)
	fmt.Println(result.Data) //b

	result = Find(bSet)
	fmt.Println(result.Data) //b

	result = Find(oneSet)
	fmt.Println(result.Data) //2

	result = Find(twoSet)
	fmt.Println(result.Data) //2
}
