FROM golang:1.16.4 as build

WORKDIR /mc-bedrock-runner

COPY go.mod .
COPY go.sum  .
RUN go mod download

COPY . .
RUN make build

FROM ubuntu

RUN apt-get update && \
  apt install curl

WORKDIR /data
VOLUME [ "/data" ]

COPY --from=build /mc-bedrock-runner/build/mc-bedrock-runner .

RUN curl -fsSOL https://minecraft.azureedge.net/bin-linux/bedrock-server-1.16.221.01.zip && unzip bedrock-server-1.16.221.01.zip
ENV LD_LIBRARY_PATH=.

ENTRYPOINT ["./mc-bedrock-runner"]
CMD ["./bedrock_server"]
