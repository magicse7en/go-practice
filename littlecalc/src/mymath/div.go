package littlemath

import "errors"

func Div(a, b int) (ret int, err error) {
	if b == 0 {
		err = errors.New("divider cannot be 0!")
		return
	}
	return a / b, nil
}
