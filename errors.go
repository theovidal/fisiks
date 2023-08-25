package fisiks

import (
	"fmt"
)

type WrongArgumentType struct {
	Argument string
	Expected string
}

func (w WrongArgumentType) Error() string {
	return fmt.Sprintf("Argument %s is invalid (expected type: %s)", w.Argument, w.Expected)
}
