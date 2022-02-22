https://projectcontour.io/getting-started/

Install contour
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml
kubectl get pods -n projectcontour -o wide


cd contour
kind create cluster --config=kind.config.yaml
kubectl cluster-info --context kind-kind

kubectl config get-contexts
*         kind-kind                      kind-kind                  kind-kind

kubectl get nodes


wget https://projectcontour.io/quickstart/contour.yaml
wget https://projectcontour.io/examples/kuard.yaml

kubectl apply -f contour.yaml
kubectl get pods -n projectcontour -o wide

kubectl apply -f kuard.yaml
kubectl get po,svc,ing -l app=kuard
