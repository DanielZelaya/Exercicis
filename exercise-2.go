package main

import (
	"fmt"
)

//var x int = 42
//var y string = "James Bond"
//var z bool = true

//type hogdog int

//var x hogdog
//var y int

//const (
//	a = "Hello this so ez"
//	b = 42
//	c = true

//)

func main()  {

	//Exercises - Ninja Level 2

	//Exercise 1
	//var number int
	//fmt.Println("Hello, can you give me the number than you want conver")
	//fmt.Scan(&number)
	//fmt.Printf("%d\t%b\t%x", number, number, number )

	//Exercise 2

	//a := (42 == 42)
	//b := (42 <= 43)
	//c := (42 >= 43)
	//d := (42 != 43)
	//e := (42 < 43)
	//f := (42 > 43)
	//fmt.Println(a, b, c, d, e, f)

	//Exercise 3


	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Printf("%T\n", a)
	//fmt.Printf("%T\n", b)
	//fmt.Printf("%T\n", c)

	//Exercise 4

	var d int

	fmt.Println("Tell some number")
	fmt.Scan(&d)
	fmt.Printf("%d\t%b\t%x",d, d, d)
	e := d  << 1
	fmt.Println(e)
	fmt.Printf("%b\t%x", e, e)
	///////////////////////////////////
	//Exercises - Ninja Level 1
	// Exercise 2
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// Exercise 3
	// s := fmt.Sprintf( "%v\t%v\t%v", x, y, z)
	//	fmt.Println(s)

	// Exercise 4
	// fmt.Println(x)
	// fmt.Printf("%T\n" , x)
	// x=42
	// fmt.Println(x)

	// Exercise 5

	//fmt.Println(x)
	//fmt.Printf("%T\n", x)
	//x = 42
	//fmt.Println(x)
	//y = int(x)
	//fmt.Println(y)
	//fmt.Printf("%T\n",y)




}
