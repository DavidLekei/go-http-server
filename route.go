package main

import(
	"fmt"
	"strings"
	"errors"
)

type Route struct{
	path string
	callback func(req *Request)(*Response)
	variable string
}

func checkForVariableMatch(path string, routeTable map[string]*Route)(*Route, string){

	var route *Route = nil

	lastSlash := strings.LastIndex(path, "/")

	//if path = "/user/32", then pieces will be ["user", "32"]
	firstHalf := path[:lastSlash]
	secondHalf := path[lastSlash:len(path)]

	fmt.Println("First Half: ", firstHalf, " - Second Half: ", secondHalf)

	//Check to see if the route exists without the last piece (aka no variable)
	route = checkForExactMatch(firstHalf, routeTable)

	//If it does exist, add the variable
	if(route != nil){
		return route, strings.Trim(secondHalf, "/")
	}

	// pathToCheck := "/" + pieces[0]

	// for i := 0; i < len(pieces); i++{
	// 	pathToCheck = pieces[i]
	// }

	return route, ""
}

func checkForExactMatch(path string, routeTable map[string]*Route)(*Route){
	var route *Route

	route = routeTable[path]

	if(route != nil){
		fmt.Println("Exact match for - ", path, " - EXISTS")
		return route
	}

	return nil
}

func FindRoute(req *Request)(*Route, error){

	//the path /user/32 comes in

	var routeTable map[string]*Route

	//it does not contain a ? character, so url will be set to ["/user/32"]
	url := strings.Split(req.Target, "?")

	//path will be set to "/user/32"
	path := url[0]
	//If there had been parameters, they'd be added here
	if(len(url) == 2){
		req.Target = path
		req.AddParams(url[1])
	}

	//Set the corresponding request type table to search for a route
	if(req.Method == "GET"){
		fmt.Println("Checking GET table for: ", path)
		routeTable = getRouteTable
		printRouteTable(routeTable)
	}else if(req.Method == "POST"){
		routeTable = postRouteTable
	}

	//See if there's an exact match to the path "/user/32"
	route := checkForExactMatch(path, routeTable)

	//If no exact match found, try again but by splitting via "/" and using variables
	if(route == nil){ //no exact match found
		route, variable := checkForVariableMatch(path, routeTable)

		if(variable != ""){
			//Add param to request
			req.AddParam("Variable", variable) //TODO: The param name should be the name provided by the user in the Route, instead of "Variable"
											   //	   For example, if the route is "/user/:id", then the param name should be "id"
			return route, nil
		}
	}

	if(route == nil){
		return nil, errors.New("Route - " + path + " - not found")
	}else{
		return route, nil
	}
}