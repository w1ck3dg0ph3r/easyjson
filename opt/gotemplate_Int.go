// generated by gotemplate

package opt

import (
	"fmt"

	"github.com/w1ck3dg0ph3r/easyjson/jlexer"
	"github.com/w1ck3dg0ph3r/easyjson/jwriter"
)

// template type Optional(A)

// A 'gotemplate'-based type for providing optional semantics without using pointers.
type Int struct {
	V       int
	Defined bool
}

// Creates an optional type with a given value.
func OInt(v int) Int {
	return Int{V: v, Defined: true}
}

// Get returns the value or given default in the case the value is undefined.
func (v Int) Get(deflt int) int {
	if !v.Defined {
		return deflt
	}
	return v.V
}

// MarshalEasyJSON does JSON marshaling using easyjson interface.
func (v Int) MarshalEasyJSON(w *jwriter.Writer) {
	if v.Defined {
		w.Int(v.V)
	} else {
		w.RawString("null")
	}
}

// UnmarshalEasyJSON does JSON unmarshaling using easyjson interface.
func (v *Int) UnmarshalEasyJSON(l *jlexer.Lexer) {
	if l.IsNull() {
		l.Skip()
		*v = Int{}
	} else {
		v.V = l.Int()
		v.Defined = true
	}
}

// MarshalJSON implements a standard json marshaler interface.
func (v Int) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	v.MarshalEasyJSON(&w)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON implements a standard json unmarshaler interface.
func (v *Int) UnmarshalJSON(data []byte) error {
	l := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&l)
	return l.Error()
}

// IsDefined returns whether the value is defined, a function is required so that it can
// be used in an interface.
func (v Int) IsDefined() bool {
	return v.Defined
}

// String implements a stringer interface using fmt.Sprint for the value.
func (v Int) String() string {
	if !v.Defined {
		return "<undefined>"
	}
	return fmt.Sprint(v.V)
}
