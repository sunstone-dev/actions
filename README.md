# Sunstone template update action

TODO

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

uses: sunstone-dev/actions@v1
with:
  username: 'rusenask'
  apiKey: 'my-key'
  repositoryId: 'ff6cce17-ba30-4d2f-a20e-186e7e1db0ad'
  templatePath: path/to/template.yml