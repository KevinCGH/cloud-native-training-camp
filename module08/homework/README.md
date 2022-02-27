# Module 8 Homework

## Part 1

编写 `Kubernetes` 部署脚本将 `httpserver` 部署到 `Kubernetes` 集群，以下是你可以思考的维度。
- [x] 优雅启动
- [x] 优雅终止
- [x] 资源需求和 QoS 保证
- [x] 探活
- [ ] 日常运维需求，日志等级
- [ ] 配置和代码分离

### 构建`httpserver` 镜像 并推送到仓库
[module03/homework](../../module03/homework/README.md)

###  环境准备

```shell
$ kubectl create ns mod8

$ kubectl create -f imagesecret.yaml
```
### 启动集群

```shell
$ kubectl create -f deployment.yaml
```

## Part 2
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：

- Service
- Ingress

> 可以考虑的细节
> - 如何确保整个应用的高可用。
> - 如何通过证书保证 `httpServer` 的通讯安全。
