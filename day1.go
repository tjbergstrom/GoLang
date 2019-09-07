package main

import "fmt"

var str = "hellow world"
var str2 = "hello" + " " + "world"
var x = int64 = 1
int i = intint u = uint
var -64 = int64
var u64 = uint64
var i32 = int32
type MyInt int
var i MyInt - MyInt(x)
var f32 = float32 = 1.5
var f64 = float64 = float64(f32)
const n = 3 // 


func main(){
	fmt.Println("hello world")
	fmt.Println(str)
	fmt.Println(x)
	fmt.Printf("%v %t %T \n", str, str, str)

	_ = x
	_ = f64 // this allows you to run if the var hasnt been used

	str := "hello" // declare and initialize... 

	local variable overrides global

	var str []string // a slice of strings

	var myInts [3]int // space for 3 ints - array

	if myInts == nil {
		myInts = make([]int, 3)
		myInts = make([]int, 0, 3) // 0 is length, 3 is capacity?
	}
	for i := range myInts {
		myInts[i] i
	}
	fmt.Println(myints)

	for i := 0; i < cap(myInts); i++ {
		myInts = append(myInts, i)
		myInts[i] = i
	}
	fmt.Println(myInts, len(myInts), cap(MyInts))

	fmt.Println(myInts[:2]) // prints these indexis
	fmt.Println(myInts[1:3])



}




func main(){
	if true && true || true {

	}

	// no while loops just for

	var myInts = []int{0, 1, 2, 3}
	var myInts = []int{} // initialized and not nil but nothing there
	for i := 0; i < 5; i++{

	}

	for i:= 0; i < len(myInts); i++ // these two are the same
	for i:= range myInts {

	}
	for i, v := range myInts { // v is the value of int at that index
		println(i, v) // prints index i 0 1 2 3 4, v prints the value a copy
		myInts[i]++ //increments the values in the array (after printing in the above line)
		v++ // doesnt modify the array
		fmt.println(i, *v) // pointer to that specific value in the array
		(*v)++ // the pointer increments values in the array
	}


}

// a struct of array or an array of struct
func Ranger(n int) []struct{}{
	return make([]struct{}, n)
}


func ranger(n, m, x int, f, t float64) []int {
	return nil // needs a return value, dont need to use args tho

}

func ranger(n int) (ints []int, err error) { //naming return values
	defer finc(){ // defer a function ...it will still run even if panic
		if err != nil {
			int = nil
		}
	}()
	int = make([]int, n) // gets overide by return nill statement
	return nil, nil // needs a return value, dont need to use args tho

	if n%2 == 0 {
		return nil, print("not odd")
		err := fmt.Errorf("not odd") // compiler error you overwrote the argument and didnt use it
		panic("not odd") // still runs and differs - prints out print lines until the panic executes
		// panic is for errors that should never happen
	}
	return make([]int, n), nil // return slice with odd lenth
	// main {
		int, err := oddSlice(5)
		if err != nil {
			log.Fatal(err)
		}
		for i  range int {
			ints[i] = i + 1
		}
		print(ints)
	}
	
}

// maps
var Ages map[string]int
var LastName map[string]string // cant assign one map to the other because different typ
var myMap map[key]value
func main(){
	Ages = make(map[string]int) // initialized, now you can write to and read from
	Ages["Adam Levy"] 26
	Ages["tim"] 30
	Ages ["Carol"] = 10
	for name, age := range Ages // name is key, age is value
		print(name, age)
		// there is no order to the map, it prints out in a different order every time
		// use different structure, like sorted array, of you need sorted
		// range thru this, sort it, put it in slice
	for name, age := range ages {
		print(name, age)
		ages["adam"]++
	}
	println(ages)

	age, ok := ages["sally"]
	if !ok {
		print("no age")
	} else {
		print(age)
	}

	for name ages range ages	
		print name age
			(*ages["adam"])++ // this prints out the address, deref the pointer
	print ages
}

var LastNames = map[string]string {
	"adam": "levy", // note commas at end of line
	"bob": "fish",
}



//structs
type MyStruct struct {} // takes up no space

type Person struct {
	FirstName 	string
	LastName 	string
	Age 		int
	Pets 		[]Pet
	Mom			&mom
}



type Pet struct {
	Bane string
	Hairless bool
}

func main(){

	mom := Person {
		FirstName: "Liz",
		LastName: "fish",
	}

	me := Person{} // expression, liter, uninitialized
	me := Person{
		FirstName: "Adam", 
		LastName: "LEvy",
	}
	me.Age = 26
	dog := Pet {Name: "Sport"}
	me.Pets = append(me.Pets, dog)
	print(me)

	for i := range stuff {
		thing := &stuff[i]
		(*thing)++ //pointers work like c++ pass by ref
	}
}




