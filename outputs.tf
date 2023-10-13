output "bucket_name" {
  description = "Bucket name for the static website hosting"
  value = module.terrahouse_aws.bucket_name
}

output "website_endpoint" {
  description = "s3 static website hosting"
  value = module.terrahouse_aws.website_endpoint
}

locals {
  root_path = path.root
}

output "root_path" {
  value = local.root_path
}