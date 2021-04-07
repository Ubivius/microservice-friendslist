# Microservice-friendslist Helm Charts

## Usage

[Helm](https://helm.sh) must be installed to use the charts.
Please refer to Helm's [documentation](https://helm.sh/docs/) to get started.

Once Helm is set up properly, add the repo as follows:

## Get Repo Info

```console
helm repo add microservice-friendslist <path to chartmuseum or ArtifactHub>
helm repo update
```

_See [helm repo](https://helm.sh/docs/helm/helm_repo/) for command documentation._

## Installing the Chart

To install the chart directly from the Github repo with the release name `microservice-friendslist`:

```console
helm install microservice-friendslist chart/ --values chart/values.yaml
```

To install the chart from ArtifactHub with the release name `microservice-friendslist`:

```console
helm install microservice-friendslist <org name>/microservice-friendslist --version <version> 
```

## Uninstalling the Chart

To uninstall/delete the microservice-friendslist deployment:

```console
helm delete microservice-friendslist
```
