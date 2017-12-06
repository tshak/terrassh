# TerraSSH

TerraSSH is a tool to make it easy to SSH into [Terraform](https://terraform.io) managed instances. This is particularly useful in development scenarios where you don't have a stable inventory of machines.

## Installation


### Prerequisites

Ensure that [Terraform](https://terraform.io) is in your path.

### Install with Golang

```
go get github.com/tshak/terrassh
```

### Install Manually

0) Download the release for your platform
0) Extract and copy to your path (e.g. `/usr/local/bin`)

> Note: While Windows binaries are available, they have not been tested. Your best bet is to use [Windows Subsystem for Linux](https://msdn.microsoft.com/en-us/commandline/wsl/install-win10). PR's welcome.

## Terraform Configuration

TerraSSH utilizes [Output Variables](https://www.terraform.io/intro/getting-started/outputs.html) to allow you to configure one or more sets of machines to easily SSH into.

There are three output variables you must set. A `prefix` is used so that you can distinguish between different groups of machines. For example, if your prefix is `web` and the variable suffix below is `_ssh_hosts`, the output variable must be named `web_ssh_hosts`. All variables are required.

| variable suffix | description |
| -| -|
| **_ssh_hosts** | A string array of hosts (IP or hostnames) |
| **_ssh_key_path** | The key path to the SSH private key |
| **_ssh_username** | The SSH username |

You must run a `terraform refresh` or `terraform apply` in order for any output variables to take affect.

See [example.tf](example.tf) for a basic example. See the [aws_example](aws_example/) folder for a full AWS example.

## Usage

After configuring Terraform, from the root of your Terraform workspace, you may run

```$ terrassh <prefix> (hostIndex)```

Run `terrassh` without any arguments to get detailed help.

### Examples
If you defined `web_hosts` which contains 3 hosts, to SSH into the last host use:

```$ terrassh web 2```

Or to SSH into the first host:

```$ terrassh web```



