package customerservice

// Describe the service

type Service interface {
	Healt() bool
	Customer(name string)string
}

// CustomerService implementation of the Service interface

type CustomerService struct{}

// Health implementation of the Service.

func (CustomerService) Health() bool {
	return true
}

// CustomerHello implementation of the Service.
func (CustomerService) Customer(name string) (custumerHello string) {
	custumerHello = "Hello " + name
	return
}