package strings

import (
	"fmt"
	//"strings"  // this package contains different string methods for string operations
)

func StringsFun() {
	s1 := "HEllo world"

	// you can use backtick to create strings with multiple lines
	s2 := `Line1
Line2`

	fmt.Println(s1, s2)

	// length of string

	fmt.Println(len(s1))

	// iterate over string as well

	for _, r := range s1 {
		fmt.Println(string(r)) // rune
	}
	// convert a string into rune slice
	rn := []rune("heythere")
	fmt.Println(rn)

	//you can iterate over it as well

	for i := 0; i < len(rn); i++ {
		fmt.Println(string(rn[i]))
	}

	// string to byte slice

	bt := []byte("heythere")

	fmt.Println(bt) // unicode value of alphabets

	//strings.Join()

}
