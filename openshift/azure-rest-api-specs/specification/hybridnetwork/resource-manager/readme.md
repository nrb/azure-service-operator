# hybridnetwork

> see https://aka.ms/autorest

This is the AutoRest configuration file for hybridnetwork.

## Getting Started

To build the SDKs for My API, simply install AutoRest via `npm` (`npm install -g autorest`) and then run:

> `autorest readme.md`

To see additional help and options, run:

> `autorest --help`

## For other options on installation see [Installing AutoRest](https://aka.ms/autorest/install) on the AutoRest github page.

## Configuration

### Basic Information

These are the global settings for the hybridnetwork.

```yaml
openapi-type: arm
openapi-subtype: rpaas
tag: package-2023-09-01
```
### Tag: package-2023-09-01

These settings apply only when `--tag=package-2023-09-01` is specified on the command line.

```yaml $(tag) == 'package-2023-09-01'
input-file:
  - Microsoft.HybridNetwork/stable/2023-09-01/common.json
  - Microsoft.HybridNetwork/stable/2023-09-01/configurationGroupSchema.json
  - Microsoft.HybridNetwork/stable/2023-09-01/configurationGroupValues.json
  - Microsoft.HybridNetwork/stable/2023-09-01/networkFunction.json
  - Microsoft.HybridNetwork/stable/2023-09-01/networkFunctionDefinition.json
  - Microsoft.HybridNetwork/stable/2023-09-01/networkServiceDesign.json
  - Microsoft.HybridNetwork/stable/2023-09-01/operation.json
  - Microsoft.HybridNetwork/stable/2023-09-01/publisher.json
  - Microsoft.HybridNetwork/stable/2023-09-01/pureProxyArtifact.json
  - Microsoft.HybridNetwork/stable/2023-09-01/site.json
  - Microsoft.HybridNetwork/stable/2023-09-01/siteNetworkService.json

suppressions:
  - code: PatchSkuProperty
    from: siteNetworkService.json
    reason: sku cannot be patched
```

### Tag: package-2022-01-01-preview

These settings apply only when `--tag=package-2022-01-01-preview` is specified on the command line.

```yaml $(tag) == 'package-2022-01-01-preview'
input-file:
  - Microsoft.HybridNetwork/preview/2022-01-01-preview/common.json
  - Microsoft.HybridNetwork/preview/2022-01-01-preview/device.json
  - Microsoft.HybridNetwork/preview/2022-01-01-preview/networkFunction.json
  - Microsoft.HybridNetwork/preview/2022-01-01-preview/networkFunctionVendor.json
  - Microsoft.HybridNetwork/preview/2022-01-01-preview/operation.json
  - Microsoft.HybridNetwork/preview/2022-01-01-preview/vendor.json
  - Microsoft.HybridNetwork/preview/2022-01-01-preview/vendorNetworkFunction.json
```
### Tag: package-2021-05-01

These settings apply only when `--tag=package-2021-05-01` is specified on the command line.

``` yaml $(tag) == 'package-2021-05-01'
input-file:
  - Microsoft.HybridNetwork/stable/2021-05-01/common.json
  - Microsoft.HybridNetwork/stable/2021-05-01/networkFunction.json
  - Microsoft.HybridNetwork/stable/2021-05-01/device.json
  - Microsoft.HybridNetwork/stable/2021-05-01/operation.json
  - Microsoft.HybridNetwork/stable/2021-05-01/vendor.json
  - Microsoft.HybridNetwork/stable/2021-05-01/networkFunctionVendor.json
  - Microsoft.HybridNetwork/stable/2021-05-01/vendorNetworkFunction.json
```

### Tag: package-2020-01-01-preview

These settings apply only when `--tag=package-2020-01-01-preview` is specified on the command line.

``` yaml $(tag) == 'package-2020-01-01-preview'
input-file:
  - Microsoft.HybridNetwork/preview/2020-01-01-preview/common.json
  - Microsoft.HybridNetwork/preview/2020-01-01-preview/networkFunction.json
  - Microsoft.HybridNetwork/preview/2020-01-01-preview/device.json
  - Microsoft.HybridNetwork/preview/2020-01-01-preview/operation.json
  - Microsoft.HybridNetwork/preview/2020-01-01-preview/vendor.json
  - Microsoft.HybridNetwork/preview/2020-01-01-preview/networkFunctionVendor.json
  - Microsoft.HybridNetwork/preview/2020-01-01-preview/vendorNetworkFunction.json
```

# Code Generation

## Swagger to SDK

This section describes what SDK should be generated by the automatic system.
This is not used by Autorest itself.

``` yaml $(swagger-to-sdk)
swagger-to-sdk:
  - repo: azure-sdk-for-python
  - repo: azure-sdk-for-java
  - repo: azure-sdk-for-go
  - repo: azure-sdk-for-js
  - repo: azure-sdk-for-ruby
    after_scripts:
      - bundle install && rake arm:regen_all_profiles['azure_mgmt_hybridnetwork']
  - repo: azure-resource-manager-schemas
  - repo: azure-powershell
```

## Go

See configuration in [readme.go.md](./readme.go.md)

## Python

See configuration in [readme.python.md](./readme.python.md)

## Ruby

See configuration in [readme.ruby.md](./readme.ruby.md)

## TypeScript

See configuration in [readme.typescript.md](./readme.typescript.md)

## CSharp

See configuration in [readme.csharp.md](./readme.csharp.md)