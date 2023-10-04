output "bucket_name" {
  description = "Bucket name for the static website hosting"
  value = module.terrahouse_aws.bucket_name
}