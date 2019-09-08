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


// OOP inheritence, will have the variables name age 
// but doesnt inherit member functions like grow older
type Fireperson Person {
	P Person // doesnt inherit structs
	Person // inherits member function
	FirstName string // this is allowed
	Age int64 //this is allowd
}

funct (fp FirePerson) Job() string {
	return "youre on fire"
	PersonNumFires int
	isLit bool
	
}

func (fp *FirePerson) ClockInTime(t time.Time) error {

}

func (fp *FirePerso) DoWork error {
fp.
}


type Worker interface {
	DoWork() error
	Job() string
	ClockOut(t time.Time) error
	ClockIn(t time.Time) error

}

//structs
type MyStruct struct {} // takes up no space

type Person struct {
	FirstName 	string
	LastName 	string
	Age 		int
	Pets 		[]Pet
	Mom			*Person
}



type Pet struct {
	Bane string
	Hairless bool
}

//member function, declare outside of main block
func (p Person) String() string { // extra bit before the function identifier is how you tie it to the person type
	return fmt.Sprintf("Person{Name: %v %v, Age: %v, Mom's Name: %v}",
	p.firstName, p.LastName, p.Age, p.Mom.FirstName, [.Mom.LastName])
}

func (p Person) GrowOlder(years int) error {
	p.Age += years // this is a COPY passed by copy so its only changing the local copy // like p->Age
					// needs to accept pointer to person to change original
					// called a popinter receiver 
					func (p *Person) GrowOlder(years int) error {
	if p.Age > 99 { // like p->Age
		return fmt.Errorf("u r ded")
	}
	return nil
}

func main(){

	var firePerson = FirePerson{Person: Person{FirstName: "Bob"}}
	var wk Worker = &firePerson
	print(
		if err := wk.Clockin(time.Now()); err != nil {
			log.Fatal(err)
		}
		if err := wk.DoWork(); err != nil {
			log.Fatal(err)
		}
		if err := wk.ClockOut(time.Now().Add(time.Hour)); err != nill {
			log.Fatal(err)
		}
		firePerson := wk.(*FirePerson)
		fmt.Println(firePerson)
	)


	mom := Person {
		FirstName: "Liz",
		LastName: "fish",
	}
	mom.FirstName // automatically dereferences pointer (wouldnt work in c++)
	// mom->firstName would be c equivalent


	me := Person{} // expression, liter, uninitialized
	me := Person{
		FirstName: "Adam", 
		LastName: "LEvy",
		Mo:	mom,
	}
	me.Age = 26
	dog := Pet {Name: "Sport"}
	me.Pets = append(me.Pets, dog)
	print(me)

	str := me.String()
	fmt.print(str)

	var err errorfor err == nil { // equiv to while true / while not dead
		fmt.Println(me)
		err = me.GrowOlder(1)
	}


	fireMan := FirePerson{Person: me}
	fireMan.FirstName // get empty string because it prefers the string in the firePerson struct
	fireman.Person.FirstName 


	for i := range stuff {
		thing := &stuff[i]
		(*thing)++ //pointers work like c++ pass by ref
	}
}








package main

import {
	"fmt"
	"log"
	"time"
}



func main() {
	var work chan int // make the channel before you use it
	work = make(chan int)
	jobs := make(chan int) // same as above two lines
	jobs := make(chan int, 5) // makes 5 jobs
	for i := 0;  i<4; i++ {
		go doWork(jobs) // pass in the channel so worker can read from it
		// this creates workers
	}
	for i := 0;  i<4; i++ {
		jobs <- i // write to the channel - every time it hits this, it blocks until worker does the line in doWork
		// this is the jobs the workers do
		// writing to the channgel is blocking until has been read
	}
	close(jobs)

	time.Sleep(3 * time.Second)
}

// job does the work
func doWork(job chan int) { //receives channel arg
func doWork(w int, jobs chan int) {
	n := <-job //reads from channel and prints result
	fmt.Println("n") // doesnt print in order becasue its waiting for any of the workers to read and do this
	//second way of doin it
	for n := range jobs {
		print("%v says %v", w, n)
	}
	
}









Go Imports
