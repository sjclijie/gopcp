package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("Sending user email to", u.name)
}

func main() {

	user := user{name: "lijie", email: "sjclijie@vip.qq.com", }

	sendNotification(user)
}

func sendNotification(n notifier) {
	n.notify()
}
