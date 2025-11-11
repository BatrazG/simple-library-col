package notifications

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю на email %s сообщение: %s\n", e.Email, message)
}

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Оправляю на номер %s СМС-сообщение: %s\n", s.PhoneNumber, message)
}
