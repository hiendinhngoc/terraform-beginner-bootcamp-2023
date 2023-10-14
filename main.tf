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
  name = "My favorite artist"
  description = <<DESCRIPTION
My favorite artist is Den Vau, he is a famous rapper in Viet Nam, his real name is Nguyen Duc Cuong. He is one of the few successful artists from underground.
  DESCRIPTION
  domain_name = module.home_arcanum_hosting.domain_name
  # domain_name = "coderpush.cloudfront.net"
  town = "melomaniac-mansion"
  content_version = var.arcanum.content_version
}

# module "home_payday_hosting" {
#   source = "./modules/terrahome_aws"
#   public_path = var.payday.public_path
#   user_uuid = var.user_uuid
#   content_version = var.payday.content_version
# }

# resource "terratowns_home" "home_payday" {
#   name = "Pushing Boundary"
#   description = <<DESCRIPTION
# scale your team. nail your dream.
#   DESCRIPTION
#   domain_name = module.home_payday_hosting.domain_name
#   # domain_name = "coderpush.cloudfront.net"
#   town = "melomaniac-mansion"
#   content_version = var.payday.content_version
# }