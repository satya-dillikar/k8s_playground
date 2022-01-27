https://kubernetesbyexample.com/jobs/

curl -O https://raw.githubusercontent.com/openshift-evangelists/kbe/main/specs/jobs/job.yaml
kubectl apply -f job.yaml
kubectl get po
kubectl get jobs
kubectl describe jobs/countdown

kubectl get pods
NAME              READY   STATUS      RESTARTS   AGE
countdown-2wf5t   0/1     Completed   0          2m43s

kubectl logs countdown-2wf5t
