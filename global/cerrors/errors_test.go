package cerrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestCErrors(t *testing.T) {
	err := New500ServiceError()
	var e ServiceError

	if errors.As(err, &e) {
		fmt.Println(e.ErrCode())
	}

	fmt.Println(e.ErrMessage())
}
