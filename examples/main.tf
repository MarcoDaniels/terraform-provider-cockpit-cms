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

data "cockpit_collections" "all" {}

/*
provider "cockpit-cms" {
  base_url = "http://localhost:8080/api"
}

# data "cockpit-cms_collections" "all" {}

data "cockpit-cms_field" "names" {
  name  = "test-from-terraform"
  type  = "text"
  label = "no-label"
}

output "field" {
  value = data.cockpit-cms_field.names
}

resource "cockpit-cms_collection" "collection" {
  name = "from-terraform"
  data {
    fields {
      name  = data.cockpit-cms_field.names.name
      type  = data.cockpit-cms_field.names.type
      label = data.cockpit-cms_field.names.label
    }
    fields {
      name  = "test-from-terraform-1"
      type  = "number"
      label = "this is number"
    }
  }
}


/*
output "all_collections" {
  value = data.cockpit-cms_collections.all
}
*/
