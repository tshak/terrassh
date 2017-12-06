This folder provides an end-to-end example which builds two machines, sets up an SSH key, and configures TerraSSH. See [outputs.tf](outputs.tf) for the TerraSSH configuration.

## Running

First, export your AWS env vars:

```
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
```

This is not a heavily variablized example. You may wish to modify some aspects before running. In particular, [network.tf](network.tf) creates a VPC and subnet with routes. You may want to test this in an existing VPC.

After running `terraform plan` to make sure that this is really what you want to build, run `terraform apply`. Then, assuming `terrassh` is in your path, simply run `terrassh web_ssh` or `terrassh web_ssh 1` to SSH into the first and second instance.
