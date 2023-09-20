//TODO: 
/*
	TODO:

		There should be 2 Route tables - one for GET and one for POST requests
		Then AddRoute() should be removed/refactored to 2 functions, GET() and POST(), which will add the route/callback to the respective table

		Needs to be able to handle the HTML template files loading in Javascript files as well... not sure how to do that at the moment
*/

package server

import(
	"fmt"
	"net"
	"errors"
	"strings"
)

var DEFAULT_PORT = 3456
var MAX_BUFFER_SIZE = 4096

var routeTable map[string]*Route

//TODO: in the future, add support for parsing RESTful type paths such as /user/<userId>/...
// func parsePath(path string, req *Request)(*Route){
// 	var key string

// 	pathOptions := strings.Split(path, "/")
// 	tryPath := ""

// 	for i, p := range pathOptions{
// 		if(i == 0){
// 			continue
// 		}

// 		tryPath = tryPath + "/" + p

// 		key = tryPath + req.Method
// 		if(routeTable[key] != nil){
// 			return routeTable[key]
// 		}
// 	}

// 	return nil
// }

func findRoute(req *Request)(*Route, error){

	// /users/32
	// /posts/security/some-title
	// /users/32?displayAs=table

	url := strings.Split(req.Target, "?")

	path := url[0]
	if(len(url) == 2){
		req.Target = path
		req.AddParams(url[1])
	}

	//var route *Route = parsePath(path, req)

	key := path + req.Method
	var route *Route = routeTable[key]


	if(route == nil){
		return nil, errors.New("Route - " + path + " - not found")
	}else{
		return route, nil
	}
}

func handleConnection(conn net.Conn){
	var inBuffer []byte = make([]byte, MAX_BUFFER_SIZE)
	//var outBuffer []byte = make([]byte, MAX_BUFFER_SIZE)

	bytesRead, readError := conn.Read(inBuffer)

	if(readError != nil){
		fmt.Println("handleConnection() - ERROR")
	}else{
		req := CreateFromBytes(inBuffer, bytesRead)

		route, routeError := findRoute(req)
		if(routeError != nil){
			fmt.Println("ERROR: ", routeError)
			//Check if there's a default "error" page/route defined, and if so, return that
		}else{

			var res *Response
			res = route.callback(req)
			conn.Write(res.Encode())
		}


		//conn.Write(outBuffer)
		//conn.Close(
	}
}

func listen(port int){
	var host string = fmt.Sprintf("localhost:%d", port)
	fmt.Println("Listening on: ", host)
	listener, listenError := net.Listen("tcp", host)

	if(listenError != nil){
		fmt.Println("TODO: Handle error")
	}

	defer listener.Close()

	for{
		conn, connError := listener.Accept()
		if connError != nil{
			fmt.Println("TODO: Handle error")
		}

		go handleConnection(conn)
	}

}

func checkPathForVariable(path string)(variable string, newPath string){
	if(strings.Contains(path, ":") == false){
		return "", path
	}
	
	var _newPath string

	pathSplit := strings.Split(path, ":")
	_newPath = pathSplit[0]
	_newPath = _newPath[:len(_newPath) - 1]
	//Example:
	//path = /users/:id/settings
	//should parse like:
	//strings.Split(path, ":")[1] = "id/settings"
	//strings.Split("id/settings", "/")[0] = "id"
	return strings.Split(pathSplit[1], "/")[0], _newPath
}

func AddRoute(path string, methods []string, callback func(req *Request)(*Response)){
	if(routeTable == nil){
		createRouteTable()
	}

	for _, m := range methods{

		variable, path := checkPathForVariable(path)

		fmt.Println("Adding New Route with Path: ", path, " - Variable: " , variable)
		
		//TODO: I do not like this as a key, but I'm not exactly sure what to do better at the moment. This should DEFINITELY be changed in the future.
		key := path + m
		route := &Route{path: path, method: m, callback: callback, variable: variable}
		routeTable[key] = route
	}
}


//TODO:Take in a config
//Config can check for things like:
//Only allow certain http/https versions
//Only allow requests for certain Content-Type's
//Run on a specific port
//Etc
func Run(port ...int){
	fmt.Println("Starting HTTP Server...")


	if(port == nil){
		listen(DEFAULT_PORT)
	}else{
		listen(port[0])
	}

	fmt.Println("Shutting down server")
}

/*

	INTERNAL FUNCTIONS

*/

func createRouteTable(){
	routeTable = make(map[string]*Route)
}

func printRouteTable(){
	var keys []string

	for k := range routeTable{
		keys = append(keys, k)
	}

	for _, r := range keys{
		fmt.Println("Route: ", routeTable[r].path, " - ", routeTable[r].method, " - ", routeTable[r].callback)
	}
}