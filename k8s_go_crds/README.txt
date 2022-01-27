https://insujang.github.io/2020-02-13/programming-kubernetes-crd/

cd /Users/sdillikar/github/projects/k8s_crds/kubernetes-test-controller
$ git clone https://github.com/insujang/kubernetes-test-controller
$ cd kubernetes-test-controller
$ docker build -t controller-test .
$ docker run -it --rm --net=host -v $KUBECONFIG:$KUBECONFIG -e KUBECONFIG=$KUBECONFIG controller-test



---------
https://insujang.github.io/2020-02-13/programming-kubernetes-crd/#fn:1

➜  k8s_crds git:(main) ✗ tree -d -L 1 .
.
├── code-generator
├── kubernetes-test-controller
├── my-k8s-test-ctrl
└── sample-controller

4 directories
➜  k8s_crds git:(main) ✗
cd /Users/sdillikar/github/projects/go_playground/k8s_crds/code-generator

go mod edit -replace=satya.com/my-k8s-test-ctrl=../my-k8s-test-ctrl

➜  code-generator git:(e95606b) ✗ cat go.mod
// This is a generated file. Do not edit directly.

module k8s.io/code-generator

go 1.12


replace satya.com/my-k8s-test-ctrl => ../my-k8s-test-ctrl


cd /Users/sdillikar/github/projects/go_playground/k8s_crds/my-k8s-test-ctrl

➜  my-k8s-test-ctrl git:(main) ✗ tree .
.
├── go.mod
├── go.sum
└── lib
    └── testresource
        ├── generated
        └── v1beta1
            ├── doc.go
            ├── register.go
            └── types.go

4 directories, 5 files
➜  my-k8s-test-ctrl git:(main) ✗ cat go.mod
module satya.com/my-k8s-test-ctrl

go 1.17

require k8s.io/apimachinery v0.22.4

➜  my-k8s-test-ctrl git:(main) ✗ ../code-generator/generate-groups.sh all satya.com/my-k8s-test-ctrl/lib/testresource/generated satya.com/my-k8s-test-ctrl/lib testresource:v1beta1 --go-header-file ../code-generator/hack/boilerplate.go.txt --output-base .
Generating deepcopy funcs
Generating clientset for testresource:v1beta1 at satya.com/my-k8s-test-ctrl/lib/testresource/generated/clientset
Generating listers for testresource:v1beta1 at satya.com/my-k8s-test-ctrl/lib/testresource/generated/listers
Generating informers for testresource:v1beta1 at satya.com/my-k8s-test-ctrl/lib/testresource/generated/informers
➜  my-k8s-test-ctrl git:(main) ✗ ls
go.mod    go.sum    lib       satya.com

cp -r satya.com/my-k8s-test-ctrl/* ./
rm -rf satya.com
rm -rf go.sum
go mod tidy
go build ./...