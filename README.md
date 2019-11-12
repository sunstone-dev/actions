# Sunstone Action

A GitHub Action to automatically update your private repositories in the https://apps.sunstone.dev service. Sunstone is an easy to use template generator that doesn't require any CLI (use with `curl` or any other HTTP client or straight from the `kubectl`).

For example, to install `latest` Redis to a `default` namespace with a volume pd:

```bash
kubectl apply -f https://sunstone.dev/redis?tag=latest&redis-data-pd=my-volume&namespace=default
```

You can instrument any applications like this, for example installing a custom [Dotscience](https://dotscience.com) Tensorflow model deployment operator with your own token, you just need to:

```bash
kubectl apply -f https://sunstone.dev/dotscience?token=my-token
```

# Usage

## Create your template

Create your template if you don't have one (examples can be found here https://about.sunstone.dev/docs/).

## Create/login into your Sunstone account

You can create new private repositories here: https://apps.sunstone.dev/private-repositories. You will be able to login with your GitHub account.Create a new private repository, you will need ID for it in the action.

## Use the action

Use the `sunstone-dev/actions@v1` action in your workflow file.

Note that you should use checkout action before the `sunstone-dev`:

```yaml
name: CI

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout
      uses: actions/checkout@v1
    - name: "update"
      uses: sunstone-dev/actions@v1
      with:
        username: 'rusenask'
        apiKey: 'XXX'
        repositoryId: 'XXXXXX-XXXX-XXXX-XXXX-XXXXXX'
        templatePath: example.yaml
```

## Inputs

### `username`

**Required** Your account username.

### `apiKey`

**Required** API key for your account, get one from https://apps.sunstone.dev/account.

### `repositoryId`

**Required** Repository ID.

### `templatePath`

**Required** Path to your template (for example "deployments/my-deployment.yml").

## Outputs

### `status`

Update status

## Example usage

```yaml
uses: sunstone-dev/actions@v1
with:
  username: 'rusenask'
  apiKey: 'my-key'
  repositoryId: 'ff6cce17-ba30-4d2f-a20e-186e7e1db0ad'
  templatePath: path/to/template.yml
```
