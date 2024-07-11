FROM gitpod/workspace-full

USER gitpod

RUN brew update \
    && brew install sqlc \
    && go install github.com/air-verse/air@latest