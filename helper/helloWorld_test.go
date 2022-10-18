package helper

import (
	"fmt"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko!" {
		//unit test failed
		// t.Fail()
		t.Error("Text is Wrong Dumbass, Result Must be 'Hello Eko!'") //t.fail but you can put arguments inside it
	}
	fmt.Println("Test_HelloWorld Done")
}

func Test_HelloWorldAzie(t *testing.T) {
	result := HelloWorld("Azie")
	if result != "Hello Azie!" {
		// t.FailNow()
		t.Fatal("Text is Wrong you dickhead, Result Must be 'Hello Azie!'") //t.FailNow but you can put arguments inside it
	}
	fmt.Println("Test_HelloWorldAzie Done")
}
