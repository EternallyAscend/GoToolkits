# GoToolkits.
[GoToolkits Description Document](./document/README.md)

## Preparation.

Environment for Ubuntu 20.04(LTS) with Golang1.17.3.

```Shell
# Modify to Tsinghua list first.
apt update
apt install vim curl wget git -y
apt install net-tools procps -y

# Add Nginx list.
echo "deb http://nginx.org/packages/mainline/ubuntu focal nginx" > /etc/apt/sources.list.d/nginx.list

curl -o /tmp/nginx_signing.key https://nginx.org/keys/nginx_signing.key
mv /tmp/nginx_signing.key /etc/apt/trusted.gpg.d/nginx_signing.asc

# Check Nginx and Redis version.
apt-cache madison nginx redis

# Install Nginx and Redis.
apt install nginx redis -y

# Install Golang1.17.2

wget https://studygolang.com/dl/golang/go1.17.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

# Start Redis with default config.
redis-server /etc/redis/redis.conf

# export PATH=$PATH:/usr/local/go/bin => ~/.profile
source ~/.profile

# Set Proxy for go install (or go get).
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
# go env -w GOPROXY=https://mirrors.aliyun.com/goproxy

# Install Go-Redis for go environment.
go install github.com/go-redis/redis/v8@latest

# Install Go Iris for go environment.
go install github.com/kataras/iris/v12@latest
```

## Include to project.
Create program like [`entry`](./entry) with [go.mod](./go.mod) file.

## Repo.
- `GitHub` [https://github.com/EternallyAscend/GoToolkits](https://github.com/EternallyAscend/GoToolkits)
- `Gitee` [https://gitee.com/EternallyAscend/GoToolkits](https://gitee.com/EternallyAscend/GoToolkits)

## Docker
```Shell
docker run -itd -p 4200:4200 -p 9000-9010:9000-9010 -v /Users/mac/Documents/Projects:/projects --name fw debian:latest

docker run -itd --name jupyter -p 8888:8888 tensorflow/tensorflow:latest-jupyter
docker run -itd --name p36tf115 -p 8888:8888 python:3.6
docker exec -d p36tf115 jupyter notebook --ip=0.0.0.0 --allow-root
docker exec -it p36tf115 jupyter notebook list

# docker run -itd -p 6379:6379 -v /root/redis/config:/etc/redis -v /root/redis/data:/data -v /root/redis/log:/var/log/redis -v /root/redis/lib:/var/lib/redis --name redis redis redis-server /etc/redis/redis.conf --appendonly yes
docker run -itd -p 6379:6379 -v /root/redis:/data --name redis redis redis-server /data/config/redis.conf

# Nginx
docker run -itd --name ContentSite -p 80:80 -p 443:443 -v /usr/share/website/ContentSite/nginx:/etc/nginx -v /usr/share/website/ContentSite/dist:/usr/share/nginx/html -v /usr/share/website/ContentSite/nginx/log:/var/log/nginx nginx:latest /bin/bash
docker run -d --name WebRunner0 -v /usr/share/website/gitlab:/gitlab  gitlab/gitlab-runner
```

```Shell
git remote set-url --add origin `address`
```
