package main
import "fmt"

func area(l,w int)int{
	Ar:= l*w
	return Ar
}
func name(a string)string{
	p:= a
	return p
}
func merge(a1,a2 string)string{
	a3=a1+a2
	return a3
}
func main(){
	fmt.Println("Area of rectangle is :", area(12,10))
	fmt.Println("next code")
	fmt.Println("Name is :", name("pranali"))
	fmt.Println("Full Name is:",merge("pranali", "kirolkar"))
}


