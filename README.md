[![Workflow Status](https://img.shields.io/github/actions/workflow/status/drpsychick/go-alexa-lambda-template/ci.yaml)](https://github.com/DrPsychick/go-alexa-lambda-template/actions)
[![Coverage Status](https://coveralls.io/repos/github/DrPsychick/go-alexa-lambda-template/badge.svg?branch=main)](https://coveralls.io/github/DrPsychick/go-alexa-lambda-template?branch=main)
[![Contributors](https://img.shields.io/github/contributors/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template/graphs/contributors)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template/pulls)
[![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub stars](https://img.shields.io/github/stars/drpsychick/go-alexa-lambda-template.svg)](https://github.com/drpsychick/go-alexa-lambda-template)

# go-alexa-lambda-template
Template project using [go-alexa-lambda](https://github.com/DrPsychick/go-alexa-lambda) library.

## How to use it
### Checkout the template
```shell
# this must be in your GOPATH, e.g. in the path to your git repo (see below)
git clone https://github.com/DrPsychick/go-alexa-lambda-template.git
mv go-alexa-lambda-template my-skill-name
cd my-skill-name
rm -rf .git
```

### Test, build and run it
```shell
# Makefile requires mmake
go install github.com/tj/mmake/cmd/mmake@v1.4.0
alias make='mmake'
make test
make build

# generate skill and interactionModel(s)
./mydemoskill make --skill --models
ls -la ./alexa/skill.json ./alexa/*/*/en-US.json

# test how lambda responds to a launch request
./test/test-lambda.sh launch
```

### Push it to your own git
```shell
# create your own git repo and push it
GITREPO=github.com/yourname/yourproject
git init .
git remote add origin https://$GITREPO.git
# replace module paths of the template with your own
find . -name \*.go -exec sed -i "" -e "s#github.com/drpsychick/go-alexa-lambda-template#$GITREPO#" {} ';'
git add .
git commit -m "initial commit"
git push
```

# Build your own skill
## Replace existing Intent with your own
This templates is ready to use with an existing demo intent. Here are some example phrases:
* Alexa open my demo skill and make coffee.
* Alexa prepare lunch with my demo skill.
* Alexa start my demo skill. Yes? Sudo make me a sandwich.

To replace it with your own first intent, do the following:
## Replace sample constants and translations
* Replace **Intent** and **Slot** name in `loca/loca.go`
  * Replace `DoSomething` with your intent name
  * Replace `Task` with your slot type
* Update your english locale accordingly `loca/en-US.go`
  * Like above replace `DoSomething` and `Task`
  * Change texts and phrases to match your intent.
* Make sure all references to the changed `loca.` constants are updated in all files.

## Replace demo with your own logic
* Adjust `skill.go` to define your skill manifest and interaction model.
* Adjust `app.go` and add the behavioural or business logic for your intents.
* Adjust `lambda/lambda.go` to add more intents or change how to process Alexa requests.
* Of course, adjust `loca` files to your needs, adding keys and changing texts.

## Verify skill manifest and interaction model
* run `make build; ./mydemoskill make --skill --models`
* check `alexa/skill.json` and `alexa/interactionModel/custom/en-US.json`


# Before you code
## Interaction concept
Make a concept for your skill on how people will interact with it, e.g. use a mind map.
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
