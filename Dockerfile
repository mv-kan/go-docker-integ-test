FROM golang:1.21rc3-alpine3.17

WORKDIR /app
COPY . ./

RUN go build -o sendthing-app .

# EXPOSE 10123
CMD [ "./sendthing-app" ]