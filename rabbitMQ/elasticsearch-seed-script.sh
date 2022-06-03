#!/bin/sh

echo ELASTICSEARCH_URL: "$ELASTICSEARCH_URL"
echo INDEX_NAME: "$INDEX_NAME"

until $(curl -XGET "$ELASTICSEARCH_URL/_cluster/health?wait_for_status=green" > /dev/null); do
    printf 'WAITING FOR THE ELASTICSEARCH ENDPOINT BE AVAILABLE, trying again in 5 seconds \n'
    sleep 5
done

# The bulk operation to insert multiple documents into the index
curl -XPOST "$ELASTICSEARCH_URL/$INDEX_NAME/_bulk" -H 'Content-Type: application/x-ndjson' --data-binary @index-bulk-payload.json
