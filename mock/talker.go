package mock

type Talker interface {
    SayHello(word string)(response string)
}