package usecase

type emailManager struct {
}

func NewEmailManager() EmailManager {
	return &emailManager{}
}

func (m *emailManager) Send(subject, addressee, body string) error {
	panic("implement me")
}
