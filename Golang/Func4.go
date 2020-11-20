
  
package main

import "fmt"

func main() {

	StudentRoll := make(map[int]string)
	StudentRoll[5] = "Sai"
	StudentRoll[12] = "Abhinav"
	StudentRoll[15] = "Mahesh"
	StudentRoll[20] = "Sohail"

	fmt.Println(StudentRoll[12])

	Employee := map[string]map[string]string{
		"Febin": map[string]string{
			"Position":   "Engineer",
			"EmployeeId": "20",
		},
		"Favas": map[string]string{
			"Position":   "Developer",
			"EmployeeId": "15",
		},
	}

	if temp, name := Employee["Abhinav"]; name {
		fmt.Println(temp["Position"], temp["EmployeeId"])
	}
}