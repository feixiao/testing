package mock



import "fmt"

type Person struct{
	name string
}

func NewPerson(name string)*Person{
	return &Person{
		name:name,
	}
}


func (p *Person)SayHello(name string)string {
	return fmt.Sprintf("Hello %s, welcome come to our company, my name is %s",name,p.name)
}
