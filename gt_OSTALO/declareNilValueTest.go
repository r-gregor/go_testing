package main

import "fmt"

type Config struct {
	path string
}

func (c *Config) Path() string {
	if c == nil {
		return "/usr/home"
	}
	return c.path
}

func main() {
	var c1 *Config    // default to nil !!
	var c2 = &Config{ // initialized
		path: "/export",
	}
	fmt.Println(c1.Path(), c2.Path())
}
