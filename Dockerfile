FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /domain-driven-design-layout

CMD [ "/domain-driven-design-layout" ]
