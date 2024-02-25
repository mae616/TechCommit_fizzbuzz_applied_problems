# Use the official Golang image as the base image
FROM golang:latest

RUN mkdir -p /app/fizzbuzz

# Set the working directory inside the container
WORKDIR /app/fizzbuzz

COPY ./src /app/fizzbuzz

RUN go mod init github.com/mae616/TechCommit_fizzbuzz_applied_problems && go install ./cmd/fizzbuzz/fizzbuzz.go

