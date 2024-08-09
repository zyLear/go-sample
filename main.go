package main

type Server interface {
	WithTimeout(int) Server
	WithTLS(bool) Server
	Start() Server
	Await() int
}

type HttpServer struct {
	timeout int
}

func (h HttpServer) WithTLS(b bool) Server {
	// no TLS support, ignore
	return h
}

func (h HttpServer) WithTimeout(i int) Server {
	h.timeout = i
	return h
}

func (h HttpServer) Start() Server {

	return h
}

func (h HttpServer) Await() int {
	return 0
}

type HttpsServer struct {
	HttpServer
	tlsEnabled bool
}

func (h HttpsServer) WithTLS(b bool) Server {
	h.tlsEnabled = b
	return h
}

//func main() {
//	HttpsServer{}.WithTimeout(10).WithTLS(true).Start().Await()
//}

type Animal interface {
	Eat()
}

type Mammal interface {
	Animal
	Lactate()
}

type Dog struct {
	Mammal
}

var _ Mammal = Dog{}
var _ Server = HttpsServer{}

func main() {
	d := Dog{}
	d.Eat()
}
