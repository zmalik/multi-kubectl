# multi-kubectl
kubectl to multiple clusters present in contexts


### Installation:


### Usage:
use the flags `--match-ctx` for any match and `--ctx` for exact match. 

## Requirements:

- kubectl 
- configured KUBECONFIG

### Examples:
```
❯ multi-kubectl get pods -A --match-ctx=kind
context : kind-kind
NAMESPACE            NAME                                         READY   STATUS    RESTARTS   AGE
kube-system          coredns-f9fd979d6-lwglv                      1/1     Running   1          44h
kube-system          coredns-f9fd979d6-tdqtf                      1/1     Running   0          44h
kube-system          etcd-kind-control-plane                      1/1     Running   1          44h
kube-system          kindnet-xdnmf                                1/1     Running   6          44h
kube-system          kube-apiserver-kind-control-plane            1/1     Running   4          44h
kube-system          kube-controller-manager-kind-control-plane   1/1     Running   31         44h
kube-system          kube-proxy-xhqn8                             1/1     Running   0          44h
kube-system          kube-scheduler-kind-control-plane            1/1     Running   28         44h
local-path-storage   local-path-provisioner-78776bfc44-8hm86      1/1     Running   27         44h

context : kind-second
NAMESPACE            NAME                                           READY   STATUS    RESTARTS   AGE
kube-system          coredns-f9fd979d6-5gdlk                        1/1     Running   0          44h
kube-system          coredns-f9fd979d6-m4ffq                        1/1     Running   0          44h
kube-system          etcd-second-control-plane                      1/1     Running   2          44h
kube-system          kindnet-zkwqj                                  1/1     Running   8          44h
kube-system          kube-apiserver-second-control-plane            1/1     Running   9          44h
kube-system          kube-controller-manager-second-control-plane   1/1     Running   23         44h
kube-system          kube-proxy-2v78d                               1/1     Running   0          44h
kube-system          kube-scheduler-second-control-plane            1/1     Running   22         44h
local-path-storage   local-path-provisioner-78776bfc44-bcpxz        1/1     Running   30         44h
```

