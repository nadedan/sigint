package some

import (
	"fmt"

	"github.com/nadedan/sigint"
)

type thing struct{}

func NewThing() *thing {
	t := &thing{}
	sigint.Defer(t.close)
	return t
}

func (t *thing) close() {
	fmt.Println("Closing the thing")
}
