# Pixiu Autoscaler

`pixiu-autoscaler` 通过为 `workload` 添加 `annotations` 的方式，自动维护对应 `HorizontalPodAutoscaler` 的生命周期.

## Prerequisites

在 `kubernetes` 集群中， 需要先完成 `Metrics Server` 组件的安装，请参考 [Metrics Server](https://github.com/kubernetes-incubator/metrics-server)

`kubectl top node/pod` 验证 `Metrics Server` 已成功安装

``` bash
# kubectl top node
NAME          CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%
kubez         333m         16%    1225Mi          65%

# kubectl top pod
NAME                     CPU(cores)   MEMORY(bytes)
test1-54cd855b77-q67h6   1m           3Mi
```
## 如需使用Prometheus数据作为监控指标

在 `kubernetes` 集群中， 需要先完成 `prometheus` 组件的安装，请参考 [Prometheus](https://github.com/prometheus/prometheus)

在 `kubernetes` 集群中， 需要先完成 `prometheus-adapter` 组件的安装，请参考 [prometheus-adapter](https://github.com/kubernetes-sigs/prometheus-adapter)

以上组件使用 [kubez-ansible](https://github.com/pixiu-io/kubez-ansible) 均可自动安装
## Installing

The steps can be found in [Installation](./deploy)

## Getting Started

在 `workload` 的 `annotations` 中添加所需注释即可自动创建对应的 `HPA`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    ...
    # 可选，默认 minReplicas 为 1
    hpa.caoyingjunz.io/minReplicas: "2"  # MINPODS
    # 可选，默认 maxReplicas 为 6
    hpa.caoyingjunz.io/maxReplicas: "6"  # MAXPODS
    ...

    # 支持多种 TARGETS 类型，若开启，至少选择一种，可同时支持多个 TARGETS
    # CPU, in cores. (500m = .5 cores)
    # Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)

    # 使用率 examples
    cpu.hpa.caoyingjunz.io/targetAverageUtilization: "80"
    memory.hpa.caoyingjunz.io/targetAverageUtilization: "60"

    # 使用值 examples
    cpu.hpa.caoyingjunz.io/targetAverageValue: 600m
    memory.hpa.caoyingjunz.io/targetAverageValue: 60Mi

    # targetCustomMetric 请填写Prometheus指标名称，请确认能实际拿到数据。
    # prometheus examples
    prometheus.hpa.caoyingjunz.io/targetAverageValue: "80"
    hpa.caoyingjunz.io/targetCustomMetric: 'node_load1'
    ...
  name: test1
  namespace: default
  ...
```

`pixiu-autoscaler-controller` 会根据注释的变化，自动同步 `HPA` 的生命周期.

Copyright 2019 caoyingjunz (cao.yingjunz@gmail.com) Apache License 2.0
