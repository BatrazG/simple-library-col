package main

<<<<<<< HEAD
import "fmt"

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю на email %s сообщение: %s\n", e.Email, message)
=======
import (
	"fmt"
)

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
>>>>>>> upstream/main
}

type SMSNotifier struct {
	PhoneNumber string
}

<<<<<<< HEAD
func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Оправляю на номер %s СМС-сообщение: %s\n", s.PhoneNumber, message)
}

type Notifier interface {
	Notify(message string)
=======
func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю email на %s: %s\n", e.EmailAddress, message)
}

func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Отправляю SMS на номер %s: %s\n", s.PhoneNumber, message)
>>>>>>> upstream/main
}
