# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    hashicup = {
      versions = "0.3.4"
      source = "hashicorp.com/Proletter/hashicup"
    }
  }
}

variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

data "hashicup_coffees" "all" {}

# Returns all coffees
output "all_coffees" {
  value = data.hashicup_coffees.all.coffees
}

# Only returns packer spiced latte
output "coffee" {
  value = {
    for coffee in data.hashicup_coffees.all.coffees :
    coffee.id => coffee
    if coffee.name == var.coffee_name
  }
}
