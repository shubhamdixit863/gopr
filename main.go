package main

import (
	"fmt"
	"gointerview/strings"
)

func Foo(ar [4]int) {
	ar[1] = 9

}

func Foo2(ar []int) {
	ar[1] = 9

}

func main() {

	// Arrays are of fixed sizes
	var myArray [4]int
	c := &myArray[0] // integer takes 8 bytes or 64 bits
	myArray[1] = 9
	d := &myArray[1]
	fmt.Println(c, d)

	// Arrays are value types
	myArray2 := [4]int{}
	myArray2[1] = 9
	fmt.Println(myArray == myArray2) // so arrays can be compared and they can be equal
	// if the values in both the arrays are equal

	// slices in go
	// Dynamic arrays ,you dont have to define the length initially

	myslice := []int{1, 2}
	myslice[1] = 9
	//fmt.Println(myslice)

	// Using make function
	// this make function creates a memory space

	mySlice2 := make([]int, 3) // length is 3 capacity is 3
	//mySlice3 := make([]int, 3,7)  // length is 3 capacity is 7

	fmt.Printf("Address of the slice: %p\n", &mySlice2)
	mySlice2[0] = 9
	mySlice2[1] = 90
	mySlice2[2] = 91
	//mySlice2[3] = 17
	// append will create a new memory location once the capacity of underlying is reached
	// it will copy all the items and will paste in the new underlying
	/**
	In Go, a slice is a reference to an underlying array.
	When you append items to a slice and its capacity is reached,
	Go automatically creates a new array with more capacity to accommodate the new elements.
	This process is often referred to as "growing the slice".

	1-When appending an element to a slice, Go checks if there is enough capacity
	in the underlying array to accommodate the new element.
	2-If there is enough capacity, Go will insert the element into the existing array.
	3-If there is not enough capacity, Go will allocate a new array with a larger capacity.
	The exact growth strategy isn't specified in the language spec but typically the new array's capacity is double the old capacity,
	as long as the number of elements is small.
	The growth factor may be reduced for very large slices to avoid excessive memory allocation.
	4- The elements from the original array are copied to the new array.
	5- The slice is then updated to reference the new array.
	*/
	mySlice2 = append(mySlice2, 88, 90, 99, 90)

	//mySlice3 := make([]int, 3) // length is 3 capacity is 3

	fmt.Printf("Address of the slice: %p\n", &mySlice2)

	///fmt.Println(mySlice2==mySlice3)  // equality operator is not defined in slice

	// nil slice
	//var slc []int
	//fmt.Print(slc)
	arr := [4]int{}
	fmt.Println("Initial", arr)

	Foo(arr)
	fmt.Println("After Modify", arr)

	slc := make([]int, 4) // slc is already an address of underlying array
	fmt.Println("Initial Slice", slc)

	Foo2(slc)
	fmt.Println("After Modify Slice", slc)

	// go is always pass by value that value can be or cannot be an address
	VariadicFoo(9, 90, 89, 98, 71, 78, 89, 12)

	// Map a simple hashmap ,maps are also reference types
	// when you create a map in golang an underlying hash table is created

	// nil map ,it does not have an underlying hash table its just a descriptor
	var mp map[string]string
	//mp["name"] = 8  //  this is not allowed
	fmt.Println("Value access", mp["something"])

	// You have to initialize the map
	mp2 := map[string]int{
		"name": 10,
	}
	mp2["something"] = 999

	fmt.Println(mp2["name"], mp2["something"])

	// Using make function
	mk := make(map[string]string)
	mk["name"] = "shawn"
	fmt.Println(mk)

	//O(N) , in map getting and setting is of constant order O(1)
	//

	strings.StringsFun()

}

// variadic functions

func VariadicFoo(data ...int) {
	for _, v := range data {
		fmt.Println(v) // data will be available in form of slice

	}
}
