#!/usr/bin/env bash

BASE_URL=http://localhost:3000

curl -sSL \
  -X POST $BASE_URL/api/v1/bookmarks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Google",
    "url": "https://google.com"
  }' | jq
