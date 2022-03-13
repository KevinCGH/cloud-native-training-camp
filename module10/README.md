# Module 10 Homework

- [x] 为`httpserver`添加0-2秒的随机延时
- [x] 为`httpserver`添加延时的`metric`
- [x] 将`httpserver`部署到集群中，并完成Prometheus的配置
- [x] 在Prometheus界面中查询延迟指标数据
- [ ] 创建一个grafana dashboard展现延时分配情况

## 修改 httpserver
- [metrics.go](../module02/homework/httpserver/server/metrics/metrics.go)
- [server.go](../module02/homework/httpserver/server/server.go)

## 重新部署 httpserver 到集群
```shell
$ kubectl apply -n mod8 -f ../module08/homework/01/deployment.yaml
```

## 部署 Prometheus

> [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack)

```shell
$ helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
$ helm repo update
$ helm show values prometheus-community/kube-prometheus-stack > /tmp/values.yaml

$ helm -n prometheus-stack install prometheus-community/kube-prometheus-stack

$ kubectl create secret generic additional-configs --from-file=prometheus-additional.yaml -n prometheus-stack
```

## 查询延迟指标数据

![search latency](./Screenshot%202022-03-13%20at%206.25.35%20PM.png)

## 创建 Grafana Dashboard

 > TODO