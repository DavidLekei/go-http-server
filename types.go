package main

type Header struct{
	Key string
	Value string
}

type KVPair interface{
	SetKey(key string)
	SetValue(value string)
	GetKey()(key string)
	GetValue()(value string)
}

type Route struct{
	path string
	callback func(req *Request)(*Response)
	variable string
}