package bot_dialog

import (
	"github.com/michael-golfi/Grott/grott/connector"
	"github.com/michael-golfi/Grott/grott/activity"
	"log"
	"github.com/michael-golfi/Grott/grott/storage"
)

type SimpleDialog struct {}

func (d SimpleDialog) MessageReceived(ctx *storage.DialogContext, msg *activity.Activity) {
	conn := connector.NewClientConnector(msg.ServiceUrl)
	headers := map[string]string{}
	msg.Text = "Hello User!"

	from := msg.From
	recipient := msg.Recipient
	msg.From = recipient
	msg.Recipient = from

	if _, err := conn.Respond(*msg, msg.Conversation.Id, msg.Id, headers); err != nil {
		log.Printf("Could not Respond to Message: %s", err.Error())
	}
}