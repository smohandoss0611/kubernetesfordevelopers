kubectl create -f go-demo-2-api-rs.yml

kubectl create -f go-demo-2-api-svc.yml

kubectl get all

nohup kubectl port-forward service/go-demo-2-api --address 0.0.0.0  3000:8080 > /dev/null 2>&1 &

#Please wait a few seconds before running the following command 
curl -i "0.0.0.0:3000/demo/hello"