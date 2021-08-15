FROM node:14.17.0-alpine AS base

ARG GROUP_ID=1100
ARG USER_ID=1100

ENV NODE_ENV=production
ENV APP_DIR /app

RUN mkdir -p ${APP_DIR}
WORKDIR ${APP_DIR}

# set timezone
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN apk add --no-cache python2 bash make g++ cmake

RUN npm install pm2 -g

# create non-root user
RUN addgroup --system --gid ${GROUP_ID} app && \
    adduser --system --uid ${USER_ID} app && \
    chown app:app ${APP_DIR}
USER ${USER_ID}:${GROUP_ID}


COPY --chown=app:app ./entry-point.sh .
ENTRYPOINT ["./entry-point.sh"]

################################################################################

FROM base AS builder

# copy entire files
COPY --chown=app:app . .

################################################################################

FROM builder AS built

# specify 'local', 'stg' or 'prod'
ARG BUILD_TARGET
# copy suitable .env file
COPY --chown=app:app env.${BUILD_TARGET} .env

# build reactjs
RUN yarn install --production && yarn build

################################################################################

FROM base AS app

# copy files from build stage
COPY --from=built ${APP_DIR}/package.json .
COPY --from=built ${APP_DIR}/package-lock.json .
COPY --from=built ${APP_DIR}/.env .
COPY --from=built ${APP_DIR}/build ./build
COPY --from=built ${APP_DIR}/server.js server.js

# install npm packages for production
RUN yarn install --production

EXPOSE 3000
CMD ["node", "server.js"]