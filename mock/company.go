package mock

type Company struct {
	Usher Talker
}

func NewCompany(t Talker) *Company{
	return &Company{
		Usher:t,
	}
}

func ( c *Company) Meeting(gusetName string)string{
	return c.Usher.SayHello(gusetName)
}
