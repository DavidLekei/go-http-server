package main

import(
	"fmt"
	//"errors"
	"strings"
	"encoding/gob"
	"bytes"
	"strconv"
)

var CRLF string = "\r"

type Request struct{
	Method string
	Target string
	Version string
	Headers map[string]string
	Body string
	headerCount int
	Params map[string]string
}

type Response struct{
	StatusLine string
	Headers string
	Body string
}

type Context struct{

}

type HttpRequest interface{
	Get()
	Post()
	Update()
	GetHeaderValue(header string)
	AddHeader(header string, value string)
	AddParams(paramString string)
	AddVariable(name string, value string)
	Print()
}

func CreateRequestFromBytes(bytes []byte, len int)(*Request){
	message := string(bytes[:len])

	lines := strings.Split(message, "\n")

	startLine := strings.Split(lines[0], " ")

	var r *Request = new(Request)
	r.Method = startLine[0]
	r.Target = startLine[1]
	r.Version = startLine[2]
	r.Headers = make(map[string]string)
	r.headerCount = 0
	r.Params = make(map[string]string)

	lineCount := 1
	line := lines[lineCount]

	for line != CRLF{
		header := strings.Split(line, ":")
		r.Headers[header[0]] = header[1]
		r.headerCount++
		lineCount++
		line = lines[lineCount]
	}


	return r
}

func CreateRequest()(*Request){
	return new(Request)
}

func Respond(statusCode int, body string)(*Response){
	//TODO: Get default response content-type from some kind of config
	responseHeaders := make(map[string]string)

	responseHeaders["X-Content-Type-Options"] = "nosniff"
	responseHeaders["Content-Type"] = "text/html"
	responseHeaders["Content-Length"] = strconv.Itoa(len(body))
	responseHeaders["Date"] = "TODO"

	return encodeResponse(statusCode, body, responseHeaders)
}

func JSON(statusCode int, body string)(*Response){
	responseHeaders := make(map[string]string)

	responseHeaders["Content-Type"] = "application/json"
	responseHeaders["Content-Length"] = strconv.Itoa(len(body))

	return encodeResponse(statusCode, body, responseHeaders)
}

func encodeResponse(statusCode int, body string, headers map[string]string)(*Response){
	
	//Update the 'accepts' header
	//headers["Content-Type"] = "application/json"

	//TODO: Status "Message" should be retrieved via a lookup table (an array of strings of size 600 or whatever, where the element at that status code number contains the message)
	//IE: statusLookup[200] = "OK"
	statusLine := fmt.Sprintf("HTTP/1.1 %d OK", statusCode)
	statusLine = statusLine + "\r\n"

	headerString := ""

	var keys []string
	for k := range headers{
		keys = append(keys, k)
	}

	for _, k := range keys{
		headerString = headerString + k + ":" + headers[k] + "\n"
	}
	headerString = headerString + "\r\n"

	//TODO: Not sure if a 2nd CRLF is necessary
	//headerString = headerString + CRLF

	return &Response{
		StatusLine: statusLine,
		Headers: headerString,
		Body: body,
	}
}

func (req Request) Get(){

}

func (req Request) Post(){
	
}

func (req Request) Update(){
	
}

func (req Request) AddHeader(header string, value string){
	req.Headers[header] = value
}

func (req Request) GetHeaderValue(header string)(value string){
	return req.Headers[header]
}

func (req Request) AddParam(key string, value string){
	req.Params[key] = value
}

func (req Request) AddParams(paramString string){

	params := strings.Split(paramString, "&")

	for _, p := range params{
		keyValuePair := strings.Split(p, "=")
		req.AddParam(keyValuePair[0], keyValuePair[1])
	}

}

func (req Request) Print(){
	fmt.Println("----------------\n\nRequest:")
	fmt.Println("Method: ", req.Method)
	fmt.Println("Target: ", req.Target)
	fmt.Println("Version: ", req.Version)
	fmt.Println("Headers (", req.headerCount, ") :")

	var keys []string
	for k := range req.Headers{
		keys = append(keys, k)
	}

	for _, k := range keys{
		fmt.Println(k, ":", req.Headers[k])
	}

	fmt.Println("Body: ", req.Body)
	fmt.Println("----------------")
}

// func (req Request) Encode()(bytes []byte){

// }

func (res Response) Encode()([]byte){
	responseString := res.StatusLine + res.Headers + res.Body + CRLF

	fmt.Println("StatusLine: ", res.StatusLine)
	fmt.Println("Headers: ", res.Headers)
	fmt.Println("Body: ", res.Body)


	//fmt.Println("responseString: ", responseString)

	return ([]byte(responseString))

	// var buf bytes.Buffer
	// enc := gob.NewEncoder(&buf)
	// err := enc.Encode(res)

	// if(err != nil){
	// 	fmt.Println("ERROR: Could not encode Response", err)
	// }

	// return buf.Bytes()
}

func (req Request) Decode(){
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(req)

	if(err != nil){
		fmt.Println("ERROR: Could not decode Request - ", err)
	}
} 

func (res Response) Decode(){

}

func (res Response) Print(){
	fmt.Print("Response:")
	fmt.Print("\n\tStatusLine: ", res.StatusLine)
	fmt.Print("\n\tHeaders: ", res.Headers)
	fmt.Print("\n\tBody: ", res.Body)
	fmt.Print("\n\t")
}