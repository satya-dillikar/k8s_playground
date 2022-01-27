https://docs.bitnami.com/tutorials/configure-rbac-in-your-kubernetes-cluster/

minikube start -p minikube

kubectl create namespace office

-----CREATE RSA PRIVATE KEY-----
cd /Users/sdillikar/myprojects/golang-projects/k8s-rbac
openssl genrsa -out employee.key 2048
> created employee.key


-----CREATE CERTIFICATE REQUEST-----
openssl req -new -key employee.key -out employee.csr -subj "/CN=employee/O=bitnami"
> created employee.csr

-----CREATE CERTIFICATE-----
openssl x509 -req -in employee.csr -CA /Users/sdillikar/.minikube/ca.crt -CAkey /Users/sdillikar/.minikube/ca.key -CAcreateserial -out employee.crt -days 500
>created employee.crt



kubectl config set-credentials employee --client-certificate=/Users/sdillikar/myprojects/golang-projects/k8s-rbac/employee.crt  --client-key=/Users/sdillikar/myprojects/golang-projects/k8s-rbac/employee.key

kubectl config view
kubectl config get-users

kubectl config set-context employee-context --cluster=minikube --namespace=office --user=employee

kubectl config get-contexts

kubectl --context=employee-context get pods
>>Error from server (Forbidden): pods is forbidden: User "employee" cannot list resource "pods" in API group "" in the namespace "office"

kubectl create -f role-deployment-manager.yaml
kubectl get roles -n office

kubectl create -f rolebinding-deployment-manager.yaml
get rolebinding -n office

kubectl --context=employee-context run --image bitnami/dokuwiki mydokuwiki
kubectl --context=employee-context get pods

kubectl --context=employee-context get pods --namespace=default
Error from server (Forbidden): pods is forbidden: User "employee" cannot list resource "pods" in API group "" in the namespace "default"

----------------------------------------------------------------

https://theithollow.com/2019/05/20/kubernetes-role-based-access/

kubectl create namespace hollowteam

kubectl apply -f service-account.yaml
kubectl apply -f sa-hollowteam.yaml

kubectl get serviceaccount -n hollowteam


kubectl apply -f role-hollowteam-full-access.yaml
kubectl get roles -n hollowteam

kubectl apply -f rolebinding-service-acc.yaml
kubectl apply -f rolebinding-hollowteam.yaml
kubectl get rolebinding -n hollowteam

kubectl describe sa hollowteam-user -n hollowteam
#get hollowteam-user-token-pbfml from 'kubectl describe sa hollowteam-user -n hollowteam' command output
#kubectl get secret [user token] -n [namespace] -o "jsonpath={.data.token}" | base64 -D
kubectl get secret hollowteam-user-token-pbfml -n hollowteam -o "jsonpath={.data.token}" | base64 -D

>>>

kubectl describe sa hollowteam-user -n hollowteam
Name:                hollowteam-user
Namespace:           hollowteam
Labels:              <none>
Annotations:         <none>
Image pull secrets:  <none>
Mountable secrets:   hollowteam-user-token-pbfml
Tokens:              hollowteam-user-token-pbfml
Events:              <none>
➜  k8s-rbac kubectl get secret hollowteam-user-token-pbfml -n hollowteam
NAME                          TYPE                                  DATA   AGE
hollowteam-user-token-pbfml   kubernetes.io/service-account-token   3      4m46s

kubectl get secret hollowteam-user-token-pbfml -n hollowteam -o "jsonpath={.data.token}" | base64 -D
eyJhbGciOiJSUzI1NiIsImtpZCI6Ijc2b3MyVnVoX1Z6WVpzTTB1SFZHb003TGNTUDA0MHJib04yQkZPZUx4bGMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJob2xsb3d0ZWFtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImhvbGxvd3RlYW0tdXNlci10b2tlbi1wYmZtbCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJob2xsb3d0ZWFtLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI0YmM4ZjAzZS03OGMyLTQ0NjYtYjIwNi02YmFhNzNkNmUwNDgiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6aG9sbG93dGVhbTpob2xsb3d0ZWFtLXVzZXIifQ.LZqpr5l0St7LPcYU68u1gy8XDPa3uZ6PulJpQ_KP1QaQBFuQSZTSzPtj8jk-dpx682fE-OLaudcfXSlINgIsgcg-d4mo633tFwaTvU7tXtx6R8KESYRoG9W_BO1Pc-5LhGZz59YRdtkgJjuGHBOfhmk1lfuDgPWvf5hYanICPEtjiH3S9C9PR-QE35K8ISAGdsxtpF83t5uyQfvdl6osbBEMi3h2keTYVe-9iCaQ63g7Zya8hy4y--k9EQP_LSH8H-G4pdEHmGQook13fCw1TIShylJ9Rfpxrkb6FrbF9kUY2HAAHkigg__klJQ8AKjsxIsekwJV4eso3w3nB5wKag

➜  k8s-rbac kubectl get secret hollowteam-user-token-pbfml -n hollowteam -o "jsonpath={.data['ca\.crt']}"

LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCakNDQWU2Z0F3SUJBZ0lCQVRBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwdGFXNXAKYTNWaVpVTkJNQjRYRFRJeE1ETXdOVEF4TVRZeE9Gb1hEVE14TURNd05EQXhNVFl4T0Zvd0ZURVRNQkVHQTFVRQpBeE1LYldsdWFXdDFZbVZEUVRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTW9DCkRaZ0k2T2x2S2Z4TzA4ajFrbUQ3b0xTYzYwVEl2NCs3ckw4b0l0LzIrblQ4YjFJL051eG1KRU9QUFlVeVM3VVEKaU9lL0tXZDJOQzV2dE1XY2lmaUhMdm9EUUFxNTBHdC9zUVI1bzdYN2ZiZmNmdmpESHA2SEpkcHNHdGJXalRMRwpCZU5BVlhsTWdjTmJoQXZkRHk4N1lmbVhoTkxxSUpkVnA2c2JLRTJUZ0t2WnlRbmlmajZTQ1VxbitIaThvU1pmCk5QY2lCQm42Wnc5WS93M0xOcU5CbkE4dXlFcnZVajJ0dmZKc2pnYXZHZ3gwejNYSVcwQkFzbWlPdGlhRXVpY2oKa2pncmpmelFWTWRDQ1RQNU80TXN6OGIyMk94SEJqQ25HVk55dnExd0J0dUg2dUJOdGRtL0c5NGg5bTNQRlh6YgozNStwRHBYMVE1QjdQU0lhMkw4Q0F3RUFBYU5oTUY4d0RnWURWUjBQQVFIL0JBUURBZ0trTUIwR0ExVWRKUVFXCk1CUUdDQ3NHQVFVRkJ3TUNCZ2dyQmdFRkJRY0RBVEFQQmdOVkhSTUJBZjhFQlRBREFRSC9NQjBHQTFVZERnUVcKQkJSb0taendDL21DQUJBeFZRcEsrUWhZZHdsdGtEQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUF3V0VRdlc4WQpoMmd3eW5saVJ2NENuc0E2Z2R6OEpVdnBtWlFYaFIyaktkN3RzVkZPazFoSFpBMkJORG82SHlFZ1FxdFZwTGtRClVvSS9TRUR1QUVSZnhGcnZIdU9LTm1xNGUrcVEzdXZCMjVlMGU3SUJFSGpocTZiUE5uVzhJUEJBczczTjVERjkKVkJGN29sK2lnWElmUmJLYU54SERNbkJJLytQM3FRV0ticEZLWGdrdHZ5eThLcVozSW96Uk0rNlhjYlhtWWxJTwpOZ240NjlFUG55anN2ajRRakxPOG1La2JKNTBST2NHc3duTnovUlVaNHozRlAzSWZzOC9ZV1Z5WTl5QUtaeE82ClpVK0RpSHBya2Zkd1VuUGViUEhtUTVnbnBzMFhuTGUvNEFMMWljaDlvZEdnc3pVMW1WektzTHh6T2o3SmZMQWoKNkpmOFNPSXY0Ylc1Z1E9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==

<<<

kubectl get secret [user token] -n [namespace] -o "jsonpath={.data['ca\.crt']}"

kubectl get secret hollowteam-user-token-pbfml -n hollowteam -o "jsonpath={.data['ca\.crt']}"

copy base64 user-token and certificate to KUBECONFIG file.
