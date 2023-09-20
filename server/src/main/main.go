package main

import(
	"fmt"
	"go-http-server/server/server"
)

func getUsers(req *server.Request)(*server.Response){

	if(req.Method == "GET"){
	}

	if(req.Method == "POST"){
		fmt.Println("Handling POST Request")
	}

	return server.Respond(200, "\"response\": \"This is a test response body\"")
}

func home(req *server.Request)(*server.Response){
	fmt.Println("Home Route")

	return server.Render("/home/home.html")
}

func main(){
	fmt.Println("----")

	server.AddRoute("/users/:userId", []string{"POST", "GET"}, getUsers)
	server.AddRoute("/home", []string{"GET"}, home)

	server.Run()
}