* speficication of how groups of pods can communicate with each other and other network endpoints
* in other words, they are the firewall rules of a k8s cluster
* by default, all pods in a cluster are non-isolated and accept traffic from any source
* as soon as you have a network policy that selects some pods, those pods become isolated and reject any traffic that is not allowed by a network policy
* network policies are additive, i.e. they don't conflict and evaluation order has no effect
* network policies are effective only within a namespace

*WARNING*: To use network policies, you must be using a [network plugin](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/) which supports NetworkPolicy. Creating a NetworkPolicy resource without a controller that implements it will have no effect.

Only coffeeshop can talk to payment-processor API:

```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: api-allow
spec:
  podSelector:
    matchLabels:
      app: payment-processor
      role: api
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: coffeeshop
```

<img src="https://user-images.githubusercontent.com/1047259/130800106-f114c4ad-04a0-42ef-9f23-54800c95ad96.png" style="max-width:100%;height:auto;"> 

Isolating all pods within a namespace:

```
# Deny all ingress and all egress traffic.
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-all
spec:
  podSelector: {} # all pods in NS
  policyTypes:    # types of traffic
  - Ingress
  - Egress
```

More

* https://kubernetes.io/docs/concepts/services-networking/network-policies/
* https://docs.giantswarm.io/guides/limiting-pod-communication-with-network-policies/
