package code11_3

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	t.Log("A")
}

func TestB(t *testing.T) {
	t.Log("B")
}

func TestAK(t *testing.T) {
	t.Log("AK")
}

func TestC(t *testing.T) {
	t.Log("C")
}

func normal() {
	fmt.Println("normal func")
}
