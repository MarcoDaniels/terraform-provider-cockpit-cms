terraform {
  required_providers {
    cockpit = {
      source = "marcodaniels.com/tf/cockpit"
    }
  }
}

provider "cockpit" {
  api_endpoint = "http://localhost:8080/api"
}

resource "cockpit_collection" "coll" {
  name   = "my-collection-1"
  label  = "From Terraform"
  fields = [
    {
      name  = "title",
      label = "Title",
      type  = "Text"
    },
    {
      name  = "position",
      label = "Position",
      type  = "Number"
    },
    {
      name  = "description",
      label = "Description",
      type  = "Text"
    }
  ]
}

output "collection" {
  value = cockpit_collection.coll
}

data "cockpit_collections" "all" {}

output "all_collections" {
  value = data.cockpit_collections.all
}

/*
data "cockpit-cms_field" "names" {
  name  = "test-from-terraform"
  type  = "text"
  label = "no-label"
}
*/
