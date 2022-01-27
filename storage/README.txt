
---------

helm repo add helm-repo https://shefeekj.github.io/helm-repo/
helm install wp --set storageClassName=<storage-class-name> helm-repo/wordpress
Example :  helm install wp --set storageClassName=standard helm-repo/wordpress 




helm repo add harbor-helm-repo https://projects.registry.vmware.com/chartrepo/vida
kubectl get storageclass
NAME                 PROVISIONER              AGE
hostpath (default)   docker.io/hostpath       264d
standard (default)   csi.vsphere.vmware.com   2d22h
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
