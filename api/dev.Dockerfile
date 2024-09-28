FROM golang:1.23.1
# syntax=docker/dockerfile:1.2

RUN apt update && apt upgrade -y && \
    apt-get install -y docker.io 
## add gcloud
RUN curl https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz > /tmp/google-cloud-sdk.tar.gz

# Installing the package
RUN mkdir -p /usr/local/gcloud \
  && tar -C /usr/local/gcloud -xvf /tmp/google-cloud-sdk.tar.gz \
  && /usr/local/gcloud/google-cloud-sdk/install.sh

# Adding the package path to local
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

COPY . .

CMD air