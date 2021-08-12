# hcorp-microservices

## Consul

### Instalar consul

$ helm repo add hashicorp https://helm.releases.hashicorp.com
$ helm install consul hashicorp/consul -f config.yaml

### Ver UI

$ kubectl port-forward service/consul-server 8500:8500

## Kong

### Instalar Kong

$ helm repo add kong https://charts.konghq.com
$ helm repo update
$ helm install kong/kong --generate-name --set ingressController.installCRDs=false

## Vault

### Instalar vault

$ helm repo add hashicorp https://helm.releases.hashicorp.com
$ helm install vault hashicorp/vault
