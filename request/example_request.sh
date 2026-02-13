#!/usr/bin/env bash

BASE_URL=http://localhost:3000
BASE_URL=https://halucigenia-backend-995b6e4c7c37.herokuapp.com

curl -sSL \
  -X POST $BASE_URL/api/v1/bookmarks \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com",
    "title": "Example"
  }' | jq
