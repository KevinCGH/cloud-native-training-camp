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

$ kubectl create -n mod8 -f imagesecret.yaml
```
### 启动集群

```shell
$ kubectl create -n mod8 -f configmap.yaml
$ kubectl create -n mod8 -f deployment.yaml
```

## Part 2
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：

- [x] Service
- [x] Ingress

> 可以考虑的细节
> - 如何确保整个应用的高可用。
> - 如何通过证书保证 `httpServer` 的通讯安全。

### Service

```shell
$ kubectl create -n mod8 -f service.yaml
$ kubectl get svc -n mod8

NAME        TYPE          CLUSTER-IP    EXTERNAL-IP     PORT(S)       AGE
httpservice LoadBalancer  10.0.144.36   20.56.239.4     80:31679/TCP  27s
```

### Ingress

#### Install ingress
```shell
$ helm repo add nginx-stable https://helm.nginx.com/stable

$ helm repo update

$ helm install ingress-nginx ingress-nginx/ingress-nginx \
    --create-namespace --namescpace ingress
```
#### Get info
```shell
$ kubectl get pod -n ingress

$ helm list -n ingress

NAME            NAMESPACE   REVISION  UPDATED                               STATUS      CHART                   APP VERSION
ingress-nginx   ingress     1         2022-03-06 14:32:23.00218 +0800 CST   deployed    ingress-nginx-4.0.13    1.1.0
```

#### Create Cert Manager
```shell
$ helm repo add jetstack https://charts.jetstack.io
$ helm repo update
$ kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.1/cert-manager.crds.yaml

$ helm install \
    cert-manager jetstack/cert-manager \
    --namespace cert-manager \
    --create-namespace \
    --version v1.7.1 

$ kubectl -n mod8 create -f issuer.yaml
```
#### Create ingress with TLS
```shell
$ kubectl -n mod8 create -f https-ingress.yaml
```