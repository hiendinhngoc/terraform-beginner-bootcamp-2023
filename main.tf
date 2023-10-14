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
  cloud {
    organization = "CoderPush"

    workspaces {
      name = "terra-house-1"
    }
  }
}

provider "terratowns" {
  endpoint = var.terratowns_endpoint
  user_uuid = var.user_uuid
  token = var.terratowns_access_token
}

module "home_arcanum_hosting" {
  source = "./modules/terrahome_aws"
  user_uuid = var.user_uuid
  public_path = var.arcanum.public_path
  content_version = var.arcanum.content_version
}

resource "terratowns_home" "home" {
  name = "Empower Pioneers"
  description = <<DESCRIPTION
CoderPush is a remote-first software consultant for businesses.
From small companies to big enterprises, we support them to scale their products and services.
With no fixed address, we free ourselves from the physical location to extend our abilities in supporting businesses.
We have scaled our own dream of reaching the world to be true. Now it's your chance. Let us give you a helping hand.
  DESCRIPTION
  domain_name = module.home_arcanum_hosting.domain_name
  # domain_name = "coderpush.cloudfront.net"
  town = "missingo"
  content_version = var.arcanum.content_version
}

module "home_payday_hosting" {
  source = "./modules/terrahome_aws"
  public_path = var.payday.public_path
  user_uuid = var.user_uuid
  content_version = var.payday.content_version
}

resource "terratowns_home" "home_payday" {
  name = "Pushing Boundary"
  description = <<DESCRIPTION
scale your team. nail your dream.
  DESCRIPTION
  domain_name = module.home_payday_hosting.domain_name
  # domain_name = "coderpush.cloudfront.net"
  town = "missingo"
  content_version = var.payday.content_version
}