package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func dataPerson() {

	values := &Person{
		Name: "Test 123",
		Age:  24,
		Exercise: &Exercise{
			Fitness: 186,
			Pilates: 174,
		},
	}

	data, err := proto.Marshal(values)
	if err != nil {
		log.Fatal("Marsal Convert Error!", err)
	}

	fmt.Printf("gelen data %v\n", data)

	newData := &Person{}

	err = proto.Unmarshal(data, newData)

	if err != nil {
		log.Fatal("unmarsing conver error", err)
	}

	fmt.Println("Yukarıda 2li oldu , şimdi unmarsal convert yapip datayı bize vericek")
	fmt.Println("***********************************")
	fmt.Printf("User: %v\n", newData.GetName())
	fmt.Printf("Age: %v\n", newData.Age)
	fmt.Printf("Fitness Exercise: %v Day\n", newData.Exercise.Fitness)
	fmt.Printf("Pilates Exercise: %v Day\n", newData.Exercise.Pilates)
}

func main() {

	fmt.Println("Burası Main Başlangıç Start Code")
	dataPerson()
}
