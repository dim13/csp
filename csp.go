// Package csp implements comma separated flag parameters
package csp

import (
	"strconv"
	"strings"
)

// String is wrapper for comma separated strings flag
type String struct {
	S *[]string
}

func (m String) String() string {
	return strings.Join(*m.S, ",")
}

// Set flag value
func (m String) Set(s string) error {
	*m.S = strings.Split(s, ",")
	return nil
}

// Int is wrapper for comma separated integer flag
type Int struct {
	I *[]int
}

func (m Int) String() string {
	var s []string
	for _, v := range *m.I {
		s = append(s, strconv.Itoa(v))
	}
	return strings.Join(s, ",")
}

// Set flag value
func (m Int) Set(s string) error {
	for _, v := range strings.Split(s, ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*m.I = append(*m.I, i)
	}
	return nil
}
