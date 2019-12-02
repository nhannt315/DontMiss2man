#!/usr/bin/env bash
kill -9 `cat /home/ubuntu/jdict-web/backend/tmp/pids/puma.pid`
pm2 kill
cd /home/ubuntu/DontMiss2man
git checkout -- .
git pull origin develop --rebase
cd /home/ubuntu/DontMiss2man/frontend/
yarn
yarn build
pm2 start server.js
cd /home/ubuntu/DontMiss2man/backend/
bundle install
rails db:migrate
bundle exec puma -e development -p 4000 --pidfile tmp/pids/puma.pid -d
