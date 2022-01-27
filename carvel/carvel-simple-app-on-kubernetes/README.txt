
### Step 1: Deploying application

kapp deploy -a simple-app -f config-step-1-minimal/
kapp inspect -a simple-app --tree
kapp logs -f -a simple-app

### Step 1a: Viewing application

kubectl port-forward svc/simple-app 8080:80


### Step 1b: Modifying application configuration

Modify `HELLO_MSG` environment value from `stranger` to something else in `config-step-1-minimal/config.yml`, and run:

kapp deploy -a simple-app -f config-step-1-minimal/ --diff-changes


### Step 2: Configuration templating

kapp deploy -a simple-app -c -f <(ytt -f config-step-2-template/ -v hello_msg=another-stranger)

New message should be returned from the app in the browser.

### Step 2a: Configuration patching

kapp deploy -a simple-app -c -f <(ytt -f config-step-2-template/ -f config-step-2a-overlays/custom-scale.yml)


### Step 3: Building container images locally

#VERY IMPORTANT: Use this Terminal only for all command
eval $(minikube docker-env)
kapp deploy -a simple-app -c -f <(ytt -f config-step-3-build-local/ | kbld -f-)

### Step 3a: Modifying application source code

Uncomment `fmt.Fprintf(w, "<p>local change</p>")` line in `app.go`, and re-run above command:

kapp deploy -a simple-app -c -f <(ytt -f config-step-3-build-local/ | kbld -f-)


### Step 4: Building and pushing container images to registry

aws ecr-public  get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/satyad-pubic-registry/carvel-contents/packages/k8s-simple-app

kapp deploy -a simple-app -c -f <(ytt -f config-step-4-build-and-push/ -v push_images_repo=public.ecr.aws/satyad-pubic-registry/carvel-contents/packages/k8s-simple-app | kbld -f-)

kapp deploy -a simple-app2 -c -f <(ytt -f config-step-4-build-and-push/ -v push_images_repo=localhost:5000/k8s-simple-app | kbld -f-)


### Step 5: Clean up cluster resources

kapp delete -a simple-app

----------------------------------------------------------------
WITH LOCAL REPOSITORY
----------------------------------------------------------------

https://github.com/vmware-tanzu/carvel/tree/develop/tutorials/katacoda/kapp-controller-package-management


docker run -d -p 5000:5000 --net=bridge --restart=always --name registry registry:2
curl -X GET http://192.168.1.211:5000/v2/_catalog
export REPO_HOST="192.168.1.211:5000"

STEP2: (imgpkg)
------
cd config-step-3-build-local
mkdir -p package-contents/config/
#create new files config.yml & values.yml or copy from directory 'config-step-2-template'
cp config.yml package-contents/config/config.yml
cp values.yml package-contents/config/values.yml


mkdir -p package-contents/.imgpkg
kbld -f package-contents/config/ --imgpkg-lock-output package-contents/.imgpkg/images.yml

imgpkg push -b 192.168.1.211:5000/packages/simple-app:1.0.0 -f package-contents
curl -X GET http://192.168.1.211:5000/v2/_catalog
curl -X GET http://192.168.1.211:5000/v2/packages/simple-app/tags/list


STEP3: (imgpkg)
------
cd config-step-3-build-local
#create new files metadata.yml & 1.0.0.yml
mkdir -p my-pkg-repo/.imgpkg my-pkg-repo/packages/simple-app.corp.com
kbld -f my-pkg-repo/packages/ --imgpkg-lock-output my-pkg-repo/.imgpkg/images.yml

imgpkg push -b 192.168.1.211:5000/packages/my-pkg-repo:1.0.0 -f my-pkg-repo
curl 192.168.1.211:5000/v2/_catalog
curl 192.168.1.211:5000/v2/packages/my-pkg-repo
curl 192.168.1.211:5000/v2/packages/my-pkg-repo/tags/list

#create repo.yml file

kapp deploy -a repo -f repo.yml -y
watch kubectl get packagerepository
kubectl get packagemetadatas
kubectl get packages --field-selector spec.refName=simple-app.corp.com
kubectl get package simple-app.corp.com.1.0.0 -o yaml

#create pkginstall.yml file
kapp deploy -a default-ns-rbac -f default-ns.yml -y
kapp deploy -a pkg-demo -f pkginstall.yml -y
kubectl get pods
kubectl port-forward service/simple-app 8080:80 &
curl http://127.0.0.1:8080
