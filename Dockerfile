FROM tensorflow/tensorflow:2.18.0 AS go-tf

ARG TARGETARCH
RUN if [ $TARGETARCH = "arm64" ]; then \
    curl -L "https://go.dev/dl/go1.23.4.linux-arm64.tar.gz" | tar -C /usr/local -xz \
    ; fi
RUN if [ $TARGETARCH = "amd64" ]; then \
    curl -L "https://go.dev/dl/go1.23.4.linux-amd64.tar.gz" | tar -C /usr/local -xz \
    ; fi

ENV GOPATH /usr/local/go
ENV GO111MODULE=on
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN chmod -R 777 "$GOPATH"
# RUN export PATH=$PATH:/usr/local/go/bin

####################################################################################################

FROM go-tf AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY service service

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/server

####################################################################################################

FROM tensorflow/tensorflow:2.18.0 AS runtime

WORKDIR /app

COPY --from=build /app/server server

EXPOSE 8080

CMD ["/app/server"]
