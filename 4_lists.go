package main
import "fmt"


func main() {
	// Declare an Array
	// here length is defined
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// here length is inferred
	var arr3 = [...]int{1,2,3}
	arr4 := [...]int{4,5,6,7,8}

	fmt.Println(arr3)
	fmt.Println(arr4)

	// Declare an array of strings
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	// Access Elements of an Array
	prices := [3]int{10,20,30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])
	prices[2] = 50
	fmt.Println(prices)

	// Array Initialization
	arr5 := [5]int{} // not initialized
	arr6 := [5]int{1,2} // partially initialized
	arr7 := [5]int{1,2,3,4,5} // fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	// Initialize Only Specific Elements
	arr8 := [5]int{1:10,2:40} // [0 10 40 0 0]
	fmt.Println(arr8)

	// Find the Length of an Array
	arr9 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr10 := [...]int{1,2,3,4,5,6}

	fmt.Println(len(arr9))
	fmt.Println(len(arr10))
}
