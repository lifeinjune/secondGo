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
	contact   contactInfo
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
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94504,
		}, // for multi-line struct need to have comma on at the end of it
	}
	fmt.Printf("%+v", jim)

}
