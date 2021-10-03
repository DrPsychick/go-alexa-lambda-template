[![Build Status](https://app.travis-ci.com/DrPsychick/go-alexa-lambda-template.svg?branch=master)](https://app.travis-ci.com/DrPsychick/go-alexa-lambda-template)
[![Coverage Status](https://coveralls.io/repos/github/DrPsychick/go-alexa-lambda-template/badge.svg?branch=master)](https://coveralls.io/github/DrPsychick/go-alexa-lambda-template?branch=master)
[![Contributors](https://img.shields.io/github/contributors/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template/graphs/contributors)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template/pulls)
[![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub stars](https://img.shields.io/github/stars/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template)

# go-alexa-lambda-template
Template project using [go-alexa-lambda](https://github.com/DrPsychick/go-alexa-lambda) library.

## How to use it
### Checkout the template
```shell
git clone https://github.com/DrPsychick/go-alexa-lambda-template.git
mv go-alexa-lambda-template my-skill-name
cd my-skill-name
rm -rf .git
# create your own git repo and push it
git add remote ...
git add .
git commit -m "initial commit"
git push
```

### Replace existing Intent with your own
This templates is ready to use with an existing demo intent. Here are some example phrases:
* Alexa open my demo skill and make coffee.
* Alexa prepare lunch with my demo skill.
* Alexa start my demo skill. Yes? Sudo make me a sandwich.

To replace it with your own first intent, do the following:
#### Replace sample constants and translations
* Replace Intent and Slot name in `loca/loca.go`
  * Replace `DoSomething` with your intent name
  * Replace `Task` with your slot type
* Update your english locale accordingly `loca/en-US.go`
  * Like above replace `DoSomething` and `Task`
  * Change texts and phrases to match your intent.
* Make sure all references to the changed `loca.` constants are updated in all files.

#### Replace demo with your own logic
* Adjust `skill.go` to define your skill manifest and interaction model.
* Adjust `app.go` and add the behavioural or business logic for your intent.
* Adjust `lambda/lambda.go` to add more intents or change how to process Alexa requests.
* Of course, adjust `loca` files to your needs, adding keys and changing texts.

#### Verify skill manifest and interaction model
* run `make build; ./mydemoskill make`
* check `alexa/skill.json` and `alexa/interactionModel/custom/en-US.json `


## Before you code
Make a concept for your skill, e.g. use a mind map
* How will users launch it?
* How will users interact with your skill?
* Write down example phrases and dialogs with your skill.
* Read them out loud and check if it "feels" like a smooth conversation.
* Make sure to repeat the above steps for every language (locale). Especially your invocation will likely be specific to the locale.

## Testing your skill
* Create sample lambda requests to ensure your lambda function answers correctly.
* Create a few common dialogs you want to have tested after deploy.
  * Important: `ask cli dialog` does not work with prompts delegated to Alexa! It will just timeout.
* Manually test dialogs with prompt delegation to Alexa with the Alexa SDK website.