FROM gitpod/workspace-full:latest

USER gitpod

RUN brew update \
    && brew install sqlc \
    && go install github.com/air-verse/air@latest