package log

import (
	"fmt"

	"github.com/arcane-craft/monadic/io"
	"github.com/arcane-craft/monadic/lazy"
)

func Print(a ...lazy.Any) io.IO[int] {
	return io.ApplyVarP1R(fmt.Print, a...)
}
