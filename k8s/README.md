```
kubectl apply -f .\backend\file.deployment.yml
kubectl apply -f .\backend\file.service.loadbalancer.yml

kubectl apply -f .\frontend\file.deployment.yml
kubectl apply -f .\frontend\file.service.loadbalancer.yml

kubectl delete services frontend backend
kubectl delete deployment frontend backend

```
