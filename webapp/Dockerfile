FROM golang

WORKDIR /app
# Current workdir is /app
COPY main.go go.mod ./ 
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-webapp
CMD [ "/go-webapp" ]