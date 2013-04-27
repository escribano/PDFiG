package pdf

import "github.com/mawicks/PDFiG/containers"

type Array struct {
	array containers.Array
}

// Constructor for Name object
func NewArray() *Array {
	return &Array{containers.NewDynamicArray(4)}
}

func (a *Array) Clone() Object {
	newArray := NewArray()
	size := a.Size()
	for i := 0; i<size; i++ {
		o := a.At(i)
		newArray.array.PushBack(o.Clone())
	}
	return newArray
}

func (a *Array) Dereference(...File) Object {
	return a
}

func (a *Array) Add(o Object) {
	a.array.PushBack(o)
}

func (a *Array) Size() int {
	return int(a.array.Size())
}

func (a *Array) At(i int) Object {
	return (*a.array.At(uint(i))).(Object)
}

func (a *Array) Serialize(w Writer, file ...File) {
	w.WriteByte('[')
	size := a.Size()
	for i := 0; i < size; i++ {
		if i != 0 {
			w.WriteByte(' ')
		}
		o := a.At(i)
		o.Serialize(w, file...)
	}
	w.WriteByte(']')
}
