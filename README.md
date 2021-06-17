# cloud-storage
File store server.

## 分块上传与断点续传
* 分块上传：将文件分成多块，独立传输，上传完成后再合并
* 断点续传：传输暂停或异常中断后，可基于原有进度重传

### 说明
* 小文件不建议分块上传
* 可并行分块上传，且可以无序传输
* 分块上传可极大提高传输效率(设置合理的并行数量)
* 减少传输失败后重试的流量和时间

## 分块上传

## redis
```bash
docker run -itd --name redis-dev -p 6379:6379 redis
```
