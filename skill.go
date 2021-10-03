package mydemoskill

import (
	"github.com/drpsychick/go-alexa-lambda-template/loca"
	"github.com/drpsychick/go-alexa-lambda/skill"
)

// NewSkill returns a configured SkillBuilder.
func NewSkill() *skill.SkillBuilder {
	return skill.NewSkillBuilder().
		WithLocaleRegistry(loca.Registry).
		WithCategory(skill.CategoryGames).
		WithPrivacyFlag(skill.FlagIsExportCompliant, true)
}

// CreateSkillModels generates and returns a list of Models.
func CreateSkillModels(s *skill.SkillBuilder) (map[string]*skill.Model, error) {
	m := s.Model().
		WithDelegationStrategy(skill.DelegationSkillResponse)

	// Prompts are only needed for the interactionModel.
	// DoSomething elicitation prompt for slot Task.
	m.WithElicitationSlotPrompt(loca.DoSomething, loca.TypeTaskName)

	// Variations must be set here and not in lambda functions.
	m.ElicitationPrompt(loca.DoSomething, loca.TypeTaskName).
		WithVariation("PlainText").
		WithVariation("SSML")

	return s.BuildModels()
}
