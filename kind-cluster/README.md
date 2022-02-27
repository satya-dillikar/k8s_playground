# KIND CLUSTER with LOCAL REGISTRY

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* X

### Installing

* X

### Executing program

```
cd k8s_playground/kind-cluster
docker system prune -a
./kind-with-registry.sh
```

```
cd ../carvel/carvel-simple-app-on-kubernetes
docker build -t my_simple_app .
docker tag my_simple_app:latest localhost:5000/my_simple_app:1.0.0
docker push localhost:5000/my_simple_app:1.0.0
```

```
curl -X GET http://localhost:5000/v2/_catalog
```

```
kubectl run my-simple-app9 --image=localhost:5000/my_simple_app:1.0.0
kubectl get pods
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
* [local-registry](https://kind.sigs.k8s.io/docs/user/local-registry/)