package inspeqtor

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"strings"
	"testing"
)

func makeAction(actionName, notifType string, config map[string]string) (Action, error) {
	return Actions[actionName](nil, &AlertRoute{"", notifType, config})
}

func mockService(name string) *Service {
	return &Service{name, 0, 0, nil, nil, nil, services.MockInit()}
}

func TestRestartAlert(t *testing.T) {
	s := mockService("foobar")
	res, err := Actions["restart"](s, nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGmailNotifier(t *testing.T) {
	action, err := makeAction("alert", "gmail", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"email":    "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, action)
}

func TestEmailNotifier(t *testing.T) {
	action, err := makeAction("alert", "email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"hostname": "smtp.example.com",
		"email":    "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, action)
}

func TestInvalidNotifier(t *testing.T) {
	action, err := makeAction("alert", "emaul", map[string]string{})
	assert.NotNil(t, err)
	assert.Nil(t, action)
}

func TestMissingEmailNotifier(t *testing.T) {
	action, err := makeAction("alert", "email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"email":    "mike@example.org",
	})
	assert.NotNil(t, err)
	assert.Nil(t, action)
}

func TestEmailTrigger(t *testing.T) {
	svc := Service{"mysql", 0, services.Down, nil, nil, metrics.NewProcessStore(), nil}
	alert := &Alert{
		&Rule{&svc, "memory", "rss", GT, 64 * 1024 * 1024, 0, 1, 0, Ok, nil},
	}

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailNotifier, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "[mysql]") > 0, "email does not contain expected content")
		assert.True(t, strings.Index(content, "memory(rss)") > 0, "email does not contain expected content")
		return nil
	})
	assert.Nil(t, err)
}

func validEmailSetup() *EmailNotifier {
	return &EmailNotifier{
		"mike", "fuzzbucket", "smtp.gmail.com", "mike@example.org", "mperham@gmail.com"}
}