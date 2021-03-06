package pdf

// PDF "Null" object
// Implements: pdf.Object
type Null struct{}

var nullSingleInstance Null

func NewNull() Object {
	return &nullSingleInstance
}

func (n *Null) Serialize(w Writer, file ...File) {
	w.WriteString("null")
	return
}

func (n *Null) Clone() Object {
	// All copies of null point to the same instance.
	return n
}

func (n *Null) Dereference() Object {
	return n
}

// Null is an immutable singleton, so the return value of Protect() can
// safely be cast back to Null.
func (n *Null) Protect() Object {
	return n
}

// Protect and unprotected interfaces are the same for null.
// Simply return the instance.
// The return value can safely be cast back to Null.
func (n *Null) Unprotect() Object {
	return n
}
