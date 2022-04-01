package main


import (
	"fmt"
	"bytes"
	"os"
)


//Here we have to caluclate density of Metal and gas. SO to avoid duplicacy of function
//We can use Interface having common Density function, which each struct will implement.
type Dense interface{
	Density() float64
}

type Metal struct{
	mass float64
	volume float64
}

type Gas struct{
	pressure float64
	temperature float64	
}

type Shape interface{
	area() int
}

type Triangle struct{
	height int
	base int
}

type Square struct{
	side int
}

func (t Triangle) area() int{
	return (t.base*t.height)
}

//Using pointer receiver
func (t *Triangle) angles() int{
	return (t.base*t.height)
}

func (s Square) area() int{
	return (s.side*s.side)
}

//structure function, first m *Metal indicates this a functon of Metal struct called using pointer variable
//This is called as pointer receiver
func (m *Metal) Density() float64{
	return (m.mass * m.volume)
}

func (g *Gas) Density() float64{   //both Density functions are of type interface receiving object of Metal and Gas
	return (g.pressure * (g.temperature+273))
}


//compare the density of two metals
// func isDenser(a, b *Metal) bool {
// 	return (a.Density() > b.Density()) //calling density() of struct using pointer variable
// }

//Using interface, it will never be of type pointer
func isDenser(a, b Dense) bool {
 	return (a.Density() > b.Density()) //calling density() of struct using pointer variable
}

//Usage of empty interface 
func describe(str interface{}){
	//Type switch to know the type variable like type assertion
	switch str.(type){  //this gives type of str  //switch v := str.(type), can use v varaible as well
	case int:
		fmt.Println("str is integer, value is ", str)
	case string:
		fmt.Println("str is string, value is ", str)
	}
}


func main(){

	describe(42)
	describe("Hello")
	gold := Metal{500, 10}
	silver := Metal{200, 20}

	fmt.Println("Is silver denser than gold - ", isDenser(&silver, &gold))

	oxygen := Gas{
		pressure: 20,
		temperature: 10,
	}

	hydrogen := Gas{
		pressure: 1,
		temperature: 5,
	}

	fmt.Println("Is oxygen denser than hydrogen - ", isDenser(&oxygen, &hydrogen))

	var buf bytes.Buffer
	fmt.Fprintf(os.Stdout, "Hello")   //Here Fprintf is interface function which can accept stdout or byte buffer etc...
	//It is defined in fmt package and used here. First param is of type io.Writer. This provides Write method which is implemented by os and bytes package
	//io.Writer provides an abstraction of all types to which bytes can be written, 
	//which includes Files, memory buffers, Network connections, HTTP Clients.
	fmt.Fprintf(&buf, "Good Morning")
	fmt.Println(buf)

	//Create variables of concreate types using interface
	s := []Shape{
		Triangle{2,3},
		Square{5},
	}

	s1 := Triangle{1,2}

	p := &s1
	
	fmt.Println("Area of triangle ", s[0].area(), p.angles())
	fmt.Println("Area of Square ", s[1].area())

	//get the concrete type 
	conType, ok := s[0].(Triangle)  //check if shape object is Triangle, this is type assertions
	if ok {
		fmt.Println("Shape is triangle.", conType.base)
	}
}