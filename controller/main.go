package controller

import "github.com/ant0ine/go-json-rest/rest"

type Main struct{
}

func NewMain() *Main{
	return &Main{}
}

func (self *Main) EchoHandler(w rest.ResponseWriter, r *rest.Request){
	w.WriteJson(map[string]string{"Body": "Api is working"})
}