package contracts

type EmailService interface {
	Send(body any, email string)
}
