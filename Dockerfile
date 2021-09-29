FROM golang:1.16-alpine
LABEL maintainer="SpectreH"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.description="Ascii-art-web in docker"

RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/
COPY . /usr/src/app/

RUN go mod download
RUN go build -o /ascii-web
EXPOSE 8000

CMD [ "/ascii-web" ]