FROM golang
WORKDIR /app/src
COPY . /app/src/
CMD [ "tail", "-f" ]