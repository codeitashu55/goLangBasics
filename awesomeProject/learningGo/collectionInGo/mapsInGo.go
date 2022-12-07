package collectionInGo

import (
	"fmt"
)

func BasicsOfMap() {

	//basic extended syntax
	employee := map[string]int{
		"ashu": 103,
		"bbhi": 101,
	}
	employee["callu"] = 102

	fmt.Println("contents of map -> ", employee)

	// retrieving values form map using key
	name := "ashu"
	println(name, "salary is ", employee[name])
	//the string is case sensitive and if the values is not present zero value of the type is returned
	println("Ashu salary is", employee["Ashu"])

	//code to check if a key is present in map or not
	// the map[key] -> kind of had an polymorphism feature
	value, isPresent := employee[name]

	if isPresent {
		println(name, "is present with salary", value)
	} else {
		println(name, "is not present in the list")
	}

	//iterating over all the elements of map -> the order in which we will get the element is not fixed they keep on changing
	fmt.Print("\nmap before deletion:->\n")
	printMap(employee)

	//deleting key form the map
	delete(employee, name)
	fmt.Print("\nmap after deletion:->\n")
	printMap(employee)

}

func printMap(employee map[string]int) {
	for key, value := range employee {
		fmt.Printf("employee[%s] = %d \n", key, value)
	}
}

type employee struct {
	country string
	salary  int
}

func MapsWithStructure() {
	employeeInfo := map[string]employee{
		"emp1": {
			"India",
			1000,
		},
		"emp2": {
			"Ghana",
			1020,
		},
		"emp3": {
			"USA",
			001,
		},
		"emp4": {
			"USA",
			001,
		},
	}

	for key, value := range employeeInfo {
		fmt.Printf("%s is from %s and earns $%d per month\n", key, value.country, value.salary)
	}

	//length of map
	fmt.Printf("\nlength of map is %d", len(employeeInfo))

	//map works in a similar way as the array so if assign them to other variable they will point to same data-structure underneath
	ref := employeeInfo
	ref["emp1"] = employee{
		"Bangladesh",
		100,
	}

	println()
	for key, value := range employeeInfo {
		fmt.Printf("%s is from %s and earns $%d per month\n", key, value.country, value.salary)
	}
}
