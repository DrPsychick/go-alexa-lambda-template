package main

import (
	"fmt"
	"strings"

	"github.com/drpsychick/go-alexa-lambda-template/loca"
	"github.com/urfave/cli/v2"
)

func runCheck(c *cli.Context) error {
	// check all `*_Samples` localization for punctuation
	// Workaround: since we cannot access "TextSnippets" and search for "*_Samples", we must use a static list :(
	samplesList := []string{
		loca.DoSomethingSamples, loca.AMAZONStopSamples, loca.AMAZONHelpSamples, loca.AMAZONCancelSamples,
	}
	var foundIn []string
	// pointRegex := regexp.MustCompile(`[a-zA-Z]\.`)
	for _, loc := range loca.Registry.GetLocales() {
		//
		for _, sample := range samplesList {
			for _, t := range loc.GetAll(sample) {
				// Sample utterances can consist of only unicode characters, spaces, periods for abbreviations,
				// underscores, possessive apostrophes, and hyphens.
				if strings.ContainsAny(t, ".,!?/") {
					foundIn = append(foundIn, sample)
					continue
				}
				// if pointRegex.MatchString(t) {
				//	foundIn = append(foundIn, sample)
				// }
			}
		}
	}
	if len(foundIn) > 0 {
		msg := "Sample utterances can consist of only unicode characters, spaces, periods for " +
			"abbreviations, underscores, possessive apostrophes, and hyphens."
		return fmt.Errorf(
			"found punctuation in sample utterances: %s (%s)", strings.Join(foundIn, ", "), msg,
		)
	}
	return nil
}
