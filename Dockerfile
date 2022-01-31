FROM golang:1.17.6-alpine

COPY . /app

WORKDIR /app

RUN go mod download

RUN mkdir -p ./bin
RUN go build -o /bin/tcms /app/cmd/tcms

EXPOSE 8888

CMD [ "/bin/tcms" ]
