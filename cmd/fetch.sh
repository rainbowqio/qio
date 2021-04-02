#!/usr/bin/env zsh

OWNER='blamelesshq'
REPO='blameless-deploy'
PATH='.circleci/config.yml'
GH_TOKEN=''
GH_API="https://api.github.com"
GH_REPO="$GH_API/repos/$OWNER/$REPO/contents/$PATH"
AUTH="Authorization: token $GH_TOKEN"
CONTENT="application/vnd.github.v3.raw"
CURL_ARGS="-LJO#"
CURL="/usr/bin/curl"
YQ=$(which yq)

# Validate Token
$CURL -o /dev/null -sH "$AUTH" $GH_REPO || { echo "Error: Invalid repo, token or network issue!";  exit 1; }

# Download Raw File
$CURL $CURL_ARGS -H "$AUTH" -H "Accept: $CONTENT" "$GH_REPO"
echo "$0 done." >&2

# Process YAML into TOML
#
#$YQ e '.jobs.[].environment.INSTANCE_VALUES_DIR' customertry.yaml | awk -F'_' 'BEGIN {print "[blameless.prod.instance]"} /prod/{print $3" = \""$1"_"$2"\""}' > prodinstance.toml

