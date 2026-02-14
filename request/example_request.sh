#!/usr/bin/env bash

BASE_URL=http://localhost:3000

# curl -sSL \
#   -X POST $BASE_URL/api/v1/bookmarks \
#   -H "Content-Type: application/json" \
#   -d '{
#     "title": "Google",
#     "url": "https://google.com"
#   }' | jq

curl -sSL \
  -X GET $BASE_URL/api/v1/bookmarks | jq

BOOKMARK_ID=fee950eb-8fad-4b2c-adca-51087ff003d4
curl -sSL \
  -X GET "$BASE_URL/api/v1/bookmarks/$BOOKMARK_ID" | jq

curl -sSL \
  -X DELETE "$BASE_URL/api/v1/bookmarks/$BOOKMARK_ID" | jq
