#!/bin/bash
cp /nexus/secret/.credentials /nexus

if [ ! -f /nexus/.credentials ]; then
    echo ".credentials file in /nexus/secret is missing"
    exit 1
fi

if [ -z "$KEEP_LIMIT" ]; then
    KEEP_LIMIT=200
fi
echo "Keeping the last $KEEP_LIMIT images"
cd /nexus
/nexus/nexus-cli cleanup -k $KEEP_LIMIT

