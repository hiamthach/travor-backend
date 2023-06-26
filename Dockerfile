# Build stage
FROM golang:1.20-alpine3.18
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY app.env ./
#RUN go mod download
RUN go get -d -v ./...
#RUN go mod vendor

COPY . .

RUN go build -o /travor-backend

EXPOSE 8088

CMD [ "/travor-backend" ]