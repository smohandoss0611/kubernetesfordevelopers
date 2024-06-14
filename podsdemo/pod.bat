k3d cluster create  ces-sales-prod
kubectl get nodes
kubectl run ces-sales-db --image mongo
docker exec -it k3d-ces-sales-prod-server-0 ctr container ls | grep mongo 
kubectl delete pod db
kubectl create -f db.yml

kubectl describe -f db.yml

kubectl exec db -- ps aux

kubectl exec -it db -- sh

echo 'db.stats()' | mongo localhost:27017/test

exit

kubectl logs db

kubectl exec db --  pkill mongod

kubectl get pods

kubectl delete -f db.yml

kubectl get pods
