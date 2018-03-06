FROM debian:jessie-slim
COPY dist/ /website
WORKDIR "/website/"
ENTRYPOINT [ "/website/website" ]