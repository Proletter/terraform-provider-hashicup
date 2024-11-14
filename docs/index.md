---
page_title: "Provider: hashicup"
subcategory: ""
description: |-
  Terraform provider for interacting with hashicup API.
---

# hashicup Provider

-> Visit the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers?utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) Learn tutorials for an interactive getting started experience.

The hashicup provider is used to interact with a fictional coffee-shop application, hashicup. This provider is meant to serve as an educational tool to show users how:
1. use providers to [create, read, update and delete (CRUD) resources](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers) using Terraform.
1. create a custom Terraform provider.

To learn how to re-create the hashicup provider, refer to the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers?utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) Learn tutorials.

Use the navigation to the left to read about the available resources.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "hashicup" {
  username = "education"
  password = "test123"
}
```

## Schema

### Optional

- **username** (String, Optional) Username to authenticate to hashicup API
- **password** (String, Optional) Password to authenticate to hashicup API
- **host** (String, Optional) hashicup API address (defaults to `localhost:19090`)