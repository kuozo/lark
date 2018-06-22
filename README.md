# Lark

```
Status: alpha
```

> Get Cloud Project Release Information For DingDing Robot (Toy Project)

## support tracker project

* K8S
* istio

## How to work

### Build Linux

```
./scripts/build.sh
```

### Run your Local cmd

```
bin/lark --token <your_dingding_token>
```

### Docker 

* Build Docker

```
docker build -t lark:latest .
```

* Docker from quay

```
docker pull quay.io/klnchu/lark:latest
```

* run 

```
docker run -it -e DING_ROBOT_TOKEN=xxxxx lark:latest /usr/local/bin/lark
```
