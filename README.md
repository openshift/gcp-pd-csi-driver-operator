# gcp-pd-csi-driver-operator

An operator to deploy the [GCP PD CSI Driver](https://github.com/openshift/gcp-pd-csi-driver) in OKD.

This operator is installed by the [cluster-storage-operator](https://github.com/openshift/cluster-storage-operator).

# Quick start

To build and run the operator locally:

```shell
# Create only the resources the operator needs to run via CLI
oc apply -f manifests/00_crd.yaml
oc apply -f manifests/01_namespace.yaml
oc apply -f manifests/02_credentials.yaml
oc apply -f manifests/09_cr.yaml

# Build the operator
make

# Set the environment variables
export DRIVER_IMAGE=gcr.io/gke-release/gcp-compute-persistent-disk-csi-driver:v1.0.1-gke.0
export PROVISIONER_IMAGE=quay.io/k8scsi/csi-provisioner:canary
export ATTACHER_IMAGE=quay.io/k8scsi/csi-attacher:canary
export RESIZER_IMAGE=quay.io/k8scsi/csi-resizer:canary
export SNAPSHOTTER_IMAGE=quay.io/k8scsi/csi-snapshotter:canary
export NODE_DRIVER_REGISTRAR_IMAGE=quay.io/openshift/origin-csi-node-driver-registrar:latest
export LIVENESS_PROBE_IMAGE=quay.io/openshift/origin-csi-livenessprobe:latest

# Run the operator via CLI
./gcp-pd-csi-driver-operator start --kubeconfig $MY_KUBECONFIG --namespace openshift-cluster-csi-drivers
```

To run the latest build of the operator:

```shell
# Set the environment variables
export DRIVER_IMAGE=gcr.io/gke-release/gcp-compute-persistent-disk-csi-driver:v1.0.1-gke.0
export PROVISIONER_IMAGE=quay.io/k8scsi/csi-provisioner:canary
export ATTACHER_IMAGE=quay.io/k8scsi/csi-attacher:canary
export RESIZER_IMAGE=quay.io/k8scsi/csi-resizer:canary
export SNAPSHOTTER_IMAGE=quay.io/k8scsi/csi-snapshotter:canary
export NODE_DRIVER_REGISTRAR_IMAGE=quay.io/openshift/origin-csi-node-driver-registrar:latest
export LIVENESS_PROBE_IMAGE=quay.io/openshift/origin-csi-livenessprobe:latest

# Deploy the operator and everything it needs
oc apply -f manifests/
```
