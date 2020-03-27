// Package csp implements comma separated flag parameters
package csp

import (
	"strconv"
	"strings"
)

type String struct {
	S *[]string
}

func (m String) String() string {
	return strings.Join(*m.S, ",")
}

func (m String) Set(s string) error {
	*m.S = strings.Split(s, ",")
	return nil
}

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
