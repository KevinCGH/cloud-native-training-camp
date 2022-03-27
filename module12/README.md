# Module 12 Homework

把我们的 httpserver 服务以 Istio Ingress Gateway 的形式发布出来。以下是你需要考虑的几点：

- [x] 如何实现安全保证
- [x] 七层路由规则
- [x] 考虑 open tracing 的接入

## 部署Istio
使用 `helm` 安装 `istio`
> https://istio.io/latest/zh/docs/setup/install/helm/
```shell
# 创建命名空间
$ kubectl create namespace istio-system

# 安装 Istio base chart
$ helm install istio-base manifests/charts/base -n istio-system

# 安装 Istio discovery chart
$ helm install istiod manifests/charts/istio-control/istio-discovery \
    --set global.hub="docker.io/istio" \
    --set global.tag="1.13.2" \
    -n istio-system
    
$ kubectl get pods -n istio-system
```

## 通过`istio gateway`部署 httpserver 服务
```shell
$ kubectl create ns mod12
$ kubectl create -f ./deployment.yaml -n mod12
$ kubectl create -f ./istio-gateway.yaml -n mod12
```

查看服务
```shell
$ kubectl get svc -n mod12
istio-ingressgateway LoadBalancer 10.137.149.50
```

## 七层路由规则
```shell
$ kubectl create -f ./nginx.yaml -n mod12
$ kubectl apply -f ./istio-gateway-v2.yaml -n mod12
```

## Open Tracing 接入
接入istio 自带的监控组件和面板

```shell
$ kubectl create -f istio-1.13.2/samples/addons/prometheus.yaml
```