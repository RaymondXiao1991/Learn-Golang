package main

import (
	"fmt"
	"unsafe"
)

func sliceToArrayPtr() {
    var b = []int{11, 12, 13}
    var p = (*[3]int)(b)
    p[1] = p[1] + 10
    fmt.Printf("%v\n", b)
}

func sliceToArrayPtr2() {
    var b = []int{11, 12, 13}
    var p0 = (*[0]int)(b)
    fmt.Printf("%v\n", *p0)
    var p1 = (*[1]int)(b)
    fmt.Printf("%v\n", *p1)
    var p2 = (*[2]int)(b)
    fmt.Printf("%v\n", *p2)
    var p3 = (*[3]int)(b)
    fmt.Printf("%v\n", *p3)
}

func sliceToArrayPtrWithPanic() {
    var b = []int{11, 12, 13}
    var p = (*[3]int)(b[:1])
    fmt.Printf("%v\n", *p)
}

/*
// this can cause compile error
func slice2array() {
	var b = []int{11, 12, 13}
	var a = [3]int(b)
	fmt.Printf("%v\n", a)
}
*/


func slice2ArrayPtrWithHack() {
	var b = []int{11, 12, 13}
    var p = (*[3]int)(unsafe.Pointer(&b[0]))
    p[1] += 10
    fmt.Printf("%v\n", p)
    fmt.Printf("%v\n", b)
}

func array2slice() {
	var a = [5]int{11, 12, 13, 14, 15}
    var b = a[0:len(a)]
    b[1] += 10
    fmt.Printf("%v\n", b)
}

func main() {
    sliceToArrayPtr()
    sliceToArrayPtr2()
    slice2ArrayPtrWithHack()
    array2slice()
    //sliceToArrayPtrWithPanic()
}

