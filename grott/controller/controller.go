package types

import "github.com/michael-golfi/Grott/grott/activity"

type Controller interface {
	Post(msg *activity.Activity)
	HandleSystemMessage(msg *activity.Activity)
}
