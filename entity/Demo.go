package entity

import (
	"BcRPCCode1/constants"
	"fmt"
)

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

type user struct {
	hash string
}
func (u user) notify() { // 这里的接收者是值
	fmt.Println(u.hash)
}

func Demo() {
	u := user{constants.GETBESTBLOCKHASH}
	sendNotification(u) // 传递的是指针
}

