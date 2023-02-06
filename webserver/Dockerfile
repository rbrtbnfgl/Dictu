FROM golang

COPY ./webserver/index.html /
COPY ./webserver/styles.css /
COPY ./webserver/execute.go /
COPY ./dictu /usr/bin/
ENTRYPOINT ["go", "run", "/execute.go"]
