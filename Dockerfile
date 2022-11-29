FROM golang:1.19-alpine

WORKDIR /go/delivery

COPY . .

RUN go build PetrosianBot.go

CMD ["./PetrosianBot"]