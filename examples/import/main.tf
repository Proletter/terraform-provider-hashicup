# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    hashicup = {
      version = "0.3.4"
      source  = "hashicorp.com/Proletter/hashicup"
    }
  }
}

provider "hashicup" {
  username = "education"
  password = "test123"
}

resource "hashicup_order" "sample" {}

output "sample_order" {
  value = hashicup_order.sample
}
