FROM golang:1.13
WORKDIR /var/app
RUN mkdir -p /go/src/github.com/lbernardo
RUN cd  /go/src/github.com/lbernardo ; git clone https://github.com/lbernardo/aws-local
COPY apigateway.sh /bin/apigateway
RUN ln -s /go/src/github.com/lbernardo/aws-local/linux/awslocal /bin/awslocal
ENTRYPOINT ["sh","apigateway"]