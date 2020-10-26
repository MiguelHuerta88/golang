package main

import "fmt"

// define a struct
/*type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}*/

// we can also do this, this will create field contactInfo of contactInfo type
type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	/* Different ways to create person instances */
	/*p1 := person{"Alex", "Anderson"}
	fmt.Println(p1)

	p2 := person{firstName: "Miguel", lastName: "Huerta"}
	fmt.Println(p2)

	p3 := person{}
	p3.firstName = "Susie"
	p3.lastName = "Espino"
	//fmt.Println(p3)
	fmt.Printf("%+v", p3) // this will print out the properties of the struct

	var p4 person
	p4.firstName = "Hello"
	p4.lastName = "World"
	fmt.Println(p4)*/

	p1 := person{}
	p1.firstName = "Miguel"
	p1.lastName = "Huerta"
	p1.contactInfo.email = "guelme88@gmail.com"
	p1.contactInfo.zip = 91790
	fmt.Printf("%+v", p1)

	p2 := person{
		firstName: "susie",
		lastName:  "Espino",
		contactInfo: contactInfo{
			email: "miguel.huerta@disqo.com",
			zip:   91790,
		},
	}

	/**
	 * POINTERS
	 * &p give your reference to address
	 * *p allows your to modify that value
	 */
	//p2Pointer := &p2
	p2.updateName("Operation") // we can do this as a shortcut. go will allow us to call this receiver with a pointer or how we have it
	p2.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}
