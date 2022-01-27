cd github/projects/k8s_playground/carvel/hellocron

git clone https://github.com/vmware-tanzu/carvel-kapp

export KUBECONFIG=~/satya-eks-west2-cl1-kubeconfig.yaml

kapp deploy -a hellocron -f cron-job.yml -y
kapp inspect -a hellocron --tree
kapp logs -f -a hellocron
#We scheduled our CronJob to output a hello message every minute, When you're done watching the logs you can use control-c (^C) to quit.
kapp delete -a hellocron -y