openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout \
tls.key -out tls.crt -subj "/CN=dev.satya.com/O=dev.satya.com"

Generating a 2048 bit RSA private key


https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights-Prometheus-Sample-Workloads-nginx.html

kubectl apply -f deploy.yaml

kubectl get service -n ingress-nginx
NAME                                 TYPE           CLUSTER-IP       EXTERNAL-IP                                                                     PORT(S)                      AGE
ingress-nginx-controller             LoadBalancer   10.100.225.242   ae326cef8990b42edb85deb1c0d6048e-62b483cc6a1a3617.elb.us-west-2.amazonaws.com   80:32140/TCP,443:32647/TCP   6m34s

use the addr ae326cef8990b42edb85deb1c0d6048e-62b483cc6a1a3617.elb.us-west-2.amazonaws.com  and update nginx-traffic-sample_values.yaml

export SAMPLE_TRAFFIC_NAMESPACE=nginx-sample-traffic
export EXTERNAL_IP=ae326cef8990b42edb85deb1c0d6048e-62b483cc6a1a3617.elb.us-west-2.amazonaws.com

cat nginx-traffic-sample.yaml| sed "s/{{external_ip}}/$EXTERNAL_IP/g" | \nsed "s/{{namespace}}/$SAMPLE_TRAFFIC_NAMESPACE/g" >nginx-traffic-sample_values.yaml



kubectl apply -f nginx-traffic-sample_values.yaml

curl https://ae326cef8990b42edb85deb1c0d6048e-62b483cc6a1a3617.elb.us-west-2.amazonaws.com/apple -k
curl https://$EXTERNAL_IP/apple -k
curl https://$EXTERNAL_IP/banana -k
