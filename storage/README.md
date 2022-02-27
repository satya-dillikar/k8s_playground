# Persistent Volumes

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

### Dependencies

* X

### Installing

* X

```
helm repo add helm-repo https://xyz.github.io/helm-repo/
helm install wp --set storageClassName=<storage-class-name> helm-repo/wordpress
Example :  helm install wp --set storageClassName=standard helm-repo/wordpress 
```


```
helm repo add harbor-helm-repo https://xyz/chartrepo/vida
kubectl get storageclass
NAME                 PROVISIONER              AGE
hostpath (default)   docker.io/hostpath       264d
standard (default)   csi.vsphere.vmware.com   2d22h
```
### Executing program

```
helm list
kubect get pods
helm install wp --set storageClassName=hostpath harbor-helm-repo/wordpress
helm list
kubect get pods
helm delete wp
	
helm inspect wp	
helm search repo
helm status wp

Modify Chart.yaml and
helm install wp --set storageClassName=hostpath ./wordpress
 
helm upgrade -f wordpress/Chart.yaml wp --set storageClassName=hostpath harbor-helm-repo/wordpress 
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

Inspiration, code snippets, etc.
* [Author](https://softbuild.dev)