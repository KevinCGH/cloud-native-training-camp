## 错误处理

Go语言无内置 `exception` 机制，只提供`error`接口定义错误

```go
type error interface {
  Error() string
}
```



## 协程

- 进程
  - 分配系统资源（CPU时间、内存等）基本单位
  - 有独立的内存空间，切换开销大
- 线程：
- 协程