FROM golang:1.5.1
MAINTAINER Xavier Bruhiere

COPY ./go.sh /usr/local/bin/wgo

## Project Management
## -----------------------------------------------------------------------------

# Vendoring - https://github.com/Masterminds/glide
RUN apt-get update && apt-get install -y unzip && \
  curl -LkOs "https://github.com/Masterminds/glide/releases/download/0.5.1/glide-linux-amd64.zip" && \
  unzip glide-linux-amd64.zip && \
  mv linux-amd64/glide $GOPATH/bin/ && \
  rm -r *linux-amd64*

RUN go get github.com/alecthomas/gometalinter && \
  gometalinter --install --update

CMD ["wgo"]
