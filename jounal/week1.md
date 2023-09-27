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