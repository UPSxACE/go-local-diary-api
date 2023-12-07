package controllers_test

import (
	"fmt"

	"github.com/UPSxACE/go-local-diary-api/server/controllers"
)

/*This is a test of Go Doc example feature

I'm wondering where this will show in the docs.
*/
func Example(){
	fmt.Println("test")
	// Output:
	// test
}

/*Hmm, let me test maps */
func ExampleLeTestMap(){
	leMap := controllers.LeTestMap()
	fmt.Println(leMap)
	// Output:
	// map[name:Test]
}


/*Hmm, let me test structs */
func ExampleLeTestStruct(){
	leStruct := controllers.LeTestStruct()
	fmt.Println(leStruct)
	// Output:
	// {a name 15 {hmm 20}}
}
