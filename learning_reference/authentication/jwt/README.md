openssl genrsa -out demo.rsa 1024
openssl rsa -in demo.rsa -pubout > demos.rsa.pub
