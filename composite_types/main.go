package main

import "fmt"

func main() {
	var a [3]int = [3]int{2, 1, 2}

	q := [...]int{1, 2, 1}

	for index, val := range a {
		fmt.Println(index, val)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Println(a, q, a == q)

	// currency

	type Currency int
	const (
		USD Currency = iota
		EUR
		IN
	)

	symbol := [...]string{USD: "Dollar", EUR: "Pounds", IN: "Rupee"}
	fmt.Println(symbol)

	// comparing

	w := [100]int{91: -2}
	e := [100]int{98: -2}
	fmt.Println(w == e)

	// slicing

	s := []int{1, 2, 3, 4}

	// fmt.Println(s, cap(s), len(s))
	// fmt.Println(s[:4])
	fmt.Println(s)

	// fmt.Println(remove(s, 1))
	// fmt.Println(s)
	// fmt.Println(removePtr(s, 1))

	// maps

	ages := make(map[string]int)
	fmt.Println(ages)

	ages = map[string]int{
		"alice":   12,
		"beyonce": 124,
	}
	fmt.Println(ages)

	alice, ok := ages["alices"]
	fmt.Println(alice, ok)

}

func map_equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}

// func remove(slice []int, i int) []int {
// 	copy(slice[i:], slice[i+1:])
// 	return slice[:len(slice)-1]
// }

// func remove(arr *int, i int) {
// 	fmt.Println()
// }
