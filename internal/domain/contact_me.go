package domain

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type ContactMeService interface {
	SendMail(msg Message) error
}
