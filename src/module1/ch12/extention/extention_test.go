package extention__test

import (
	"fmt"
	"testing"
)

type pet struct {
}

func (p *pet) Speak() {
	fmt.Println("pet speak")
}

func (p *pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println("pet SpeakTo ", name)
}

func TestPet(t *testing.T) {
	p := new(pet)
	p.SpeakTo("fufeng")
}

type dog struct {
	p *pet
}

func (d *dog) Speak() {
	d.p.Speak()
}

func (d *dog) SpeakTo(name string) {
	d.p.SpeakTo(name)
}

func TestDog(t *testing.T) {
	d := &dog{}
	d.SpeakTo("dog")
}
