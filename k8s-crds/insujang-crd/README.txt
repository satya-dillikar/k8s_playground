https://insujang.github.io/2020-02-11/kubernetes-custom-resource/

kubectl create -f test_resource.yaml
kubectl get crds

kubectl apply -f create_testresource.yaml 
kubectl get testresources

kubectl apply -f create_testresource2.yaml 
kubectl describe crds testresources

https://insujang.github.io/2020-02-13/programming-kubernetes-crd/


