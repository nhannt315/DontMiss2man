FROM golang:1.17.3-alpine as builder

WORKDIR /go

# install requirements
RUN apk add make gcc git g++

COPY go.mod go.sum /

# install goose
RUN go install github.com/pressly/goose/cmd/goose@v2.7.0

FROM alpine as migrate_base
# set Tokyo TZ
ENV TZ=Asia/Tokyo
ENV USER_ID=1000
ENV GROUP_ID=1000

ENV APP_DIR /app

# migration.shの仕様に従う。(bash)
RUN apk add bash && mkdir /db/

# create and set app user/group to APP_DIR
RUN addgroup --gid $GROUP_ID app \
  && adduser --uid $USER_ID \
  --disabled-password \
  --ingroup app \
  --no-create-home \
  app

USER $USER_ID:$GROUP_ID
WORKDIR /

# copy goose from golang image(builder)
COPY --chown=app:app --from=builder /go/bin/goose /bin/goose

COPY --chown=app:app ./deployments/migration/migration.sh /db/migration.sh
COPY --chown=app:app ./deployments/migration/entrypoint.sh /entrypoint.sh

COPY --chown=app:app ./deployments/migration/migrations /db/migrations

ENTRYPOINT [ "bash", "/entrypoint.sh" ]
CMD ["/db/migration.sh"]
