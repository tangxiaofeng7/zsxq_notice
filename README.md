# zsxq_notice

知识星球提醒

## 操作指南：

首先编辑 config/config.yaml

```
database:
  default:
    type:  "mysql"
    link:  "root:woaini520@tcp(mysql)/zsxq"

wechat:
    key:    "dc7d86c2-ce47-4aa7-9fa1-111111111111"

zsxq_access_token:
    token:    "DFD51610-752F-369F-1BC7-1287690C57B0_1111111111111111"

zsxq_group:
  - id: "11111111111"
  - id: "2222222222222"
```

在 linux 环境:

```
cd zsxq_notice
docker-compose up -d
```
## 卸载：
```
docker stop $(docker ps -a | grep "zsxq_notice" | awk '{print $1}') 
docker rm $(docker ps -a | grep "zsxq_notice" | awk '{print $1}') 
docker rmi $(docker images | grep "zsxq_notice" | awk '{print $3}') 

```