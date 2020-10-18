# [Choice] Node.js version: 14, 12, 10
ARG NODE_VERSION=14

FROM mcr.microsoft.com/vscode/devcontainers/typescript-node:${NODE_VERSION}
ARG PURESCRIPT_VERSION=0.13.8
ENV PURESCRIPT_VERSION=$PURESCRIPT_VERSION

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    ripgrep \
    libtinfo5

RUN curl -LO https://github.com/purescript/purescript/releases/download/v${PURESCRIPT_VERSION}/linux64.tar.gz \
    && tar -xzf linux64.tar.gz \
    && mv purescript/purs /usr/local/bin/purs \
    && rm -rf linux64.tar.gz purescript

# [Optional] Uncomment if you want to install an additional version of node using nvm
# ARG EXTRA_NODE_VERSION=10
# RUN su node -c "source /usr/local/share/nvm/nvm.sh && nvm install ${EXTRA_NODE_VERSION}"

# [Optional] Uncomment if you want to install more global node packages
RUN sudo -u node npm install -g spago
