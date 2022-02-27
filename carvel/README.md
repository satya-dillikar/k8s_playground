# Carvel

Simple overview of use/purpose.

## Description

An in-depth paragraph about your project and overview of use.

## Getting Started

```
kind cluster
export KUBECONFIG=~/.kube/config
Install kapp-controller
curl -L -O https://github.com/vmware-tanzu/carvel-kapp-controller/releases/download/v0.25.0/release.yml
kubectl apply -f release.yml 

kubectl apply -f https://github.com/vmware-tanzu/carvel-kapp-controller/releases/latest/download/release.yml

kubectl delete -f https://github.com/vmware-tanzu/carvel-kapp-controller/releases/latest/download/release.yml
cd step1
kapp deploy -a simple-app -f config-step-1-minimal/
```


### Dependencies

* INSTALL CARVEL follow https://carvel.dev/kapp-controller/docs/latest/install/

```
kapp deploy -a kc -f https://github.com/vmware-tanzu/carvel-kapp-controller/releases/latest/download/release.yml
OR
kubectl apply -f https://github.com/vmware-tanzu/carvel-kapp-controller/releases/latest/download/release.yml
```

* verify
```
kubectl get all -n kapp-controller
kubectl api-resources --api-group packaging.carvel.dev
kubectl api-resources --api-group data.packaging.carvel.dev
kubectl api-resources --api-group kappctrl.k14s.io
```

* RBAC
```
cd carvel/hellocron
kapp deploy -a default-ns-rbac -f default-ns.yml -y
```

### Installing


* step0
```
docker pull dsatya6/k8s-simple-app
docker tag dsatya6/k8s-simple-app public.ecr.aws/satyad-pubic-registry/satyad/apps/k8s-simple-app:1.0.0
```

* AWS ECR UI , create public repository name: satyad-pubic-registry/satyad/apps/k8s-simple-app
```
docker push public.ecr.aws/satyad-pubic-registry/satyad/apps/k8s-simple-app:1.0.0
```

```
git clone https://github.com/vmware-tanzu/carvel-kapp
```

OR 

```
git clone https://github.com/vmware-tanzu/carvel-simple-app-on-kubernetes.git
```

* Login to AWS ECR
```
aws ecr-public  get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/satyad-pubic-registry/satyad/apps/k8s-simple-app

docker pull public.ecr.aws/satyad-pubic-registry/satyad/apps/k8s-simple-app:1.0.0

export KUBECONFIG=~/satya-eks-west2-cl1-kubeconfig.yaml
```

### Executing program

* STEP1:
```
cd step1

kapp deploy -a simple-app -f config-step-1-minimal/
kapp inspect -a simple-app --tree
kapp logs -f -a simple-app

sudo -E kwt net start

kubectl get pods
NAME                              READY   STATUS    RESTARTS   AGE
kwt-net                           1/1     Running   0          162m


open http://simple-app.default.svc.cluster.local/

Modify HELLO_MSG environment value from stranger to something else in config-step-1-minimal/config.yml, and run:

kapp deploy -a simple-app -f config-step-1-minimal/ --diff-changes

open http://simple-app.default.svc.cluster.local/
```

* To stop the app
```
kill/stop  command "sudo -E kwt net start"
kapp delete -a simple-app -y
kapp ls
```

* for below step2: DONOT USE iTerm on MACOS X instead use VSCODE Terminal

* STEP2:
```
cd step2/
cd satya-package-contents
ytt -f config

kapp deploy -a simple-app -c -f <(ytt -f config/)
sudo -E kwt net start
open http://simple-app.default.svc.cluster.local/


kapp deploy -a simple-app -c -f <(ytt -f config/ -v hello_msg=another-stranger)
open http://simple-app.default.svc.cluster.local/

kapp delete -a simple-app
```


* STEP2: (imgpkg)
```
cd step2/
mkdir -p satya-package-contents/.imgpkg
kbld -f satya-package-contents/config --imgpkg-lock-output satya-package-contents/.imgpkg/images.yml
#AWS ECR UI , create public repository name: carvel-contents/packages/k8s-simple-app
imgpkg push -b public.ecr.aws/satyad-pubic-registry/carvel-contents/packages/k8s-simple-app:1.0.0 -f satya-package-contents
imgpkg push -b localhost:5000/packages/k8s-simple-app:1.0.2 -f satya-package-contents

kapp deploy -a simple-app-v100-repo -f simple-app-v100-repo.yml -y
```

* STEP3: (imgpkg)
```
cd step3/
mkdir -p my-pkg-repo/.imgpkg my-pkg-repo/packages/simple-app.corp.com
kbld -f my-pkg-repo/packages/ --imgpkg-lock-output my-pkg-repo/.imgpkg/images.yml
#AWS ECR UI , create public repository name: carvel-contents/packages/step3/k8s-simple-app
imgpkg push -b public.ecr.aws/satyad-pubic-registry/carvel-contents/packages/step3/k8s-simple-app:1.0.0 -f my-pkg-repo
imgpkg push -b localhost:5000/packages/my-pkg-repo:1.0.3 -f my-pkg-repo
kapp deploy -a step3-simple-app-local-repo -f simple-app-local-repo.yml -y
kapp deploy -a step3-simple-app-v102-repo -f simple-app-v102-repo.yml -y
```


```

 config git:(main) ✗ kapp ls
Target cluster 'https://A732C981D646432B1656B7B0899EAA24.gr7.us-west-2.eks.amazonaws.com' (nodes: ip-192-168-0-16.us-west-2.compute.internal, 1+)

Apps in namespace 'default'

Name                                  Namespaces                             Lcs   Lca
default-ns-rbac                       default                                true  26d
kc                                    (cluster),kapp-controller,kube-system  true  31d
simple-app                            default                                true  2h
step3-simple-app-v102-repo            default                                true  16m
step3-simple-package-repository-ctrl  default                                true  5m

Lcs: Last Change Successful
Lca: Last Change Age

5 apps

Succeeded
➜  config git:(main) ✗ kubectl api-resources | grep carvel
packagemetadatas                           pkgm         data.packaging.carvel.dev/v1alpha1                                true         PackageMetadata
packages                                   pkg          data.packaging.carvel.dev/v1alpha1                                true         Package
internalpackagemetadatas                                internal.packaging.carvel.dev/v1alpha1                            true         InternalPackageMetadata
internalpackages                                        internal.packaging.carvel.dev/v1alpha1                            true         InternalPackage
packageinstalls                            pkgi         packaging.carvel.dev/v1alpha1                                     true         PackageInstall
packagerepositories                        pkgr         packaging.carvel.dev/v1alpha1                                     true         PackageRepository
➜  config git:(main) ✗ kubectl get packages
NAME                        PACKAGEMETADATA NAME   VERSION   AGE
simple-app.corp.com.1.0.0   simple-app.corp.com    1.0.0     4m6s
➜  config git:(main) ✗ kubectl get package
NAME                        PACKAGEMETADATA NAME   VERSION   AGE
simple-app.corp.com.1.0.0   simple-app.corp.com    1.0.0     4m11s
➜  config git:(main) ✗ kubectl get packagemetadata
NAME                  DISPLAY NAME   CATEGORIES   SHORT DESCRIPTION        AGE
simple-app.corp.com   Simple App     demo         Simple app for demoing   4m21s
➜  config git:(main) ✗ kubectl get packagemetadatas
NAME                  DISPLAY NAME   CATEGORIES   SHORT DESCRIPTION        AGE
simple-app.corp.com   Simple App     demo         Simple app for demoing   4m25s
➜  config git:(main) ✗ kubectl get packagerepositories
NAME                              AGE   DESCRIPTION
step3-simple-package-repository   15m   Reconcile succeeded
➜  config git:(main) ✗

```

```
cd step4
kapp deploy -a default-ns-rbac -f default-ns.yml -y
kapp deploy -a satya-pkg-demo -f pkginstall.yml -y

kapp -n default logs -a satya-pkg-demo
kapp -n default inspect -a satya-pkg-demo
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
* [packaging tutorial](https://carvel.dev/kapp-controller/docs/latest/packaging-tutorial/)
* [carvel-simple-app-on-kubernetes](https://github.com/vmware-tanzu/carvel-simple-app-on-kubernetes/blob/develop/README.md)
