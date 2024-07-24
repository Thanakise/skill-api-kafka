FROM alpine as builder


RUN apk update && apk add go

COPY . .

RUN go build -o app main.go

FROM scratch
COPY --from=builder /app .

ENTRYPOINT [ "./app" ]