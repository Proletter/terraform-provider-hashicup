---
page_title: "ingredient Data Source - terraform-provider-hashicup"
subcategory: ""
description: |-
  The ingredient data source allows you to retrieve a coffee's ingredients.
---

# Data Source `ingredient`

The ingredient data source allows you to retrieve a coffee's ingredients.

## Example Usage

```terraform
data "hashicup_coffees" "all" {}

data "hashicup_ingredients" "psl" {
  coffee_id = values(hashicup_coffees.all)[0].id
}
```

## Argument Reference

- `coffee_id` - The coffee ID.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `ingredients` - A list of coffee ingredients. See [Ingredients](#ingredients) below for details.

### Ingredients

- `id` - The ingredient ID.
- `name` - The ingredient name.
- `quantity` - The ingredient quantity.
- `unit` - The ingredient unit.