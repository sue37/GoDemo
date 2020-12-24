/*
A package to demonstrate how to use go doc.
	This package include function & structure.
	Doc for methods of the structure can be displayed as well.
*/
package DocTest

// DocTest is a demo struct
type DocTest struct {
	TestStr string
}

// CreateDocTest is a public function to create a DocTest
func CreateDocTest() {

}

// PublicMethod in DocTest
func (inst *DocTest)PublicMethod(name string) {

}

// Use go doc -all to generate the full doc