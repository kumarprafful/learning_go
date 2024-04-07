package main

import "fmt"

var a = b + c
var b = f()
var c = 1

func f() int {
	return c + 1
}

func main() {
	f()
	s := "Nello world"
	fmt.Println(s[0], len(s))
	fmt.Println(s[0:3])
	q := `Helo Wol
	aslkjdkashnc
	d
	as
	as
	s
	
	'a]a]d]asd]adaa`
	fmt.Println(q)

	const pi = 3.1412
	fmt.Println(pi)
	// pi = 1
	type Weekday uint
	const (
		Sunday Weekday = 1 << iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println(ZiB / YiB)
}
