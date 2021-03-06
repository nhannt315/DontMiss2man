#!/usr/bin/env bash
#source ~/.bashrc
#sudo timedatectl set-timezone Asia/Tokyo
kill -9 `cat /home/ubuntu/DontMiss2man/backend/tmp/pids/puma.pid`
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
bundle exec puma -e production -p 4000 --pidfile tmp/pids/puma.pid -d
bundle exec whenever --update-crontab --set environment='production'
