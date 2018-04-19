#!/bin/bash
cp /nexus/secret/.credentials /nexus

if [ ! -f /nexus/.credentials ]; then
    echo ".credentials file in /nexus/secret is missing"
    exit 1
fi

cd /nexus
/nexus/nexus-cli cleanup -k $KEEP_LIMIT

