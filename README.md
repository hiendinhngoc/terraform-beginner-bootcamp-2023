# Terraform Beginner Bootcamp 2023

## Sematic versioning
[Website](https://semver.org/)

## Install the Terraform CLI

### Considerations with the Terraform CLI changes
The Terraform CLI instalation instructions have changed due to gpg keyring changes. So we needed refer to the latest install CLI instructions via Terraform Documentation and change the scripting for install.

[Install Terraform CLI](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)


### Considerations for Linux Distribution

This project is built against Ubunutu.
Please consider checking your Linux Distrubtion and change acordingly to distrubtion needs. 

[How To Check OS Version in Linux](
https://www.cyberciti.biz/faq/how-to-check-os-version-in-linux-command-line/)

Example of checking OS Version:

```
$ cat /etc/os-release

PRETTY_NAME="Ubuntu 22.04.3 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.3 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
```

### Refactoring into Bash Scripts

While fixing the Terraform CLI gpg depreciation issues we notice that bash scripts steps were a considerable amount more code. So we decided to create a bash script to install the Terraform CLI.

This bash script is located here: [./bin/install_terraform_cli](./bin/install_terraform_cli)

- This will keep the Gitpod Task File ([.gitpod.yml](.gitpod.yml)) tidy.
- This allow us an easier to debug and execute manually Terraform CLI install
- This will allow better portablity for other projects that need to install Terraform CLI.

#### Shebang Considerations

A Shebang (prounced Sha-bang) tells the bash script what program that will interpet the script. eg. `#!/bin/bash`

ChatGPT recommended this format for bash: `#!/usr/bin/env bash`

- for portability for different OS distributions 
-  will search the user's PATH for the bash executable

https://en.wikipedia.org/wiki/Shebang_(Unix)

#### Execution Considerations

When executing the bash script we can use the `./` shorthand notiation to execute the bash script.

eg. `./bin/install_terraform_cli`

If we are using a script in .gitpod.yml  we need to point the script to a program to interpert it.

eg. `source ./bin/install_terraform_cli`

#### Linux Permissions Considerations

In order to make our bash scripts executable we need to change linux permission for the fix to be exetuable at the user mode.

```sh
chmod u+x ./bin/install_terraform_cli
```

alternatively:

```sh
chmod 744 ./bin/install_terraform_cli
```

https://en.wikipedia.org/wiki/Chmod

### Github Lifecycle (Before, Init, Command)

We need to be careful when using the Init because it will not rerun if we restart an existing workspace.

https://www.gitpod.io/docs/configure/workspaces/tasks

## Set up environment variable

### Show the environment variables
1. show all: `env`
2. show specific variable: `env | grep [variable name]`

### Set environment variables
There are serveral ways to set environment variable for the project:
1. Hard code: set directly variable to the file and then call it:
```sh
#!/usr/bin/env bash
PROJECT_ROOT = '/workspace/terraform-beginner-bootcamp'`

echo $PROJECT_ROOT
```
2. On terminal(**The variable only work with current windown**):
    a. Permanant
        - Use `export` to set: `export PROJECT_ROOT='/workspace/terraform-beginner-bootcamp'`
        - Use `unset` to unset: `unset PROJECT_ROOT`
    b. Temporary
        - `PROJECT_ROOT='/workspace/terraform-beginner-bootcamp' ./bin/message`

### Print variables
Use `echo`: `echo $HELLO`

### Global environment variable
Set the variable into a bash file(`.bash_profile`)

## Install AWS CLI
### Move the installation script to bash script
1. Create a bash script file call `install_aws_cli` in [bin](./bin/install_aws_cli) folder
2. Move the script which was put in `.gitpod.yml` file before to the new bash file:
```sh
    cd /workspace
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
    unzip awscliv2.zip
    sudo ./aws/install
    cd $THEIA_WORKSPACE_ROOT
```
3. Update command to run the script in the `.gitpod.yml` file:
```sh
    - name: aws-cli
        env:
            AWS_CLI_AUTO_PROMPT: on-partial
        before: |
            source ./bin/install_aws_cli
```

### Make sure the installation work well
1. Check credential configuration
Using command:
```sh
aws sts get-caller-identity
```

It should show the result like this:
```jso
{
    "UserId": "ABCAVUO12ZPVHJ3WIJ4KR",
    "Account": "123456789012",
    "Arn": "arn:aws:iam::123456789012:user/terraform-beginner-bootcamp"
}
```

If not, you should create a new user and then generate a new AWS CLI credentials for it. After you have your new access and screte keys, add them to your .env file(not .env.example) with the same format as from .env.example file. Then rerun the above command, it will work.

Remove file and folder(if existed)
```sh
rm -f '/workspace/awscliv2.zip'
rm -rf '/workspace/aws'
```
## Terraform basics

### Terraform registry
The [Terraform Registry](https://registry.terraform.io) is a public registry of Terraform providers and modules. It is hosted by HashiCorp and is the official source for finding and installing Terraform providers and modules.

- **[Providers](https://registry.terraform.io/browse/providers)** are plugins that allow Terraform to manage the resources of a specific cloud provider or infrastructure platform. For example, the AWS provider allows Terraform to manage EC2 instances, S3 buckets, and other AWS resources.
- **[Modules](https://registry.terraform.io/browse/modules)** are reusable Terraform configurations that can be called and configured by other configurations. Most modules manage a few closely related resources from a single provider. For example, the terraform-aws-modules/vpc module can be used to create a VPC with all the necessary subnets, security groups, and routing tables.

To use a provider or module from the Terraform Registry, just add it to your Terraform configuration file. When you run terraform init, Terraform will automatically download everything it needs from the registry.

The Terraform Registry includes a wide variety of providers and modules, covering a broad range of cloud providers, infrastructure platforms, and other services. This makes it easy to use Terraform to provision and manage infrastructure on virtually any platform.

Here are some of the benefits of using the Terraform Registry:

- It is the official source for finding and installing Terraform providers and modules.
- It includes a wide variety of providers and modules, covering a broad range of cloud providers, infrastructure platforms, and other services.
- It is easy to use and provides a consistent experience for finding and installing providers and modules.
- It is maintained by HashiCorp, the creators of Terraform.

Overall, the Terraform Registry is a valuable resource for Terraform users. It makes it easy to find and install the providers and modules you need to manage your infrastructure efficiently and effectively.

### Terraform commands
These are some basic commands

#### Working directory
A Terraform working directory typically contains:

- A Terraform configuration describing resources Terraform should manage. This configuration is expected to change over time.
- A hidden `.terraform` directory, which Terraform uses to manage cached provider plugins and modules, record which workspace is currently active, and record the last known backend configuration in case it needs to migrate state on the next run. This directory is automatically managed by Terraform, and is created during initialization.
- State data, if the configuration uses the default local backend. This is managed by Terraform in a `terraform.tfstate` file (if the directory only uses the default workspace) or a `terraform.tfstate.d` directory (if the directory uses multiple workspaces).
- `.terraform` directory and `.tfstate` file should be ignored to public to CVS

#### Initialize
A working directory must be initialized before Terraform can perform any operations in it (like provisioning infrastructure or modifying state).
To initilize the directory, run `terraform init`. This is the first command that should be run after writing a new Terraform configuration or cloning an existing one from version control. It is safe to run this command multiple times.
`init` can run with options, for more details check the [link](https://developer.hashicorp.com/terraform/cli/commands/init)

#### Provisioning Infrastructure
When people refer to "running Terraform," they generally mean performing these provisioning actions in order to affect real infrastructure objects. The Terraform binary has many other subcommands for a wide variety of administrative actions, but these basic provisioning tasks are the core of Terraform.
Terraform's provisioning workflow relies on three commands: `plan`, `apply`, and `destroy`

##### plan
The terraform plan command creates an execution plan, which lets you preview the changes that Terraform plans to make to your infrastructure. By default, when Terraform creates a plan it:

- Reads the current state of any already-existing remote objects to make sure that the Terraform state is up-to-date.
- Compares the current configuration to the prior state and noting any differences.
- Proposes a set of change actions that should, if applied, make the remote objects match the configuration.

For more details about usage, modes and options, please refer to the [link](https://developer.hashicorp.com/terraform/cli/commands/plan)

##### apply
The terraform apply command executes the actions proposed in a Terraform plan.

for more details about usage, modes and options of apply, please check the [link](https://developer.hashicorp.com/terraform/cli/commands/apply)

#### destroy
The terraform destroy command terminates resources managed by your Terraform project. This command is the inverse of terraform apply in that it terminates all the resources specified in your Terraform state. It does not destroy resources running elsewhere that are not managed by the current Terraform project.
Usage [details](https://developer.hashicorp.com/terraform/cli/commands/destroy)

### Terraform lock files
The dependency lock file is a file that belongs to the configuration as a whole, rather than to each separate module in the configuration. For that reason Terraform creates it and expects to find it in your current working directory when you run Terraform, which is also the directory containing the `.tf` files for the root module of your configuration.

The lock file is always named `.terraform.lock.hcl`, and this name is intended to signify that it is a lock file for various items that Terraform caches in the `.terraform` subdirectory of your working directory.

### Terraform Clouse Login issue on Gitpod Workspace
When attemting to run `terraform login` it will launch bash a wiswig wiew to generate a token. However it dose not work as expected in Gitpod VSCode in the browser.

The workaround is manually generate a token in [Terraform Cloud](https://app.terraform.io/app/settings/tokens?source=terraform-login), then create the file and open manually:
```sh
touch /home/gitpod/.terraform.d/credentials.tfrc.json
open /home/gitpod/.terraform.d/credentials.tfrc.json
```

Provide the following code(add your above generated token)

```json
{
    "credentials": {
        "app.terraform.io": {
            "token": "your token"
        }
    }
}
```

**Workaround**
Create a bash script to run the generation automatically in file `./bin/generate_tfrc_credentials`

## Git basics

### Git ignore
To hide a file or folder to be publiced to VCS, add the file name or path to directory to the `.gitignore` file:
```
# .tfstate files
*.tfstate
*.tfstate.*
```

### Git stash
If you are doing some things in the wrong branch and just relized that after you finished your work.
1. Just add your changes with:
`git add .` (add all your changes, you can add specific file with `git add /path/to/the/file`)
2. Then save it with:
`git stash save` (you can add a message after that so the context will be clearer)
3. Create your new branch
`git checkout my-new-branch`
you can create the branch remotely and then fetch it to local
`git fetch`
`git checkout my-new-branch`
4. Then move your change to the new branch:
`git stash apply`