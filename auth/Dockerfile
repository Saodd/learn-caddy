FROM golang:1.15.3 as build
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /code
ADD go.mod /code/go.mod
ADD go.sum /code/go.sum
RUN go mod download

ADD . /code
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o /app /code/auth/main.go


FROM alpine
COPY --from=build /app /app
CMD /app
