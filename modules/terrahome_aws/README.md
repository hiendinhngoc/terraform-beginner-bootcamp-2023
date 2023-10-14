## Terrahouse AWS

```tf
module "home_payday" {
  source = "./modules/terrahouse_aws"
  public_path = var.payday_public_path
  user_uuid = var.user_uuid
  content_version = var.content_version
}
```

The public directory expects the following:
- index.html
- error.html
- assets

All top level files in assets will be compied, but not any subdirectories.