# streams
Like Java, stream API for processing of collection of elements with less code in Golang.

## Usage:

Lets say you have a slice of following struct:

```go
type Employee struct {
	Id     int
	Name   string
	Salary float64
	Age    int
}
```
And you need to have slice of salaries of all employees having Age > 30.

Have a look at the following code using streams package:


```go

package main

import (
	"fmt"
	"github.com/nikhilg-hub/streams"
)

func main() {

  empList := streams.New([]Employee{
		Employee{1, "Bob", 23000, 27},
		Employee{2, "Alice", 24000, 32},
		Employee{3, "Marina", 27000, 31},
		Employee{4, "Marie", 43224, 23},
		Employee{5, "John", 343322, 45},
	})
  
  salarySlice := []float64{}
  
	empList.Filter(func(e Employee) bool {
		return e.Age > 30
	}).Map(func(e Employee) any {
		return e.Salary
	}).ForEach(func(salary any) {
		salarySlice = append(salarySlice, salary.(float64))
	})
  
	fmt.Println(salarySlice)

}

```
Output :
```go
[24000 27000 343322]
```
