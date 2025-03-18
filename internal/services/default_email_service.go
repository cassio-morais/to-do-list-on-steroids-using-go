package services

import (
	"log"
	"time"
)

type DefaultEmailService struct{}

func NewDefaultEmailService() *DefaultEmailService {
	return &DefaultEmailService{}
}

func (e *DefaultEmailService) Send(body interface{}, email string) {
	time.Sleep(4 * time.Second)

	log.Printf("email to %s sent succesfully", email)
}
