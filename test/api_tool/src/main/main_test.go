package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_getit(t *testing.T) {

	b, i := getit()

	if !b {

		t.Error()
	}

	fmt.Printf("%v", i)
	//t.Error() // to indicate test failed
}

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}
