package response

import (
	"net/http"
	"testing"
	"errors"
)

func TestNewResponse(t *testing.T) {
	response := New().(*response)
	if response.status != http.StatusOK {
		t.Fail()
	}
}

func TestSliceResult(t *testing.T) {

	var tests = []struct {
		slice  interface{}
		size int
	}{
		{[]string{"Helllo", "world", "you"}, 3},
		{&[]string{"Nice"}, 1},
	}

	for _, test := range tests {
		r := New().SliceResult(test.slice).(*response)
		size := r.data["result_size"].(int)

		if size != test.size {
			t.Errorf("size don't match, expected:%v got:%v", size, test.size)
		}
	}

}

func TestErrorMap(t *testing.T) {

	custom := errors.New("my custom error")

	// Register custom error
	RegisterError(custom, http.StatusNotFound)

	r := Error(custom).(*response)

	if r.status != http.StatusNotFound {
		t.Errorf("status is bad, expected: %v got:%v", http.StatusNotFound, r.status)
	}
}