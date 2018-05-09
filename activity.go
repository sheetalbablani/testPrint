package testPrint

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-Print")

const (
	Disp     = "Message"
	ovResult = "result"
)

// inputs : {Message}
// outputs: {result}
type PrintActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new PrintActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &PrintActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *PrintActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *PrintActivity) Eval(context activity.Context) (done bool, err error) {

	Message := context.GetInput(Disp).(string)
	fmt.Println("Hey beautiful ", Message)

	context.SetOutput(ovResult, nil)
	return true, nil

}
