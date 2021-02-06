```shell
docker pull caddy:2.3.0-alpine
```

```shell
docker build -t learn-caddy:auth -f auth/Dockerfile .
docker build -t learn-caddy:business -f business/Dockerfile .
docker build -t learn-caddy:caddy -f caddy.dockerfile .
```
