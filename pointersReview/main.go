package pointersreview

import "fmt"

func PracticePointer() {
	fmt.Println("Pointer Practice")
	age := 34
	fmt.Println("Age variable: ", age)

	// intialized a pointer
	// var agePointer *int

	// & - operator to get and store the address of a variable
	agePointer := &age
	fmt.Println("Pointer to the age variable address: ", agePointer)

	// use * to get the value of the pointer
	// this is called de-referencing a pointer
	fmt.Println("Age pointer value: ", *agePointer)

	adult_age := GetAdultYearsByPointer(agePointer)
	fmt.Println("Get Adult Year By Pointer: ", adult_age)
	fmt.Println("Age value: ", age)

	GetAdultYearsByOverriding(agePointer)
	fmt.Println("Age value after overriding it in function call: ", age)
}

// make a copy of the passed agrument age
func GetAdultYears(age int) int {
	return age - 18
}

// there is no copy being created for the passed agrument age
func GetAdultYearsByPointer(age *int) int {
	return *age - 18
}

// overriding the value stored in the pointer passed
func GetAdultYearsByOverriding(age *int) {
	*age = *age - 18
}
