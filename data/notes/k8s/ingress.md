# Intro

* the Service object operates at OSI L4 - it only forwards TCP and UDP connections
```
$ kubectl get svc cportal
NAME      TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
cportal   ClusterIP   10.99.61.122   <none>        8080/TCP   79d
```
* if you use Services of `type: NodePort` the clients must connect to a unique port per service
* if you use Services of `type: LoadBalancer` you allocate scarse resources (IP addresses) for each service
* for HTTP (L7) based services we can do better -> Ingress

# Ingress

<img src="https://user-images.githubusercontent.com/1047259/126613022-2bb69f61-d0d6-4933-bc9e-e636071017f8.png" style="max-width:100%;height:auto">

* k8s's HTTP-based load balancing and "virtual hosting" system
* at implementation level Ingress is different from pretty much any other k8s resource object
* in particular it is split into 1) common resource specification (the Ingress object) and 2) a controller implemetation
* there is no "standard" Ingress controller built into k8s - you have to pick and install one

## Nginx Ingress Controller

* the most popular generic Ingress controller is probably the open source [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx/)
* it reads Ingress objects and merges them into NGINX config file and then signals to the NGINX process to restart
* it has many features and options exposed via annotations
* it parses HTTP request and based on `Host` header and URL path proxies the request to a service
* it's a pod running NGINX process

```
kubectl get pods -n ingress-nginx \
  -l app.kubernetes.io/name=ingress-nginx
```

# Manifests

```
# simple-ingress.yaml
# any HTTP request is forwarded to my-service
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: simple-ingress
spec:
  backend:
    serviceName: my-service
    servicePort: 8080
```

```
# host-ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: host-ingress
spec:
  rules:
  - host: my-service.example.com
    http:
      paths:
      - backend:
          serviceName: my-service
          servicePort: 8080
```

## TLS

First we need a secret with private key and certificate for a given host (my-service.example.com in this case):

```
# tls-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  creationTimestamp: null
  name: tls-secret-name
type: kubernetes.io/tls
data:
  tls.crt: <base64 encoded certificate>
  tls.key: <base64 encoded private key>
```

Then we can reference the secret in Ingress via `secretName` field:

```
# tls-ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: tls-ingress
spec:
  tls:
  - hosts:
    - my-service.example.com
    secretName: tls-secret-name
  rules:
  - host: my-service.example.com
    http:
      paths:
      - backend:
          serviceName: my-service
          servicePort: 8080
```

You can use [cert-manager](https://cert-manager.io/docs/) and e.g. [Let's Encrypt](https://letsencrypt.org/) to automate certificates management.

# Tips and tricks

If you specify duplicate or conflicting configurations of Ingress object, the behavior is undefined.

An Ingress object can only refer to an upstream (backend) service in the same namespace. However, multiple Ingress objects in different namespaces can specify subpaths for the same host. These specifications are then merged together. This means that Ingress needs to be coordinated globally across the cluster.

Get [nginx ingress](https://kubernetes.github.io/ingress-nginx/troubleshooting/) logs

```
kubectl get po -A | grep ingress
kubectl logs -n <namespace> nginx-ingress-controller-67956bf89d-fv58j
```

# Sources

* Kubernetes Up & Running (2019)
