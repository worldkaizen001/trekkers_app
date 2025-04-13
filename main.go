package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (a Person) Ab() {
	// return "";
}
func DoSomething(ctx context.Context)  {
	// ... use ctx ...
	// ctx.Deadline()
	return;
}
type Speaker interface {
    Speak() string
}

type Dog struct {
	Breed string
	Age string
}

func (d Dog) Speak() string {
 return "woof"
}
func main() {

	//defer
	//make
	//slice
	//map  Done
	//function
	//method
	//struct
	//loops ramge and regular for loop
	//convert struct to map
	//convert map to struct
	//context
	//interfaces
	//net/http
	//os
	//bufio
	//bufio

	p := make(map[any]any)



	ar := []int{1, 2, 3, 3, 4, 5, 67, 789}
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// defer cancel()
// ctx.Done()

	v := map[string]int{
		"a": 2,
		"b": 3,
	}
   l := Dog{Breed: "Root", Age: "WW"}
   fmt.Println(l.Speak())
	p["v"] = true
	p["x"] = false

	fmt.Println(v)
	v["a"] = 5
	fmt.Println(v)
	fmt.Println(p)
	p["x"] = 50
	fmt.Println(p)

	h := Person{Name: "gdgd", Age: "ww"}
	

	fmt.Println(h)
	u, _ := json.Marshal(h)
	fmt.Println(string(u))

	for i, av := range ar {
		fmt.Println(i, av)
	}

	for i := 0; i < len(ar); i++ {
		if ar[i] == 2 {
			ar[i] = 600
			fmt.Println("We are at index 1 with vale 2")
		}

		fmt.Println(i, ar[i])
	}
	data := map[string]string{"name": "Go"}
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))


}
