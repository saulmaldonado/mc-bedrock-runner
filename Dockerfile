FROM golang:1.16.4 as build

WORKDIR /mc-bedrock-runner

COPY go.mod .
COPY go.sum  .
RUN go mod download

COPY . .
RUN make build

FROM ubuntu

RUN apt-get update && \
  apt install -y \
  curl \
  unzip

WORKDIR /data
VOLUME [ "/data" ]

RUN curl -fsSOL https://minecraft.azureedge.net/bin-linux/bedrock-server-1.16.221.01.zip && unzip bedrock-server-1.16.221.01.zip

COPY --from=build /mc-bedrock-runner/build/mc-bedrock-runner .

RUN chmod +x bedrock_server

ENV LD_LIBRARY_PATH=.

ENTRYPOINT ["./mc-bedrock-runner"]
CMD ["./bedrock_server"]
