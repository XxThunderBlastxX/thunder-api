FROM gitpod/workspace-full:latest

USER gitpod

RUN brew update \
    && brew install sqlc \
    && brew install pkl \
    && go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest \
    && go install github.com/air-verse/air@latest \