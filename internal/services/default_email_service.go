package services

import (
	"log"
	"time"
)

type DefaultEmailService struct{}

func NewDefaultEmailService() *DefaultEmailService {
	return &DefaultEmailService{}
}

func (e *DefaultEmailService) Send(body any, email string) {
	time.Sleep(4 * time.Second)

	log.Printf("email with body: %+v sent to %s succesfully", body, email)
}
