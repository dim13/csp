package csp

import "testing"

func TestString(t *testing.T) {
	s := []string{"x"}
	String{&s}.Set("a,b,c")
	for n, v := range []string{"a", "b", "c"} {
		if s[n] != v {
			t.Error("got", s)
		}
	}
}

func TestInt(t *testing.T) {
	i := []int{42}
	Int{&i}.Set("1,2,3")
	for n, v := range []int{1, 2, 3} {
		if i[n] != v {
			t.Error("got", i)
		}
	}
}
