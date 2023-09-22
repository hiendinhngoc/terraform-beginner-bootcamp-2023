terraform {
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
    aws = {
      source = "hashicorp/aws"
      version = "5.17.0"
    }
  }
}

provider "aws" {
  
}

provider "random" {
  # Configuration options
}

resource "random_string" "bucket_name" {
  lower            = true
  upper            = false
  length           = 32
  special          = false
  override_special = ""
}

resource "aws_s3_bucket" "example" {
  # https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html
  bucket = random_string.bucket_name.result
}

output "random_bucket_name" {
  value = random_string.bucket_name.result
}