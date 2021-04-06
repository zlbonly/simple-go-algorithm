package main

type User struct {
	userName string
}

func test1() *User {
	a := User{}
	return &a
}
func main() {
	test1()
}
