# YAGES

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

```
  protoc --go_out=yages --go_opt=paths=source_relative \
  --go-grpc_out=yages --go-grpc_opt=paths=source_relative \
  yages-schema.proto
```

```
  protoc --go_out=yages --go_opt=paths=source_relative \
  yages-schema.proto
```

```
  kubectl get services -n grpc-demo
  NAME       TYPE           CLUSTER-IP      EXTERNAL-IP                                                               PORT(S)          AGE
  yages      ClusterIP      10.100.136.76   <none>                                                                    9000/TCP         123m
  yages-lb   LoadBalancer   10.100.18.75    ab35c75dc42d84921873473a40e0e2c9-2078373086.us-west-2.elb.amazonaws.com   9000:31075/TCP   27m
```

```
  grpcurl --plaintext ab35c75dc42d84921873473a40e0e2c9-2078373086.us-west-2.elb.amazonaws.com:9000 yages.Echo.Ping
{
  "text": "pong"
}
```

### Installing

* How/where to download your program
* Any modifications needed to be made to files/folders

### Executing program

* START SERVER:
```
docker run  -p 9000:9000 dsatya6/yages:0.1.0
```

* CLIENT:


```
LOCAL
grpcurl --plaintext localhost:9000 list

grpcurl --plaintext localhost:9000 describe

grpcurl --plaintext localhost:9000 Echo.Ping

grpcurl --plaintext -d '{ "text" : "ereh nuf emos" }' localhost:9000 Echo.Reverse
```


```
ON AWS LoadBalancer
grpcurl --plaintext a7cbf8ceccaf24a9ea3f8133bc226a10-1100472558.us-west-2.elb.amazonaws.com:80 list
grpcurl --plaintext a7cbf8ceccaf24a9ea3f8133bc226a10-1100472558.us-west-2.elb.amazonaws.com:80 describe
grpcurl --plaintext a7cbf8ceccaf24a9ea3f8133bc226a10-1100472558.us-west-2.elb.amazonaws.com:80 Echo.Ping
grpcurl --plaintext -d '{ "text" : "ereh nuf emos" }' a7cbf8ceccaf24a9ea3f8133bc226a10-1100472558.us-west-2.elb.amazonaws.com:80 Echo.Reverse

** INGRESS grpc
grpcurl --plaintext a5ff9bc794c9c412798924a990b33791-7c4fb9dc4b1b492e.elb.us-west-2.amazonaws.com:80 list
```

* Cluster-IP service
```
kubectl get services -n grpc-demo
NAME    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
yages   ClusterIP   10.96.145.131   <none>        80/TCP    18m

 kubectl run -it --rm grpcurl --restart=Never --image=quay.io/mhausenblas/gump:0.1 -- sh
 /go $  grpcurl --plaintext 10.96.145.131:80 Echo.Ping
{
  "text": "pong"
}
/go $  grpcurl --plaintext -d '{ "text" : "ereh nuf emos" }' 10.96.145.131:80 Echo.Reverse
{
  "text": "some fun here"
}
/go $ exit
```

## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```

## Authors

Contributors names and contact info

[@SatyaDillikar](https://twitter.com/SatyaDillikar)

## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release

## License

N/A

## Acknowledgments
* [yages](https://github.com/mhausenblas/yages)