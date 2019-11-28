package interface__test

import "testing"

type Program interface {
	WriteWorld() string
}

type goProgram struct {
}

func (g *goProgram) WriteWorld() string {
	return "fmt.Println(\"world\")"
}

func TestInterface(t *testing.T) {
	//var p Program = new(goProgram)
	var p Program = &goProgram{}
	t.Log(p.WriteWorld())
}
