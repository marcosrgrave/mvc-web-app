FROM golang:1.20 AS build
WORKDIR /github.com/marcosrgrave/mvc-crud-products
COPY . .
COPY main.go .
ENV CGO_ENABLED=0
RUN go get -d -v ./...
RUN go build -a -installsuffix cgo -o mvc-crud-products .

FROM scratch AS runtime
COPY --from=build /github.com/marcosrgrave/mvc-crud-products ./
EXPOSE 8080/tcp
ENTRYPOINT ["./mvc-crud-products"]
