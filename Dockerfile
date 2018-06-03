FROM golang:alpine

ADD . /go/src/github.com/justinbarrick/go-terraform-plan
WORKDIR /go/src/github.com/justinbarrick/go-terraform-plan

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/go-terraform-plan

FROM scratch

COPY --from=0 /go/bin/go-terraform-plan /bin/terraform-plan

ENTRYPOINT ["/bin/terraform-plan"]
