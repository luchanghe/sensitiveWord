# sensitiveWord
一个基于Ac自动机算法的敏感词检测服务

# 使用方式
words.txt中每一行都认为是一个敏感词

使用 `go run main -addr {地址，默认本地4396} `启动服务

外部调用使用get或者post请求,接收的参数为 `text`

### 调用结果

```
GET http://127.0.0.1:4396?text=敏敏是一个很敏感的敏感词

HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 26 Apr 2024 03:27:41 GMT
Content-Length: 49

{
  "message": "**是一个***的***",
  "isMatch": true
}

POST http://127.0.0.1:4396

HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 26 Apr 2024 03:27:54 GMT
Content-Length: 49

{
  "message": "**是一个***的***",
  "isMatch": true
}
```

## 已完成
- 支持文本导入
## 未来计划
- 词库更新服务重启
- 词库增加第三方支持
- 输出敏感词的位置？或许前端做展示需要用到，比如将小说中的敏感词高亮什么的？
- 以后想到啥再说
