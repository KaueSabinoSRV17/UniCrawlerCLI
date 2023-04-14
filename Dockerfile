FROM golang

WORKDIR /app/src

COPY . /app/src/

ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra-cli@latest

CMD [ "tail", "-f" ]
