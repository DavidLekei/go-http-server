package main

import(
	"fmt"
)

func addUser(req *Request)(*Response){
	fmt.Println("addUser called")
	return JSON(200, "user: {\"id\": 1, \"name\": \"Test User\"}")
}

func getUsers(req *Request)(*Response){

	if(req.Method == "GET"){
	}

	if(req.Method == "POST"){
		fmt.Println("Handling POST Request")
	}

	return Respond(200, "\"response\": \"This is a test response body\"")
}

func home(req *Request)(*Response){
	fmt.Println("Home Route")

	return Render("/home/home.html")
}

func main(){
	fmt.Println("----")

	Get("/home", home)
	Get("/user", getUsers)
	Post("/user", addUser)

	Run()
}