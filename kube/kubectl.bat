kubectl run db --image mongo

kubectl get pods

docker exec -it k3d-mycluster-server-0 ctr container ls | grep mongo 

kubectl delete pod db


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