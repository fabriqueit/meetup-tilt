# Local Environment

## Prerequisites

- Apple MacBook
- Homebrew <https://brew.sh/>

## Container Runtime

You will need a container runtime in order to run container on your machine.

### Colima

We are going to use [Colima](https://github.com/abiosoft/colima) for that.

```sh
brew install colima
```

### Docker

We will need the docker **client** in order to communicate with the runtime.

```sh
brew install docker
```

### Kubectl

To communicate with our futur k8s local cluster, we need `kubectl`.

I strongly recommand to install it with the google SDK for later purposes :smirk:

```sh
brew install google-cloud-sdk
gcloud components install kubectl gke-gcloud-auth-plugin
```

Here are some very handy tools when working with k8s :

- `kubectx` and `kubens` (<https://github.com/ahmetb/kubectx>)
- `fzf` (<https://github.com/junegunn/fzf>)
- `k9s` (<https://github.com/derailed/k9s>)

```sh
brew install kubectx k9s fzf
```

### Tilt

We are going to use [Tilt](https://github.com/tilt-dev/tilt) to deploy, manage all the resources defined in this repository.

```sh
brew install tilt
```

## Usage

### Start Colima

The default VM created by Colima has 2 CPUs, 2GiB memory and 60GiB storage.

I strongly recommand to customize the default setup, here an example.

Feel free to modify to your needs.

```sh
colima start --cpu 4 --memory 8 --kubernetes
```

### Kubernetes

Verify if the cluster is up and running

```sh
# Get cluster infos
kubectl cluster-info

## This should produces:
Kubernetes control plane is running at https://127.0.0.1:6443
CoreDNS is running at https://127.0.0.1:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://127.0.0.1:6443/api/v1/namespaces/kube-system/services/https:metrics-server:https/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

# Show available clusters
kubectl config get-contexts

### This should produces:
kubectl config get-contexts
CURRENT       NAME        CLUSTER       AUTHINFO        NAMESPACE
*             colima      colima        colima          default
```

### CI/CD Tilt

#### Postgres

To persit data after seeding the postgres database, follow those steps:

Create a `dev-iac/postgres-pv.yaml` file and modify `CHANGEME` with the absolute path of the location of this repository

```yaml
cat << EOF > dev-iac/postgres-pv.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: postgres-pv-volume
  labels:
    app: postgres
spec:
  storageClassName: local-path
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: '$PWD/dev-iac/volumes/megarepo/tmp/data-postgres'
EOF
```

You are good to go :rocket:

```sh
tilt up
```

Destroy resources

```sh
tilt down
```

## Current external services

### PostgreSQL

Current DB: `postgres`

Admin user: `admin`

Admin password: `admin`

Access URL: `postgresql://admin:admin@localhost:5432/postgres?schema=public`

### RabbitMQ

AMQP Access URL: `amqp://localhost:5672`

Admin URL: `http://localhost:15672`

Admin user: `guest`

Admin password: `guest`

### Redis

Access URL: `redis://localhost:6379`
