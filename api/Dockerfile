# syntax = docker/dockerfile:experimental

FROM ruby:2.6.5-slim AS base

ENV APP_DIR /app
WORKDIR $APP_DIR

ARG GROUP_ID=1100
ARG USER_ID=1100
RUN groupadd --system --gid $GROUP_ID app && \
      useradd --system --gid $GROUP_ID -m --uid $USER_ID app && \
      mkdir -p $APP_DIR && \
      chown app:app $APP_DIR

RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

RUN --mount=type=cache,target=/var/cache/apt --mount=type=cache,target=/var/lib/apt \
      apt update \
      && apt upgrade -y \
      && apt install -y --no-install-recommends \
      git \
      openssh-client \
      make \
      gcc \
      g++ \
      default-libmysqlclient-dev \
      shared-mime-info

USER $USER_ID:$GROUP_ID

COPY --chown=app:app Gemfile .
COPY --chown=app:app Gemfile.lock .
COPY --chown=app:app entry-point.sh .

ENV BUNDLE_PATH $APP_DIR/.bundle

RUN mkdir -p $BUNDLE_PATH

# gemのインストールと実行時に不要なものを削除して軽量化
RUN --mount=type=cache,target=/app/.bundle/cache,uid=1100 \
      echo 'gem: --no-document' > ~/.gemrc \
      && bundle install --jobs 6 --without development test \
      && rm -rf /app/.bundle/cache/*.gem \
      && rm -rf /app/.bundle/cache/bundler \
      && rm -rf /app/.bundle/bundler/gems/**/.git

# RAILS_ENV=production 用軽量イメージ
FROM base AS production

COPY --chown=app:app . .

ENTRYPOINT ["./entry-point.sh"]
