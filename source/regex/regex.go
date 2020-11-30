package main

import (
	"fmt"
	"regexp"
)

func main() {
	test1()
}

func test1() {
	str := "abc abd a*c a.c a1c cba 3.14 4.3a a.34 0.12"

	reg := regexp.MustCompile(`a.c`)

	parts1 := reg.FindAllString(str, -1)
	fmt.Println(parts1)

	reg = regexp.MustCompile(`\d+\.\d+`)
	parts2 := reg.FindAllStringSubmatch(str, -1)
	fmt.Println(parts2)

	str = ""
	reg = regexp.MustCompile(`<div>(.*)</div>`)
	reg = regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	str = "email to pain@gmail.com if ok"
	reg = regexp.MustCompile("pain@gmail.com")
	email := reg.FindString(str)
	fmt.Println(email)

	//reg = regexp.MustCompile(".+@.+\\..+")
	//reg = regexp.MustCompile(`.+@.+\..+`)
	reg = regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	email = reg.FindString(str)
	fmt.Println(email)

	str = `email to pain@gmail.com if ok
           email to jack@qq.com if nice 
                page@163.com   loc@sraf.com `

	// find all
	emails := reg.FindAllString(str, -1)
	fmt.Println(emails)

	reg = regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.[a-zA-Z0-9]+`)
	matches := reg.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		fmt.Println(match)
	}
}
