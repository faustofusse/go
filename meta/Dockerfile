FROM golang:1.23 as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -v -tags lambda.norpc -o ./bin/bootstrap ./main.go

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /usr/src/app/bin/bootstrap ./bootstrap

ENTRYPOINT [ "./bootstrap" ]
