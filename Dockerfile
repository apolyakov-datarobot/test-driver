FROM golang:1.17
WORKDIR /app
COPY words /usr/share/dict/words
COPY /pkg/ .
COPY go.mod .
COPY go.sum .
RUN go get github.com/tjarratt/babble
RUN go install github.com/jstemmer/go-junit-report/v2@latest
COPY main-combined.sh /bin
COPY run-tests-junit-report.sh /bin
COPY upload-test-artifacts.sh /bin
RUN chmod +x /bin/main-combined.sh
RUN chmod +x /bin/run-tests-junit-report.sh
RUN chmod +x /bin/upload-test-artifacts.sh
RUN echo "Acquire::Check-Valid-Until \"false\";\nAcquire::Check-Date \"false\";" | cat > /etc/apt/apt.conf.d/10no--check-valid-until
RUN apt-get update
RUN apt-get install unzip
RUN	curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN	unzip awscliv2.zip && ./aws/install
ENV TEST_NAME ".*"
