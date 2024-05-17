package strings

import "github.com/arcane-craft/monadic/algebra"

type String string

func (String) Append(a String, b String) String {
	return a + b
}

func (String) Neutral() String {
	return ""
}

var _ = algebra.ImplMonoid[String]()
