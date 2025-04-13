package mail

import (
	"testing"

	"github.com/pawaspy/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1> Hello world</h1>
	<p> This is test message from Pawas Pandey. A hero. A good friend of sumit srivastav</p>`

	to := []string{"samarthdec12@gmail.com"}
	attachFiles := []string{"../Makefile"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
