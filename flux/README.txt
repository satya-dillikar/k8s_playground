
Flux - GitOps
----------------
https://searchitoperations.techtarget.com/tutorial/Try-out-this-GitOps-tutorial-with-Flux-and-Kubernetes

https://fluxcd.io/docs/get-started/

https://github.com/stefanprodan/podinfo

export GITHUB_TOKEN=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 <-dont use this
export GITHUB_USER=satya-dillikar

flux bootstrap github \
 --owner=$GITHUB_USER \
 --repository=gitops-Kubernetes-demo \
 --branch=main \
 --path=./clusters/satya-eks-west2-cl1 \
 --personal

export GITHUB_TOKEN=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

flux bootstrap github \
 --owner=$GITHUB_USER \
 --repository=gitops-Kubernetes-demo \
 --branch=main \
 --path=./clusters/my-cluster \
 --personal


git clone https://github.com/$GITHUB_USER/gitops-Kubernetes-demo
Username for 'https://github.com': satya-dillikar
Password for 'https://satya-dillikar@github.com':  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

flux create source git stefan-podinfo \
 --url=https://github.com/stefanprodan/podinfo \
 --branch=master \
 --interval=30s \
 --export > ./clusters/satya-eks-west2-cl1/stefan-podinfo-source.yaml



flux create source git satya-podinfo2 \
 --url=https://github.com/satya-dillikar/podinfo2 \
 --branch=main \
 --interval=30s \
 --export > ./clusters/satya-eks-west2-cl1/satya-podinfo-source.yaml

flux create source git satya-podinfo2  --url=https://github.com/satya-dillikar/podinfo2   --branch=main  --interval=30s  --export > ./clusters/satya-eks-west2-cl1/satya-podinfo-source.yaml

flux create kustomization satya-podinfo2 \
 --source=satya-podinfo2 \
 --path="./kustomize" \
 --prune=true \
 --validation=client \
 --interval=5m \
 --export > ./clusters/satya-eks-west2-cl1/podinfo-kustomization.yaml

flux create kustomization satya-podinfo2  --source=GitRepository/satya-podinfo2  --path="./kustomize"  --prune=true  --validation=client  --interval=5m 

 kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/$VERSION_KUBE_DASHBOARD/aio/deploy/recommended.yaml



flux create source helm podinfo \
--namespace=default \
--url=https://stefanprodan.github.io/podinfo \
--interval=10m

flux create kustomization satya-podinfo2  --source=podinfo2  --path="./kustomize"  --prune=true  --validation=client  --interval=5m 

—

flux create source git podinfo \
  --url=https://github.com/stefanprodan/podinfo \
  --branch=master \
  --interval=30s \
  --export > ./clusters/my-cluster/app-podinfo/podinfo-source.yaml

flux create source helm podinfo --namespace=default --url=https://stefanprodan.github.io/podinfo --interval=10m


flux create kustomization podinfo \
  --source=podinfo \
  --path="./kustomize" \
  --prune=true \
  --validation=client \
  --interval=5m \
  --export > ./clusters/my-cluster/app-podinfo/podinfo-kustomization.yaml

flux create kustomization podinfo --source=podinfo --path="./kustomize" --prune=true --validation=client --interval=5m 

flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=Flux \
  --branch=main \
  --path=./clusters/my-cluster \
  --personal

flux create source git podinfo \
  --url=https://github.com/stefanprodan/podinfo \
  --branch=master \
  --interval=30s \
  --export > ./clusters/my-cluster/podinfo-source.yaml

flux create source git podinfo --url=https://github.com/stefanprodan/podinfo --branch=master --interval=30s --export > ./clusters/my-cluster/podinfo-source.yaml



helm repo add traefik https://helm.traefik.io/traefik
helm install my-traefik traefik/traefik \
  --version 9.18.2 \
  --namespace traefik


flux create source helm traefik --url https://helm.traefik.io/traefik --namespace traefik
flux create helmrelease my-traefik --chart traefik \
  --source HelmRepository/traefik \
  --chart-version 9.18.2 \
  --namespace traefik


helm repo add satya-helm-repo https://satya-dillikar.github.io/helm-chart/
helm install satya-podinfo-app satya-helm-repo/podinfo --namespace test --create-namespace


flux create source helm satya-src-helm —url https://satya-dillikar.github.io/helm-chart/ --namespace test-flux
flux create helmrelease satya-helm-release --chart podinfo \
  --source HelmRepository/podinfo \
  --namespace test-flux

————
Provide on-prem package management solution for non-tmc/air gapped customers


helm -n cattle-fleet-system uninstall fleet
helm -n cattle-fleet-system uninstall fleet-crd


helm -n cattle-fleet-system install --create-namespace --wait \
    fleet-crd https://github.com/rancher/fleet/releases/download/v0.3.6/fleet-crd-0.3.6.tgz

helm -n cattle-fleet-system install --create-namespace --wait \
    fleet https://github.com/rancher/fleet/releases/download/v0.3.6/fleet-0.3.6.tgz


helm -n cattle-fleet-system uninstall fleet
helm -n cattle-fleet-system uninstall fleet-crd

helm -n fleet-system uninstall fleet-crd


# Ensure kubectl is pointing to the right cluster
kubectl -n cattle-fleet-system logs -l app=fleet-agent
kubectl -n cattle-fleet-system get pods -l app=fleet-agent



Multi-Cluster Apps have been deprecated as of Rancher v2.5.0 - check out the FAQ. Check out the new Rancher Continuous Delivery capability to deploy apps across multiple clusters through GitOps.

Cc-tgt7v

--set replicaCount=2 \
--set backend=http://backend-podinfo:9898/echo \

helm upgrade --install --wait frontend \
--namespace test \
podinfo2/podinfo


flux create source helm satya-src-helm —url https://satya-dillikar.github.io/helm-chart/ --namespace test-flux
flux create helmrelease satya-helm-release --chart podinfo \
  --source HelmRepository/podinfo \
  --namespace test-flux
