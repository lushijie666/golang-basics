package main

import "fmt"

func main(){
	s := server{"/get", "8080"}
	s.CheckRequest()
}


func (s *server) CheckRequest(){
	fmt.Printf("server uri is %s \n", s.uri)
	fmt.Printf("server port is %s \n", s.port)
}

type server struct {
	uri string
	port string
}
