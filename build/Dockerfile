FROM neohsudroid/golang-builder:1.17.1 as builder
ARG CommitHash
ARG BuildTime
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X 'main.CommitHash=${CommitHash}' -X 'main.BuildTime=${BuildTime}'" -o app cmd/app/app.go
RUN upx app

FROM neohsudroid/scratch-ca-certs:latest
EXPOSE 9000
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]