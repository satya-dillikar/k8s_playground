KIND CLUSTER with LOCAL REGISTRY

https://kind.sigs.k8s.io/docs/user/local-registry/


cd k8s_playground/kind-cluster
docker system prune -a
./kind-with-registry.sh

cd ../carvel/carvel-simple-app-on-kubernetes
docker build -t my_simple_app .
docker tag my_simple_app:latest localhost:5000/my_simple_app:1.0.0
docker push localhost:5000/my_simple_app:1.0.0

curl -X GET http://localhost:5000/v2/_catalog

kubectl run my-simple-app9 --image=localhost:5000/my_simple_app:1.0.0
kubectl get pods