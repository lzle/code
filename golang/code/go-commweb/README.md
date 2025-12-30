# 呼叫中心通用接口

### 呼入

申请 `/callin/apply`

释放 `/callin/release`

### 智能外呼

开启 `/robot/start`

关闭 `/robot/stop`

监听 `/robot/listen`

拦截 `/robot/intercept`

### 技能组监控

预览 `/group/preview`

详情 `/group/detail`

### 自动外呼

开启 `/task/start`

关闭 `/task/stop`

### 接续号码

设置 `/agent/translate`
取消 `/agent/translate`

### 音频

上传 `/media`
下载 `/media`


### 配置

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

### 下载项目依赖

```
go get ./...
```