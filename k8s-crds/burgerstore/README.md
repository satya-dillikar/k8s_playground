# K8s CRD : Burgerstore

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* Describe any prerequisites, libraries, OS version, etc., needed before installing program.
* ex. Windows 10

### Installing

* How/where to download your program
* Any modifications needed to be made to files/folders

### Executing program

```
kubectl apply -f burgerstore-crd.yml
kubectl api-resources | grep burger
```

```
kubectl apply -f burgerstores-admin-edit-role.yml
kubectl apply -f burgerstores-admin-view-role.yml
```

```
kubectl get clusterrole | grep burger
kubectl describe clusterrole  burgerstores-admin-edit
```

```
kubectl apply -f burgerstore-object.yml
kubectl get BurgerStore
```

```
kubectl apply -f burgerstore-object2.yml
kubectl get BurgerStore
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
* [kubernetes-custom-resources](https://www.tutorialworks.com/kubernetes-custom-resources/)