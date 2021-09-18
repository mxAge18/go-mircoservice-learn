- 通过openssl自签名证书
```bash

~/go/src/go-mircoservice-learn/gRPC-learn/v1/certs   master  openssl
OpenSSL> genrsa -out ca.key 2048
Generating RSA private key, 2048 bit long modulus
......................................................................................................................+++
....+++
e is 65537 (0x10001)
OpenSSL>  req -new -x509 -days 365 -key ca.key -out ca.pem
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:CN
State or Province Name (full name) []:beijing
Locality Name (eg, city) []:beijing
Organization Name (eg, company) []:invokerx
Organizational Unit Name (eg, section) []:invokerx
Common Name (eg, fully qualified host name) []:localhost
Email Address []:maxu0410@163.com
OpenSSL> genrsa -out server.key 2048
Generating RSA private key, 2048 bit long modulus
......+++
...................................................................+++
e is 65537 (0x10001)
OpenSSL> req -new -key server.key -out server.csr    
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:CN
State or Province Name (full name) []:beijing
Locality Name (eg, city) []:beijing 
Organization Name (eg, company) []:invokerx
Organizational Unit Name (eg, section) []:invokerx
Common Name (eg, fully qualified host name) []:localhost
Email Address []:maxu0410@163.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:invokerxcom
OpenSSL> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 365 -in server.csr -out server.pem
Signature ok
subject=/C=CN/ST=beijing/L=beijing/O=invokerx/OU=invokerx/CN=localhost/emailAddress=maxu0410@163.com
Getting CA Private Key
OpenSSL> ecparam -genkey -name secp384r1 -out client.key
OpenSSL> req -new -key client.key -out client.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) []:CN
State or Province Name (full name) []:beijing
Locality Name (eg, city) []:beijing
Organization Name (eg, company) []:invokerx
Organizational Unit Name (eg, section) []:invokerx
Common Name (eg, fully qualified host name) []:localhost
Email Address []:maxu0410@163.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:invokerxcom
OpenSSL> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 365 -in client.csr -out client.pem
Signature ok
subject=/C=CN/ST=beijing/L=beijing/O=invokerx/OU=invokerx/CN=localhost/emailAddress=maxu0410@163.com
Getting CA Private Key
OpenSSL> 

```