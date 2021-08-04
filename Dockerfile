FROM golang:1.16
WORKDIR . 
COPY . .
ENV GOPATH=""
RUN make
CMD ["./build/gogo-bot"]
