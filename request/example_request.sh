#!/usr/bin/env bash

BASE_URL="http://localhost:3000"

curl -X POST $BASE_URL/api/v1/bookmarks \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "title": "Example"
  }'
