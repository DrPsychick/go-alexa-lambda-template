package loca

import (
	"github.com/drpsychick/go-alexa-lambda/l10n"
	"github.com/drpsychick/go-alexa-lambda/ssml"
)

var enUS = &l10n.Locale{
	Name: "en-US",
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        {"My Skill"},
		l10n.KeySkillDescription: {"Description of your skill."},
		l10n.KeySkillSummary: {
			"Skill that enables the user to ...",
		},
		l10n.KeySkillKeywords: {
			"skill", "template",
		},
		l10n.KeySkillSmallIconURI: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_small.png", //nolint:lll
		},
		l10n.KeySkillLargeIconURI: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_large.png", //nolint:lll
		},
		l10n.KeySkillPrivacyPolicyURL: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		},
		// Error: privacyAndCompliance.locales.en-US
		// - object instance has properties which are not allowed by the schema: ["termsOfUse"]
		// l10n.KeySkillTermsOfUse: []string{
		//	"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		// },
		l10n.KeySkillTestingInstructions: {
			"Instructions on how to verify that the skill works.",
		},
		// Important: make sure that your invocation suits each locale!
		l10n.KeySkillInvocation: {"my demo skill"},
		l10n.KeySkillExamplePhrases: {
			"Alexa, open my demo skill and do something with package",
			"Alexa, process that thing with my demo skill",
			"Alexa, tell my demo skill to start chores",
			// only 3 allowed!
		},

		// default intents
		// Intent "AMAZON.StopIntent"
		l10n.KeyStopTitle: {"Ending"},
		l10n.KeyStopText:  {"End.", "Good bye.", "See U!"},
		l10n.KeyStopSSML:  {ssml.Speak("Good bye.")},
		// Intent "AMAZON.HelpIntent"
		l10n.KeyHelpTitle: {"Help"},
		l10n.KeyHelpText: {
			"You can do something by saying: make XXX or prepare XXX",
			"Try saying: 'make coffee' or 'sudo make me a sandwich'",
		},
		l10n.KeyHelpSSML: {
			ssml.Speak("You can do something by saying: 'make coffee' or 'prepare lunch'."),
			ssml.Speak("Try saying: 'make ...' and then add your task name."),
			ssml.Speak("Try saying " +
				ssml.UseVoice("Justin", "sudo make me a sandwich"),
			),
		},
		l10n.KeyCancelTitle: {"Abort"},
		l10n.KeyCancelText:  {"Ok, aborting now.", "Alright, rolling back."},
		l10n.KeyCancelSSML:  {ssml.Speak("Ok I'm rolling back now.")},

		// launch request
		l10n.KeyLaunchTitle: {
			"Greeting",
		},
		l10n.KeyLaunchText: {
			"Hello!",
			"Hi!",
			"Yes?",
			"Welcome to my demo skill",
		},
		l10n.KeyLaunchSSML: {
			ssml.Speak("<voice name=\"Salli\">Hello!</voice>"),
			ssml.Speak("<emphasis level=\"strong\">Hi!</emphasis>"),
			ssml.Speak(ssml.UseVoice("Kendra", "Welcome to my demo skill")),
		},

		// standard errors
		l10n.KeyErrorTitle:                   {"Error"},
		l10n.KeyErrorText:                    {"An error occurred: %s"},
		l10n.KeyErrorSSML:                    {ssml.Speak("An error occurred.")},
		l10n.KeyErrorNotFoundTitle:           {"Not found error"},
		l10n.KeyErrorNotFoundText:            {"An element was not found: %s"},
		l10n.KeyErrorNotFoundSSML:            {ssml.Speak("A not found error occurred.")},
		l10n.KeyErrorLocaleNotFoundTitle:     {"Locale missing"},
		l10n.KeyErrorLocaleNotFoundText:      {"Locale for '%s' not found!"},
		l10n.KeyErrorLocaleNotFoundSSML:      {ssml.Speak("The locale '%s' is not supported.")},
		l10n.KeyErrorNoTranslationTitle:      {"Translation missing"},
		l10n.KeyErrorNoTranslationText:       {"No translation found for '%s'!"},
		l10n.KeyErrorNoTranslationSSML:       {ssml.Speak("No translation found for '%s'!")},
		l10n.KeyErrorMissingPlaceholderTitle: {"Placeholder missing"},
		l10n.KeyErrorMissingPlaceholderText:  {"A placeholder is missing in key: %s"},
		l10n.KeyErrorMissingPlaceholderSSML:  {ssml.Speak("A placeholder is missing in key '%s'")},

		// Custom errors
		RequestFailedTitle: {"Request to external service failed"},
		RequestFailedText:  {"Could not %s server %s: request failed. Please try again."},
		RequestFailedSSML: {ssml.Speak(
			"My external connection could not %s server %s. Please try again. " +
				"If that does not solve your problem, please contact support",
		)},

		// Intent: "DoSomething"
		DoSomethingSamples: {
			"Alexa open my demo skill and make {" + TypeTaskName + "}",
			"prepare {" + TypeTaskName + "}",
			"make {" + TypeTaskName + "} with my demo skill",
			"sudo make me a {" + TypeTaskName + "}",
		},
		DoSomethingTitle: {"Do something..."},
		DoSomethingText: {
			"Preparing %s.",
			"Making %s for you.",
			"%s is being processed.",
		},
		DoSomethingSSML: {
			ssml.Speak("Ok, I'm working on %s, please be patient."),
			ssml.Speak(
				ssml.UseVoice(
					"Salli",
					"%s is being processed. Please wait a few minutes.",
				),
			),
			ssml.Speak(ssml.UseVoiceLang(
				"Kendra",
				"en-US",
				"Your %s is progressing. Please be patient.",
			)),
		},
		DoSomethingUserText:       {"Sure %s, preparing %s."},
		DoSomethingUserSSML:       {ssml.Speak("Ok %s, I'm preparing %s for you.")},
		DoSomethingTaskElicitText: {"What task do you want to process?"},
		DoSomethingTaskElicitSSML: {ssml.Speak("Which task do you want to run?")},

		// Slots
		TypeTaskValues: {"coffee", "lunch", "dinner", "sandwich"},
		DoSomethingTaskSamples: {
			"prepare {" + TypeTaskName + "}",
			"make {" + TypeTaskName + "} with my demo skill",
			"sudo make me a {" + TypeTaskName + "}",
		},

		// Slot prompts
		SlotTaskElicitTitle:     {"Unknown task", "Wrong task"},
		SlotTaskElicitText:      {"I'm sorry. Which task did you mean?"},
		SlotTaskElicitSSML:      {ssml.Speak("Please try again, which task did you mean?")},
		SlotWrongTaskElicitText: {"Sorry, I don't know any task '%s'. Please try again."},
		SlotWrongTaskElicitSSML: {ssml.Speak("I don't know '%s'. Please try again.")},

		// required for tests to work (delegated to Alexa in real use)
		AMAZONStopSamples:   {"stop", "terminate", "end"},
		AMAZONHelpSamples:   {"help", "help me", "what now"},
		AMAZONCancelSamples: {"abort", "cancel"},
	},
}
