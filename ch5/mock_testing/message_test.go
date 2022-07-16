package message

import "testing"

type MockMessage struct {
	email, subject string
	body           []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body
	return nil
}

func TestAlert(t *testing.T) {
	msg := new(MockMessage)
	body := []byte("Critical Error")

	Alert(msg, body)

	if msg.subject != "Critical Error" {
		t.Errorf("Expected 'Critical Error', Got '%s'\n", msg.subject)
	}
}
