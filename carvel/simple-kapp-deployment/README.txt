# Simple kapp-controller Deployment

This repo contains artifacts to be used in a super simple demonstration of [kapp-controller](https://github.com/vmware-tanzu/carvel-kapp-controller).


https://www.tacticalprogramming.com/kapp-controller-getting-started.html

export KUBECONFIG=~/satya-eks-west2-cl1-kubeconfig.yaml
kapp deploy -a exposed-simple-app -f crd/simple-app.yml
kubectl get pods
kubectl get svc
curl -silent a5213a4119c8243d6b7c27b096002591-356296011.us-west-2.elb.amazonaws.com