#!/bin/sh

# build for lambda, then send json requests to the lambda function in docker

# check for required variables
check_env_vars () {
  missing=0
  for name; do
    if [ -z "$(eval echo \$$name)" ]; then
      echo "ENV '$name' must be set!"
      missing=1
    fi
  done
  return $missing
}

if ! check_env_vars; then
    exit 1
fi

# determine arch
docker_args="--platform linux/amd64"
if [ "$(uname -s)" != "Linux" ]; then
    docker_args="--platform linux/amd64"
fi
APP_NAME=${APP_NAME:-"mydemoskill"}

request=$1
DIR="$( cd "$( dirname "$0" )" >/dev/null 2>&1 && pwd )"

# build for lambda linux
(cd $DIR/..; export GOOS=linux; export GOARCH=amd64; go build -o ./test/app ./cmd/$APP_NAME) || exit 1

# TODO: refactor/rethink, how can this be done more elegantly (intents and locales are already defined elsewhere)
# or is this needed at all? it helps identify missing localization...
(cd $DIR;
# prepare ENV
echo "
DOCKER_LAMBDA_USE_STDIN=1
STATS_DSN=l2met://console
" > ./docker.env

intentlist="launch stopintent cancelintent helpintent DoSomething DoSomethingWrongName"
for t in $intentlist; do
    if [ -n "$request" -a "$request" != "$t" ]; then
        continue
    fi
    cat lambda_${t}.json |grep -A20 '"request"'
    # loop over locales
    for l in en-US; do
        echo "----------------------- $t ($l) ------------------------------"
        result=$(set -x; sed -e "s/LOCALE/${l}/" lambda_${t}.json | docker run $docker_args --rm -i -v "$PWD":/var/task --env-file ./docker.env lambci/lambda:go1.x app)
        err=$(echo "$result" | tr ',' '\n' | grep -i '\("content"\|"title"\):.*error.*')
        if [ -n "$err" ]; then
            failed="${failed}$l $t : $err\n"
        fi
        echo "$result" |jq .
    done
done

if [ -n "$failed" ]; then
    echo "Error(s) occurred:"
    echo "$failed"
    exit 1
fi
)