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

provider "hashicup" {
  username = "education"
  password = "test123"
}

module "psl" {
  source = "./coffee"

  coffee_name = "Packer Spiced Latte"
}

output "psl" {
  value = module.psl.coffee
}

data "hashicup_order" "order" {
  id = 1
}

output "order" {
  value = data.hashicup_order.order
}

resource "hashicup_order" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 2
  }
  items {
    coffee {
      id = 2
    }
    quantity = 3
  }
}

output "edu_order" {
  value = hashicup_order.edu
}


data "hashicup_order" "first" {
  id = 1
}

output "first_order" {
  value = data.hashicup_order.first
}
