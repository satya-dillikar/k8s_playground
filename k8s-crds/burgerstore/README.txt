https://www.tutorialworks.com/kubernetes-custom-resources/


kubectl apply -f burgerstore-crd.yml
kubectl api-resources | grep burger

kubectl apply -f burgerstores-admin-edit-role.yml

kubectl apply -f burgerstores-admin-view-role.yml

kubectl get clusterrole | grep burger
kubectl describe clusterrole  burgerstores-admin-edit

kubectl apply -f burgerstore-object.yml
kubectl get BurgerStore

kubectl apply -f burgerstore-object2.yml
kubectl get BurgerStore