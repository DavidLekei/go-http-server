package main

func (header Header) SetKey(key string){
	header.Key = key
}

func (header Header) SetValue(value string){
	header.Value = value
}

func (header Header) GetKey()(value string){
	return header.Key
}

func (header Header) GetValue()(value string){
	return header.Value
}