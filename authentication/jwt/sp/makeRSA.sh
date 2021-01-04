openssl genrsa -f4 -out demo.rsa 4096 
openssl rsa -in demo.rsa -outform PEM -pubout -out demo.rsa.pub
