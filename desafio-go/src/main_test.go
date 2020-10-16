package main

import "testing"

func TestGrettings(t *testing.T) {

	grettings := grettings() //somatória deveria retorna 10

	if grettings != "Code.education Rocks!" {
		t.Errorf("grettings () failed, expected %v, got %v", "Code.education Rocks!", grettings)
	}

}
