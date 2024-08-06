package main

import "fmt"

// Animal is type for our abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is the concerte factory for dogs
type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("Woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

// Cat is the concerte factory for cats
type Cat struct{}

func (c *Cat) Says() {
	fmt.Println("Meow")
}

func (c *Cat) LikesWater() bool {
	return false
}

// The following things are for readability
type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {

	// Careate one each of a DofFactory and CatFactory
	DogFactory := DogFactory{}
	CatFactory := CatFactory{}

	// Call the new method to create a Dog and Cat
	dog := DogFactory.New()
	cat := CatFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("A dog likes water: ", dog.LikesWater())
	fmt.Println("A cat likes water: ", cat.LikesWater())

}
