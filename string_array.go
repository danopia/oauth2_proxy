package main

import (
	"strings"
)

type StringArray []string

func (a *StringArray) Set(s string) error {
	*a = append(*a, s)
	return nil
}

// Get returns the slice of strings
func (a *StringArray) Get() interface{} {
	return []string(*a)
}

func (a *StringArray) String() string {
	return strings.Join(*a, ",")
}
