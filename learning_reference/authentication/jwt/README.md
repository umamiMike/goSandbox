This is taken mostly from [this presentation by Yamil Asusta](https://www.youtube.com/watch?v=eVlxuST7dCA)
to generate rsa keys for demo you can run the commands below or run the bash script
` openssl genrsa -out demo.rsa 1024 `
` openssl rsa -in demo.rsa -pubout > demos.rsa.pub t

to test through 
`go run jwt.go`
 to run the the file which will start a web server

from the terminal run
` curl localhost:3000/login` will return the public key for demonstration purposes   
`curl localhost:3000/api` will return "Not Authorized" with a 401   

however if you copy and paste the public key from the login route
and `curl -H "Authorization: Bearer {paste the key in here} " localhost:3000/api` will return "Authorized"
