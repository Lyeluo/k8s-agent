FROM golang:alpine AS build
RUN mkdir /app
COPY . /app
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /app
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 \
     go mod tidy && go build -o k8s-agent
      
### 
FROM alpine:3.13 as final
LABEL author="micro-liujy"
COPY --from=build /app/k8s-agent .
EXPOSE 19000
CMD ["/k8s-agent"]
