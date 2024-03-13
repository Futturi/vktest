FROM golang:1.22.0
RUN go version
ENV GOPATH=/
COPY ./ ./

RUN go mod tidy
RUN go mod download
RUN go build -o kino ./cmd/main.go
CMD [ "./kino" ]