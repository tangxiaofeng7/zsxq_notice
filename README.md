# zsxq_notice
知识星球提醒

## 操作指南：
首先编辑config/config.yaml

```
database:
  default:
    type:  "mysql"
    link:  "root:woaini520@tcp(mysql)/zsxq"

#企业微信机器人key
wechat:
    key:    "xxxxxx"

#知识星球的zsxq_access_token
zsxq_access_token:
    token:    "xxxx"
#需要通知的知识星球group
zsxq_group:
  - id: "xxxx"
  - id: "xxx"
#扫描更新时间，默认3600秒
time:
  value: 3600

```

在linux环境:
```
cd zsxq_notice
docker-compose up -d
```

