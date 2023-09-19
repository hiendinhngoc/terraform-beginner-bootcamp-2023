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