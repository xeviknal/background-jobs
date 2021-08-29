# background-jobs
A demo app to simulate background job processing

## Getting started

This solutions is composed out of two components:
1. **Background Jobs**: A web server with an API for Creating and Getting Jobs.
2. **Backgournd Processing**: A infinite process that processes the background jobs introduced by the previous componet. [See source](github.com/xeviknal/background-processing).

There is a third component, called [Background commons](github.com/xeviknal/background-commons), that is a dependency to share logic from both components mentioned above.

The two services are cloud-native: containerized and ready to deploy into a kubernetes cluster.

### How to install and deploy the components?

Both components follow the same structure.
1. `Dockerfile` to containerize the application.
2. `create-docker-image.sh` which make the container image out of the source code.
3. `build.yaml` Which contains all the artifacts to run the component on kubernetes.

A typical installation is:

```bash
chmod +x create_docker-image.sh
./create_docker-image.sh
kubectl apply -f build.yaml
```

#### Installing the database
This solution requires a database to persist the different states of a job. In order to install it, the [Background commons](github.com/xeviknal/background-commons) repo contains a helm chart with a MariaDB installation. The DB is installed into `mariadb` namespace.

```bash
kubectl apply -f mariadb-helm-chart.yaml
```

**Note**: In the case of the web server (bg-jobs) there is a service to expose the API in and out of the cluster.

### How to expose the web server?

The service is a NodePort kind. Therefore, in order to send traffic from outside of the cluster you need two things: cluster IP and node port.

```bash
export CLUSTER_IP=$(minikube ip)
export PORT=$(kubectl get svc background-jobs -n background-jobs -o go-template='{{(index .spec.ports 0).nodePort}}')
echo $CLUSTER_IP:$PORT
```

**Note**: This has only been tested using a minikube cluster.

## API

There are two endpoints available:

1. `GET /jobs/id`: that returns a json with all the available information of a job.
2. `POST /objects/{object_id:[0-9]+}/jobs/create`: that creates a job with an object associated. Only 1 job every 5min for the same object_id.

**Test examples**:

```bash
# Creating a job
curl -X POST http://$CLUSTER_IP:$PORT/objects/14/jobs/create
# Checking its status
curl http://$CLUSTER_IP:$PORT/jobs/362
```
