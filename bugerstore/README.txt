Requirement

  group: burgerstore.dev
  kind/type: BurgerStore
  version: v1alpha1

Operator: BurgerStore

mkdir -p satya.com/burgerstore
cd /Users/sdillikar/go/src/satya.com/burgerstore
go mod init satya.com/burgerstore

go env
GO111MODULE="on"


mkdir -p pkg/api/bugerstore.dev/v1alpha1

tree pkg
pkg
└── api
    └── bugerstore.dev
        └── v1alpha1
            ├── doc.go
            ├── register.go
            └── types.go


execDir=/Users/sdillikar/go/src/code-generator
modulename=satya.com/burgerstore

code-dir = pkg/apis

sub-directory should be "burgerstore.dev/v1alpha1" because this is used in the below command "burgerstore.dev:v1alpha1"


"${execDir}"/generate-groups.sh all satya.com/burgerstore/pkg/generated satya.com/burgerstore/pkg/apis burgerstore.dev:v1alpha1 --go-header-file "${execDir}"/hack/boilerplate.go.txt --output-base=.

#copy satya.com over to pkg
mv satya.com/burgerstore/pkg/generated pkg/
mv satya.com/burgerstore/pkg/apis/burgerstore.dev/v1alpha1/zz_generated.deepcopy.go pkg/apis/burgerstore.dev/v1alpha1/


controller-gen tool

Install
go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5
go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5
which controller-gen
/Users/sdillikar/go//bin/controller-gen

generate CRDs:
 controller-gen paths=satya.com/burgerstore/pkg/apis/burgerstore.dev/v1alpha1/ crd:trivialVersions=true crd:crdVersions=v1 output:crd:artifacts:config=manifests

 see folders config and maninfest
 
vi kind-config-with-ip.yaml

kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  apiServerAddress: "192.168.1.211"


kind create cluster --name local-ip-cl1 --config kind-config-with-ip.yaml
kind export kubeconfig  --name=local-ip-cl1  --kubeconfig=./kubeconfig-kind-local-ip-cl1.yaml