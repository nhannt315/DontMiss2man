#!/bin/sh -e

bundle exec rails db:migrate

exec "$@"