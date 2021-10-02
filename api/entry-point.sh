#!/bin/sh -e

RAILS_ENV=production bundle exec rails db:create
RAILS_ENV=production bundle exec rails db:migrate

exec "$@"