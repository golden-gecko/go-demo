package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println(primitive.ObjectID{})

	/*
	for i := 0; i < 100; i++ {
		t0 := time.Now().UTC()
		t1 := t0.Format("2006-01-02 15:04:05.000000")

		fmt.Println(t1)
	}
	*/
}
