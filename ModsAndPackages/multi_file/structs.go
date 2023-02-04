package main


type PersonBuilder interface{
	SetPerson(p Person) (fp Person, e error)
}

type Person struct {
	name string
	age int
	address string
}

func SetPerson ( p Person) ( rp Person, e error){
	a := Person{"che", 34, "we"}
	e := error

	return a, e
}