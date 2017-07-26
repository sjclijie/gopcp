package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println("Sending user email to", u.name)
}

type duration int

func (d duration) pretty() string {
	return "hello"
	//return fmt.Sprintf("Duration: %d", *d)
}

func main() {

	user := user{name: "lijie", email: "sjclijie@vip.qq.com", }

	sendNotification(user)

	fmt.Println("===================\n")

	duration(44).pretty()
}

func sendNotification(n notifier) {
	n.notify()
}
