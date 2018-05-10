FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/middleware
RUN go get github.com/sirupsen/logrus
RUN go build -o main . 
CMD ["/app/main"]
