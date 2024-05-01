package main

import "testing"

func TestAdd(t *testing.T) {
	res := add(1, 2)

	if res != 4 {
		t.Errorf("错误")
	}
}
