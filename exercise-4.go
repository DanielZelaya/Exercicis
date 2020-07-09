package main

//var x [5]int

func main() {

	//Exercises - Ninja Level 4
	//Exercises1
	//x[0]=10
	//x[1]=11
	//x[2]=12
	//x[3]=13
	//x[4]=14
	//for  v := range x {
	//	fmt.Println(v)
	//}
	//fmt.Printf("%T",x)

	//Exercises 2
	//a := make([]int,10)
	//var contador int = 0
	//for i := 0; i<=9; i++ {
	//	a[i] = contador
	//	contador++
	//}
	//for v := range a {
	//	fmt.Println(v)
	//}
	//fmt.Printf("%T\n",a)

	//Exercises 3
	//a := make([]int,10)
	//var contador int = 0
	//for i := 0; i<=9; i++ {
	//	a[i] = contador
	//	contador++
	//}
	//for v := range a {
	//	fmt.Println(v)
	//}
	//fmt.Printf("%T\n",a)
	//fmt.Println(a[:5])
	//fmt.Println(a[5:10])
	//fmt.Println(a[4:9])
	//fmt.Println(a[1:6])

	//Exercises 4
	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//x = append(x , 52)
	//fmt.Println(x)
	// = append(x,53,54,55)
	//fmt.Println(x)
	//y :=[]int{56, 57, 58, 59, 60}
	//x = append(x,y...)
	//fmt.Println(x)

	//Exercises 5
	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//y := []int{42, 43, 44, 48, 49, 50, 51}
	//x = append(x, y...)
	//fmt.Println(x)
	//x= append(x[:3], x[:6]...)
	//fmt.Println(x)

	//Exercises 6
	//	x := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, `	Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North 	Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` 	Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`,` Wyoming`}
	//	fmt.Println(len(x), cap(x))
	//	for v, i := range x {
	//		fmt.Println(i,v)
	//	}

	//Exercises 7
	//x := [][]string {{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Helloooooo, James."}}
	//for _,i := range x{
	//	fmt.Println(i)
	//}

	//Exercises 8
	//x := map[string][]string{
	//	"last_name": []string{"Daniel Zelaya Flores", "Francisco"},
	//	"Edad":      []string{"Tiene 27 años de edad", "Es mentira"},
	//	"Trabajo":   []string{"Programador", "Aun esta aprendiendo"},
	//}
	//fmt.Println(x)
	//for i, v := range x {
	//	fmt.Println(i)
	//	for k, v2 := range v {
	//		fmt.Println("\t", k, v2)
	//
	//	}
	//}

	//Exercises 9
	/*x := map[string][]string{
		"last_name": []string{"Daniel Zelaya Flores", "Francisco"},
		"Edad":      []string{"Tiene 27 años de edad", "Es mentira"},
		"Trabajo":   []string{"Programador", "Aun esta aprendiendo"},
	}
	fmt.Println(x)
	x["Apellido"] = []string{"javier", "jimenez"}
		for i,v := range x {
			fmt.Println(i)
			for k,v := range v{
				fmt.Println("\t",k,v)
			}
		}
	*/
	//Exercises 10
	/*x := map[string][]string{
		"last_name": []string{"Daniel Zelaya Flores", "Francisco"},
		"Edad":      []string{"Tiene 27 años de edad", "Es mentira"},
		"Trabajo":   []string{"Programador", "Aun esta aprendiendo"},
	}
	fmt.Println(x)
	x["Apellido"] = []string{"javier", "jimenez"}
	for i,v := range x {
		fmt.Println(i)
		for k,v := range v{
			fmt.Println("\t",k,v)
		}
	}
	delete(x, "Apellido")
	fmt.Println(x)
	*/

}
