# Terraform Beginner Bootcamp 2023 - Week 1

## Root Module Structure

```
PROJECT_ROOT
├── main.tf
├── variables.tf
├── providers.tf
├── outputs.tf
├── terraform.tfvars
└── README.md
```

This is a simple ASCII representation of the Terraform file structure, with each file on a new line and indented to indicate its hierarchy. The PROJECT_ROOT directory is the top-level directory for the Terraform project, and it contains all of the other files.

The following is a brief description of each file:
- *variables.tf*: This file stores the structure of input variables, which are values that can be passed into the Terraform configuration to customize the resources that are created.
- *main.tf*: This file contains the main Terraform configuration, including the definition of the resources that are to be created.
- *providers.tf*: This file defines the Terraform providers that are required for the project and their configuration.
- *outputs.tf*: This file stores the outputs of the Terraform configuration, which are values that can be used by other resources or by external scripts.
- *terraform.tfvars*: This file contains the data of variables that are to be loaded into the Terraform project.
- *README.md*: This file is a markdown file that provides documentation for the Terraform project.

Note that the terraform.tfvars file is optional, but it is generally recommended to use it to store variable values, as this makes the Terraform configuration more modular and reusable.

[Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform and Inpurt Variables

### Terraform Cloud Variables

There two types of variables in Terraform:
1. Environment Variables: The ones which set in bash terminal like: AWS credentials
2. Terraform Variables: Which we can normally set in the tfvars file

The variables should be set as sensitive so they will be hidden on the web.

### Loading Terraform Input Variables

#### var flag
We can use the `-var` to set an input variables or override a variable in the ftvars file: `tf plan -var user_uuid="exampro-user-id"`

#### var-file flag
- When we set so many variables in a var file(.tfvars) and want to spefify them all, use this flag:
`terraform apply -var-file="testing.tfvars"`

#### terraform.tfvars
This is the default file to load tin terraform variables in blunk

#### .auto.tfvars
Any files with names ending in `.auto.tfvars` or `.auto.tfvars.json` will automatically be loadded

#### order of terraform variables
Terraform loads variables in the following order, with later sources taking precedence over earlier ones:

- Environment variables
- The terraform.tfvars file, if present.
- The terraform.tfvars.json file, if present.
- Any *.auto.tfvars or *.auto.tfvars.json files, processed in lexical order of their filenames.
- Any -var and -var-file options on the command line, in the order they are provided. (This includes variables set by a Terraform Cloud workspace.)

## Dealing with Configuratin Drift
Senario: If some `.tfstate` files in your project were lost, we have to do some works to get them back:

### Fix missing resources with Terraform Import

`terraform import aws_s3_bucket.bucket bucket-name`

[Terraform Import](https://developer.hashicorp.com/terraform/cli/import)
[AWS S3 bucket import](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#import)
### Fix Manual Configuration

If someone goes and deletes or modifies cloud resource manually through ClickOps.

If we run Terraform plan is with attempt to put our infrastructure back into the expected state fixing configuration drift.

## Fix using Terraform Refresh

```sh
terraform apply -refresh-only -auto-approve
```

## Terraform Modules

### Terraform Module Structure

It's recommend to place module in a `modules` directory when locally developing modules but you can name it whatever you want.

### Passing Input Variables

We can pass input variables to our module

The module has to declare the terraform variables in its own variables.tf

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}
```

### Modules Sources

Using the source we can import the module from various places ex:
- locally
- github
- Terraform registry

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
}
```

[Module Source](https://developer.hashicorp.com/terraform/language/modules/sources)

## Working with files in Terraform

### Fileexists function

This is a built in terraform function to check the file is existed or not.
```tf
validation {
    condition     = fileexists(var.error_html_filepath)
    error_message = "The specified error.html file does not exist."
  }
```
### Path variable

### Filemd5

https://developer.hashicorp.com/terraform/language/functions/filemd5

In Terraform there is a speical variable called `path` that allows us to reference local paths(use `tf console`  to experiment the following commands):
- `path.module`: Get the path for the current module
- `path.root`: Get the path for the root module
[Path for variables](https://developer.hashicorp.com/terraform/language/expressions/references#filesystem-and-workspace-info)

```tf
resource "aws_s3_object" "index_html" {
  bucket = aws_s3_bucket.website_bucket.bucket
  key    = "index.html"
  source = "${path.root}/public/index.html"

  # The filemd5() function is available in Terraform 0.11.12 and later
  # For Terraform 0.11.11 and earlier, use the md5() function and the file() function:
  # etag = "${md5(file("path/to/file"))}"
}
```

## Terraform locals

Locals allows us to define local variables. It can be very useful when we need transform data into another format and have referenced a variable.
```tf
locals {
  s3_origin_id = "S3Origin"
}
```

[Local values](https://developer.hashicorp.com/terraform/language/values/locals)

### Terraform Data Sources

This allows use to source data from cloud resources.

This is useful when we sant to refernce cloud resources without importing them.

```tf
data "aws_caller_identity" "current" {}

output "account_id" {
  value = data.aws_caller_identity.current.account_id
}
```

[Data sources](https://developer.hashicorp.com/terraform/language/data-sources)

### Working with json

We use the jsonendcode to create the json policy inline in the hcl.
```tf
> jsonencode({"hello"="world"})
{"hello":"world"}
```
[jsonendcode](https://developer.hashicorp.com/terraform/language/functions/jsonencode)

### Lifecycle of resources

[Meta arguments lifecycle](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle)

## Terraform data

Plain data values such as Local Values and Input Variables don't have any side-effects to plan against and so they aren't valid in replace_triggered_by. You can use terraform_data's behavior of planning an action each time input changes to indirectly use a plain value to trigger replacement. [Check the link for more details](https://developer.hashicorp.com/terraform/language/resources/terraform-data)