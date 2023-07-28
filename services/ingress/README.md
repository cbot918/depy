# kind-ingress template

## Requirements
1. [kubectl](https://kubernetes.io/docs/tasks/tools/): k8s client
2. [kind](https://kind.sigs.k8s.io/): k8s cluster
2. [zbpack](https://github.com/zeabur/zbpack) (choice): 只是拿來快速製作image的工具, 範例有現成的這個可以跳過, 也可以寫dockerfile製作

## Notice: 步驟流程先紀錄下來, 還來不及嚴謹多測幾次, 如果遇到問題歡迎直接找我xD

## Steps
1. create a kind-cluster with ingress feature
```
kind create cluster --name ingress --config kind-with-ingress.yaml 
```

2. install ingress-nginx to cluster
```
kubectl apply --filename https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
```
這個步驟要等一下, 好了他會顯示
```
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```
3. 快速指啟動 pod and service

一個會echo hihi 的 service
```
kubectl run hi \
  --expose \
  --image cbot918/hi \
  --port 3334
```

4. setup ingress
```
kubectl apply -f ingress.yaml
```
到這邊就都架好了
5. 測試
```
curl hello.yale.com
// hihi
```

## Zbpack(快速製作image小幫手)
### step:
1. cd app
2. zbpack .
### 說明
1. 需要 go.mod:  go mod init test
2. 需要 go.sum:  touch go.sum
3. 直接打包當前資料夾(也可以把main.go放到cmd資料夾裡 參考[官方文件](https://zeabur.com/docs/zh-TW/guides/go))
4. image會是當前目錄的名稱