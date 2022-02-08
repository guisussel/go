package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var myBoolean bool      //false
var myVar string        //""
var myInt, otherInt int //0
var myString = "this is a string"
var anotherString string

//ints types
var myInt8 int8
var myInt16 int16
var myInt32 int32
var myInt64 int64

var myUint8 uint8
var myUint16 uint16
var myUint32 uint32
var myUint64 uint64

var myRune rune //same as int32
var myByte byte //same as uint8

//floats
var myFloat32 float32
var myFloat64 float64

//conts
var myConst = "one const var"
var myIntConst = 123
var myFloatConts = 123.4

func main() {
	fmt.Println("Calculating on goLang: ")
	//log.fatal("Oh no!!")

	//math operators
	myInt = 10
	myInt = myInt * 10
	myInt += 1

	myBoolean = myInt < otherInt

	fmt.Println("calc result: ", myInt)
	fmt.Println("boolean result: ", myBoolean)

	//string concatenation
	anotherString := "concatenating " + myString + " strings together! "
	fmt.Println("concat result: " + anotherString)

	//string conversion
	myString = strconv.Itoa(123456) //integer to ascii
	fmt.Println("conversion result: " + myString)
	myInt, _ = strconv.Atoi("123456") //ascii to integer

	fmt.Println(myBoolean)
	myBoolean, _ = strconv.ParseBool("true")
	fmt.Println(myBoolean)
	fmt.Println(reflect.TypeOf(myBoolean))

	//composite types
	//arrays
	var exampleArray [3]int
	exampleArr := [3]int{1, 2, 3}
	fmt.Println(exampleArray[0])
	fmt.Println(exampleArr[0])

	//slices
	//var exampleSlice []int
	exampleSlice := []string{"this", "is", "a", "hello", "world"}
	fmt.Println(exampleSlice[2])
	fmt.Printf("exampleArr: %v", exampleArr)

	//who slices the slices??

	firstTwo := exampleSlice[:2]
	lastThree := exampleSlice[2:4]
	lastElement := exampleSlice[len(exampleSlice)-1:]

	fmt.Printf("\nLastThree: %+v \n FirstTwo: %+v \n LastElement: %+v \n", firstTwo, lastThree, lastElement)

	//slices using make()
	slice := make([]int, 10, 15)
	fmt.Println(slice)

	//change values in arrays and slices
	fmt.Printf("\nexampleArr before the change: %v ", exampleArr)
	exampleArr[0] = 5 //change first value in exampleArr to 5
	fmt.Printf("\nexampleArr after the change: %v \n", exampleArr)

	//append
	//exampleArr = append(exampleArr, 8) //error!
	exampleSlice = append(exampleSlice, ", and This works!")
	fmt.Println(exampleSlice)

}
