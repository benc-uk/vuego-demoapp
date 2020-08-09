#!/bin/bash

VER=$1

if [[ -z "$VER" ]]; then
  echo "Error! Supply version tag!"
  exit 1
fi

read -r -d '' NOTES << EOM
\`\`\`
docker pull ghcr.io/benc-uk/vuego-demoapp:$VER
\`\`\`

\`\`\`
docker run --rm -it -p 4000:4000 ghcr.io/benc-uk/vuego-demoapp:$VER
\`\`\`
EOM

gh release create $VER --title "Release v$VER" -n "$NOTES"
