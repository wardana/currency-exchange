FROM golang:latest
# RUN go get -u github.com/golang/dep

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   

ENV APP_PATH=/go/src/github.com/wardana/currency-exchange

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock

COPY config.yaml config.yaml

# RUN go get github.com/jinzhu/gormd

RUN dep init

ADD . $APP_PATH

CMD go run main.go