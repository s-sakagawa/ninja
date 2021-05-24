# Nginx reverse proxy

Implementation of reverse proxy using Nginx

> sushi: reverse proxy server (localhost:80) \
> &emsp; --> tuna: backend server (localhost:7001) \
> &emsp; --> salmon: backend server (localhost:7002)

## Set up
```bash
$ cd reverse-proxy
$ docker-compose build
$ docker-compose up
```

## Server Request
```bash
# request tuna
$ curl --verbose localhost:80/tuna

# request salmon
$ curl --verbose localhost:80/salmon
```
