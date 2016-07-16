package bot_dialog

import (
	"github.com/michael-golfi/Grott/grott/types"
	"github.com/michael-golfi/Grott/grott/connector"
)

type SimpleDialog struct {

}

func (d SimpleDialog) MessageReceived(ctx *types.DialogContext, msg *types.Activity) error {
	conn := connector.NewClientConnector(msg.ServiceUrl)
	headers := map[string]string{}
	msg.Text = "Hello User!"

	from := msg.From
	recipient := msg.Recipient
	msg.From = recipient
	msg.Recipient = from
	_, err := conn.Respond(*msg, msg.Conversation.Id, msg.Id, headers)
	return err
}

func (d SimpleDialog) CalculateScore(msg *types.Activity) (int, error) {
	return 100, nil
}
