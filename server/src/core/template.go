package server

import(
	"os"
	"fmt"
)

var DEFAULT_TEMPLATE_DIR string = "../../pages"

//TODO: Some sort of cacheing where in the response/file is stored in a table along with the last time the file was modified. then when the same request comes in, check to see last time the file was modified
//if the file has been modified after the entry in the table, read from file, otherwise just return the cached version
//This should speed up performance by eliminating read()'s on every request
func Render(fileName string, variables ...string)(*Response){
	data, err := os.ReadFile(DEFAULT_TEMPLATE_DIR + fileName)

	if(err != nil){
		fmt.Println("TODO: Handle template missing/etc errors")
	}

	fmt.Println("Returning: ")
	fmt.Println(string(data))

	return Respond(200, string(data))
}
