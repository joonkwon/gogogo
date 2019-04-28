package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := flag.String("csv", "employee.csv", "csv filepath to load")
	flag.Parse()
	f, err := os.Open(*filepath)
	if err != nil {
		log.Fatalf("failed to open file: %s; %v\n", *filepath, err)
	}
	defer f.Close()

	employees := []Employee{}

	reader := csv.NewReader(f) //could be csv.NewReader(bufio.NewReader(f))
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("failed to read line: %v\n", err)
		}
		age, err := strconv.Atoi(strings.TrimSpace(line[2]))
		if err != nil {
			log.Printf("failed to convert string to age(int): %v", err)
		}
		salary, err := strconv.Atoi(strings.TrimSpace(line[3]))
		if err != nil {
			log.Printf("failed to convert string to salary(int): %v", err)
		}

		employee := Employee{
			FirstName: strings.TrimSpace(line[0]),
			LastName:  strings.TrimSpace(line[1]),
			Age:       age,
			Salary:    salary,
		}
		employees = append(employees, employee)
	}

	jsonByte, err := json.Marshal(&employees)
	if err != nil {
		log.Printf("failed to marshal to json: %v", err)
	}
	fmt.Println(string(jsonByte))
}

type Employee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Salary    int    `json:"salary"`
}

func (e Employee) MarshalJSON() ([]byte, error) {
	// EmployeeAlias type to strip off method(including MarshalJSON) from Employee struct
	// see http://choly.ca/post/go-json-marshalling/
	type EmployeeAliase Employee
	return json.Marshal(&struct {
		Salary string `json:"salary"`
		EmployeeAliase
	}{
		Salary:         "$" + strconv.Itoa(e.Salary),
		EmployeeAliase: EmployeeAliase(e),
	})
}

type Employess struct {
	Employess []Employee `json:"employees"`
}
