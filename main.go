package main

import (
	"fmt"
	"selenium/util"
)

func main() {
	util.SetUp()
	defer util.TearDown()

	wd := util.NewDriver()
	s, _ := wd.NewSession()
	fmt.Println("Session ID:", s)
	if err := wd.Get("https://google.com"); err != nil {
		panic(err)
	}
	t, _ := wd.Title()
	fmt.Println("Title of current page:", t)
}
