# HardwareSecurityModules

> see https://aka.ms/autorest

This is the AutoRest configuration file for hardware security module RP.

---

## Getting Started

To build the SDK for hardware security modules, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:

> `autorest`

To see additional help and options, run:

> `autorest --help`

---

## Configuration

### Basic Information

These are the global settings for the Hardware Security Modules API.

``` yaml
title: Azure HSM Resource Provider
description: The Azure management API provides a RESTful set of web services that interact with Azure HSM RP.
openapi-type: arm
tag: package-2023-12-preview
```

``` yaml !$(typescript)
modelerfour:
  flatten-models: false
```

### Tag: package-2018-10

These settings apply only when `--tag=package-2018-10` is specified on the command line.

``` yaml $(tag) == 'package-2018-10'
input-file:
- Microsoft.HardwareSecurityModules/preview/2018-10-31-preview/dedicatedhsm.json
```

### Tag: package-2021-11

These settings apply only when `--tag=package-2021-11` is specified on the command line.

``` yaml $(tag) == 'package-2021-11'
input-file:
- Microsoft.HardwareSecurityModules/stable/2021-11-30/dedicatedhsm.json
```

### Tag: package-2022-08-preview

These settings apply only when `--tag=package-2022-08-preview` is specified on the command line.

``` yaml $(tag) == 'package-2022-08-preview'
input-file:
- Microsoft.HardwareSecurityModules/preview/2022-08-31-preview/cloudhsm.json
- Microsoft.HardwareSecurityModules/stable/2021-11-30/dedicatedhsm.json
```

### Tag: package-2023-12-preview

These settings apply only when `--tag=package-2023-12-preview` is specified on the command line.

```yaml $(tag) == 'package-2023-12-preview'
input-file:
  - Microsoft.HardwareSecurityModules/preview/2023-12-10-preview/cloudhsm.json
  - Microsoft.HardwareSecurityModules/stable/2021-11-30/dedicatedhsm.json
```

# Code Generation

## Swagger to SDK

This section describes what SDK should be generated by the automatic system.
This is not used by Autorest itself.

``` yaml $(swagger-to-sdk)
swagger-to-sdk:
  - repo: azure-cli-extensions
  - repo: azure-sdk-for-python
  - repo: azure-resource-manager-schemas
  - repo: azure-sdk-for-go
  - repo: azure-powershell
  - repo: azure-sdk-for-java
```