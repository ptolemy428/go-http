# go-http

#to build docker image
docker build -t local:go-http .

#to run docker image
docker run -it --rm -p8080:8080 local:go-http
