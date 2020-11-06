package tests

import (
	"github.com/w1ck3dg0ph3r/easyjson"
	"github.com/w1ck3dg0ph3r/easyjson/jlexer"
	"github.com/w1ck3dg0ph3r/easyjson/jwriter"
)

//easyjson:json
type NestedMarshaler struct {
	Value easyjson.MarshalerUnmarshaler
	Value2 int
}

type StructWithMarshaler struct {
	Value int
}

func (s *StructWithMarshaler) UnmarshalEasyJSON(w *jlexer.Lexer) {
	s.Value = w.Int()
}

func (s *StructWithMarshaler) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int(s.Value)
}

