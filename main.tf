terraform {
  required_providers {
    terratowns = {
      source = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
#   backend "remote" {
#     hostname = "app.terraform.io"
#     organization = "CoderPush"

#     workspaces {
#       name = "terra-house-1"
#     }
#   }
# cloud {
#   organization = "CoderPush"

#   workspaces {
#     name = "terra-house-1"
#   }
# }

}

provider "terratowns" {
  endpoint = var.terratowns_endpoint
  user_uuid = var.user_uuid
  token = var.terratowns_access_token
}

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  index_html_filepath =  var.index_html_filepath
  error_html_filepath =  var.error_html_filepath
  content_version = var.content_version
  assets_path = var.assets_path
}

resource "terratowns_home" "home" {
  name = "Empower Pioneers"
  description = <<DESCRIPTION
CoderPush is a remote-first software consultant for businesses.
From small companies to big enterprises, we support them to scale their products and services.
With no fixed address, we free ourselves from the physical location to extend our abilities in supporting businesses.
We have scaled our own dream of reaching the world to be true. Now it's your chance. Let us give you a helping hand.
  DESCRIPTION
  domain_name = module.terrahouse_aws.cloudfront_url
  # domain_name = "coderpush.cloudfront.net"
  town = "missingo"
  content_version = 1
}