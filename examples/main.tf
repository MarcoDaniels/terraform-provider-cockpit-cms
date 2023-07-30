terraform {
  required_providers {
    cockpit-cms = {
      version = "0.1"
      source  = "marcodaniels.com/tf/cockpit-cms"
    }
  }
}

provider "cockpit-cms" {
  base_url = "http://localhost:8080/api"
}

data "cockpit-cms_collections" "all" {}

resource "cockpit-cms_collection" "collection" {
  name = "from-terraforms"
  data = "test-from-terraform"
}

/*
resource "cockpit-cms_collection" "collection" {
  name = "from-terraform"
  data = {
    fields = [
      {
        name     = "test-from-terraform"
        label    = ""
        type     = "text"
        default  = ""
        info     = ""
        group    = ""
        localize = false
        options  = []
        width    = "1-1"
        lst      = true
        acl      = []
      }
    ]
  }
}
*/

output "create_collection" {
  value = cockpit-cms_collection.collection
}


output "all_collections" {
  value = data.cockpit-cms_collections.all
}
