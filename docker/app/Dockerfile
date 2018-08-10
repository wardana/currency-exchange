FROM golang:latest

RUN  sleep 10 \
  && curl https://glide.sh/get | sh \
  && mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   

ENV APP_PATH=/go/src/github.com/wardana/currency-exchange

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

## move config and dependencies 
COPY glide.yaml glide.yaml
COPY glide.lock glide.lock
COPY config.yaml config.yaml

RUN glide install

ADD . $APP_PATH

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.3.0/wait /wait
RUN chmod +x /wait

## Launch the wait tool and then your application
CMD /wait && go run main.go