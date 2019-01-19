#!/bin/bash -e

curl -H "Accept: application/vnd.github.mercy-preview+json" \
  https://api.github.com/search/repositories?q=user:pankona+created:2019-01-01..2019-02-01
