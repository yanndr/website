FROM debian:jessie-slim
COPY templates/ /website/templates
COPY public/ /website/public
COPY website /website
WORKDIR "/website/"
ENTRYPOINT [ "/website/website" ]