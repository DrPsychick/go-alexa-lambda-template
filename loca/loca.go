// Package loca contains all translations.
package loca

import (
	"github.com/drpsychick/go-alexa-lambda/l10n"
)

// keys of the project.
const (
	// Intents.
	DoSomething               string = "DoSomething"
	DoSomethingSamples        string = "DoSomething_Samples"
	DoSomethingTitle          string = "DoSomething_Title"
	DoSomethingText           string = "DoSomething_Text"
	DoSomethingSSML           string = "DoSomething_SSML"
	DoSomethingUserText       string = "DoSomething_User_Text"
	DoSomethingUserSSML       string = "DoSomething_User_SSML"
	DoSomethingTaskSamples    string = "DoSomething_Task_Samples"
	DoSomethingTaskElicitText string = "DoSomething_Task_Elicit_Text"
	DoSomethingTaskElicitSSML string = "DoSomething_Task_Elicit_SSML"

	// Slots.
	SlotTaskElicitTitle     string = "Slot_Task_Elicit_Title"
	SlotTaskElicitText      string = "Slot_Task_Elicit_Text"
	SlotTaskElicitSSML      string = "Slot_Task_Elicit_SSML"
	SlotWrongTaskElicitText string = "Slot_WrongTask_Elicit_Text"
	SlotWrongTaskElicitSSML string = "Slot_WrongTask_Elicit_SSML"

	// Slot types.
	TypeTask       string = "TaskName"
	TypeTaskName   string = "Task"
	TypeTaskValues string = "TaskName_Values"

	// Errors.
	RequestFailedTitle string = "RequestFailed_Title"
	RequestFailedText  string = "RequestFailed_Text"
	RequestFailedSSML  string = "RequestFailed_SSML"

	AMAZONStopSamples   string = "AMAZON.StopIntent_Samples"
	AMAZONHelpSamples   string = "AMAZON.HelpIntent_Samples"
	AMAZONCancelSamples string = "AMAZON.CancelIntent_Samples"
)

// Registry is the global l10n registry.
var Registry = l10n.NewRegistry()

func init() {
	// default first
	locales := []*l10n.Locale{
		enUS,
	}
	for _, l := range locales {
		if err := Registry.Register(l); err != nil {
			panic("registration of locale failed")
		}
	}
}
