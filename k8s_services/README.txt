
https://aws.amazon.com/premiumsupport/knowledge-center/eks-kubernetes-services-cluster/

macos routing table

ifconfig en0
netstat -rn
netstat -rn -i en0

list my open network ports with netstat:

netstat -anvp tcp | awk 'NR<3 || /LISTEN/'
sudo lsof -PiTCP -sTCP:LISTEN
lsof -Pn -i4 | grep LISTEN
netstat -ap tcp



pod ips
10.244.1.2 (bar-app)
10.244.2.2 (foo-app)

node ips
minikube-multinode2 192.168.58.2
minikube-multinode2-m02 192.168.58.3
minikube-multinode2-m03  192.168.58.4

service ip (http-echo-service)
cluster ip 10.103.231.165 8080/TCP

➜  kubectl get svc -o wide
NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE   SELECTOR
http-echo-service   ClusterIP   10.103.231.165   <none>        8080/TCP   5s    app=http-echo
kubernetes          ClusterIP   10.96.0.1        <none>        443/TCP    10h   <none>


node port ip 10.100.153.131   8080:30917/TCP

➜   kubectl get svc -o wide
NAME                TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE   SELECTOR
http-echo-service   NodePort    10.100.153.131    <none>        8080:30917/TCP   7s    app=http-echo
kubernetes          ClusterIP   10.96.0.1      <none>        443/TCP          10h   <none>
➜  

load balancer ip 10.100.153.131  8080:30917/TCP
➜  kind-cluster kubectl get svc -o wide
NAME                TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE   SELECTOR
http-echo-service   LoadBalancer   10.100.153.131   <pending>     8080:30917/TCP   3s    app=http-echo
kubernetes          ClusterIP      10.96.0.1        <none>        443/TCP          10h   <none>

Accessing Pods from nodes

docker exec -it minikube-multinode2 curl http://10.244.1.2:5678
docker exec -it minikube-multinode2 curl http://10.244.2.2:5678
docker exec -it minikube-multinode2-m02 curl http://10.244.1.2:5678
docker exec -it minikube-multinode2-m02 curl http://10.244.2.2:5678
docker exec -it minikube-multinode2-m03 curl http://10.244.1.2:5678
docker exec -it minikube-multinode2-m03 curl http://10.244.2.2:5678


Accessing service from nodes

cluster ip 10.103.231.165 8080/TCP

docker exec -it minikube-multinode2 curl http://10.103.231.165:8080
docker exec -it minikube-multinode2 curl http://10.103.231.165:8080
docker exec -it minikube-multinode2-m02 curl http://10.103.231.165:8080
docker exec -it minikube-multinode2-m02 curl http://10.103.231.165:8080
docker exec -it minikube-multinode2-m03 curl http://10.103.231.165:8080
docker exec -it minikube-multinode2-m03 curl http://10.103.231.165:8080
Observation : already load-balancing happens with cluster-ip also.

node port ip 10.100.153.131   8080:30917/TCP
docker exec -it minikube-multinode2 curl http://10.100.153.131 :8080
docker exec -it minikube-multinode2 curl http://10.100.153.131 :8080
docker exec -it minikube-multinode2-m02 curl http://10.100.153.131 :8080
docker exec -it minikube-multinode2-m02 curl http://10.100.153.131 :8080
docker exec -it minikube-multinode2-m03 curl http://10.100.153.131 :8080
docker exec -it minikube-multinode2-m03 curl http://10.100.153.131 :8080
Observation : already load-balancing happens with nodeport ip also.

docker exec -it minikube-multinode2 curl http://192.168.58.3:30917
docker exec -it minikube-multinode2 curl http://192.168.58.3:30917
docker exec -it minikube-multinode2 curl http://192.168.58.4:30917
docker exec -it minikube-multinode2 curl http://192.168.58.4:30917

docker exec -it minikube-multinode2-m02  curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m02  curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m02  curl http://192.168.58.4:30917
docker exec -it minikube-multinode2-m02  curl http://192.168.58.4:30917

docker exec -it minikube-multinode2-m03  curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m03 curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m03 curl http://192.168.58.4:30917
docker exec -it minikube-multinode2-m03 curl http://192.168.58.4:30917

Observation : already load-balancing happens with nodeport ip also.


load balancer ip 10.100.153.131  8080:30917/TCP

docker exec -it minikube-multinode2 curl http://10.100.153.131:8080
docker exec -it minikube-multinode2 curl http://10.100.153.131:8080
docker exec -it minikube-multinode2-m02 curl http://10.100.153.131:8080
docker exec -it minikube-multinode2-m02 curl http://10.100.153.131:8080
docker exec -it minikube-multinode2-m03 curl http://10.100.153.131:8080
docker exec -it minikube-multinode2-m03 curl http://10.100.153.131:8080
Observation : already load-balancing happens with nodeport ip also.

docker exec -it minikube-multinode2 curl http://192.168.58.3:30917
docker exec -it minikube-multinode2 curl http://192.168.58.3:30917
docker exec -it minikube-multinode2 curl http://192.168.58.4:30917
docker exec -it minikube-multinode2 curl http://192.168.58.4:30917

docker exec -it minikube-multinode2-m02  curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m02  curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m02  curl http://192.168.58.4:30917
docker exec -it minikube-multinode2-m02  curl http://192.168.58.4:30917

docker exec -it minikube-multinode2-m03  curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m03 curl http://192.168.58.3:30917
docker exec -it minikube-multinode2-m03 curl http://192.168.58.4:30917
docker exec -it minikube-multinode2-m03 curl http://192.168.58.4:30917

Observation : already load-balancing happens with nodeport ip also.

accessing external ip from nodes
minikube tunnel --cleanup
  kind-cluster kubectl get svc -o wide
NAME                TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE     SELECTOR
http-echo-service   LoadBalancer   10.100.153.131   127.0.0.1     8080:30917/TCP   6m36s   app=http-echo
kubernetes          ClusterIP      10.96.0.1        <none>        443/TCP          11h     <none>

from local laptop:
for i in {1..20}; do curl http://127.0.0.1:8080 ; done
