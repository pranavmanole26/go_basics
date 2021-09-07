package main

import (
	"encoding/json"
	"fmt"
)

//Start: Creating enum for error code
type custErrorCode int

const (
	error_400 = iota
	error_402
	error_404
	error_500
)

func (ced custErrorCode) string() string {
	return []string{"400", "402", "404", "500"}[ced]
}

//End: Creating enum for error code

type myError struct {
	Code custErrorCode
	Msg  string
}

func (me *myError) Error() string {
	return fmt.Sprintf(`{"code": "%s", "msg": "%s"}`, me.Code.string(), me.Msg)
}

func newError(code custErrorCode, msg string) error {
	return &myError{Code: code, Msg: msg}
}

func search(ns []int, n int) (int, error) {
	found := false
	index := -1
	var _n int
	var err error
	for index, _n = range ns {
		if _n == n {
			found = true
			break
		}
	}
	if !found {
		err = newError(error_404, "Number not found in array")
	}
	return index, err
}

func main() {
	var nums = []int{1, 2, 3, 4}
	index, err := search(nums, 5)
	if err == nil {
		fmt.Println("Number found at index:", index)
	} else {
		var me myError
		json.Unmarshal([]byte(err.Error()), &me)
		fmt.Printf("Code: %s, error message: %s", me.Code.string(), me.Msg)
	}
}
