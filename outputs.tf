output "bucket_name" {
  description = "Bucket name for the static website hosting"
  value = module.home_arcanum_hosting.bucket_name
}

output "website_endpoint" {
  description = "s3 static website hosting"
  value = module.home_arcanum_hosting.website_endpoint
}

output "cloudfront_url" {
  description = "The Cloudfront distribution domain name"
  value = module.home_arcanum_hosting.domain_name
}