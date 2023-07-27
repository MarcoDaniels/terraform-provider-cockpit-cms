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

output "all_collections" {
  value = data.cockpit-cms_collections.all
}