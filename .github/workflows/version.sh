#!/usr/bin/env bash


set -e

#get highest tag number
VERSION=`git describe --abbrev=0 --tags`

#replace . with space so can split into an array
VERSION_BITS=(${VERSION//./ })

#get number parts and increase last one by 1
VNUM1=${VERSION_BITS[0]}
VNUM2=${VERSION_BITS[1]}
VNUM3=${VERSION_BITS[2]}
VNUM4=${VERSION_BITS[3]}
VNUM4=$((VNUM4+1))

#create new tag
NEW_TAG="$VNUM1.$VNUM2.$VNUM3.$VNUM4"

echo "Updating $VERSION to $NEW_TAG"

#get current hash and see if it already has a tag
GIT_COMMIT=`git rev-parse HEAD`
echo "I'm about to describe..."
NEEDS_TAG=`git describe --contains $GIT_COMMIT 2>/dev/null`

#only tag if no tag already
#to publish, need to be logged in to npm, and with clean working directory: `npm login; git stash`
if [ -z "$NEEDS_TAG" ]; then
  echo "I'm about to git config..."
  git config --local user.email "action@github.com"
  git config --local user.name "GitHub Action"
  git tag -a $NEW_TAG -m "Applying tag bump"
  git push origin $NEW_TAG
else
  echo "Already a tag on this commit"
fi
