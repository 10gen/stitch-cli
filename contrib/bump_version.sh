#!/bin/sh

set -e

VERSION=$1
GITHASH=$2

if [ "$GITHASH" == "" ]; then
  GITHASH=`git rev-parse HEAD`
fi


DIR=`aws s3 ls stitch-clis --recursive | grep $GITHASH | sort | tail -n 1 | awk '{print $4}' | cut -f 1-1 -d "/"`

if [ "$DIR" == "" ]; then
  echo "error: failed to find release for git hash $GITHASH"
  exit 1
fi

echo "updating 'version.json'..."
cat <<EOF > version.json
{
  "version": "$VERSION",
  "baseDirectory": "$DIR"
}
EOF

npm version --no-git-tag-version $VERSION
git add ./version.json ./package*
git commit -m "$VERSION"
git tag -m "$VERSION" -a "v$VERSION"

echo "Success!\n"

CONTRIB=$(cat <<EOF
{
  "version": "$VERSION",
  "info": {
    "linux-amd64": {
      "url": "https://s3.amazonaws.com/stitch-clis/$DIR/linux-amd64/stitch-cli"
    },
    "macos-amd64": {
      "url": "https://s3.amazonaws.com/stitch-clis/$DIR/macos-amd64/stitch-cli"
    },
    "windows-amd64": {
      "url": "https://s3.amazonaws.com/stitch-clis/$DIR/windows-amd64/stitch-cli.exe"
    }
  }
}
EOF
)

if type "pbcopy" &> /dev/null; then
  echo "copying CONTRIB body to clipboard...\n"
  echo "$CONTRIB" | pbcopy
fi
echo "$CONTRIB"

