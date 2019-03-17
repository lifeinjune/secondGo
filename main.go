package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//definition of struct call person
type person struct { //data structure that collection of data
	firstName string
	lastName  string
	contactInfo
	//contact   contactInfo // or just contacInfo can be used
}

func main() {
	/*alex := person{"Alex", "Anderson"}                //not recommended
	may := person{firstName: "May", lastName: "June"} //define person struc call may
	var ally person                                   // 0 value assigned to it which is "", for int 0, float 0, bool false
	fmt.Println(alex)
	fmt.Println(may)
	fmt.Println(ally)
	ally.firstName = "ally" //assign value to struct
	ally.lastName = "olive"
	fmt.Printf("%+v", alex) //%+v will give name and value of alex*/
	jim := person{
		firstName: "Jim",
		lastName:  "Jimmy",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94504,
		}, // for multi-line struct need to have comma on at the end of it
	}
	jimPointer := &jim //&variable: & is operator which gives the memory address of the variable
	jimPointer.updateName("Ji")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// go is pass by value language
//p is different variable and copies the value of the jim into it
//so even if i modify p value, jim value still same
/*func (p person) updateName(newFirst string) {
	p.firstName = newFirst
}*/
func (pointerToPerson *person) updateName(newFirst string) { //*type is pointer to the type NOT OPERATOR, receiver of pointer to the type
	(*pointerToPerson).firstName = newFirst //* give access to the pointer's value

}

// address to value by *address
// value to address by &value
