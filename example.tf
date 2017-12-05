// Set the hosts available for SSH
// Normally this would be dynamically populated
// See the the aws_example folder for a more detailed example
output "example_ssh_hosts" {
  value =  ["10.0.0.1", "10.0.0.2", "10.0.0.3"]
}

// Set the path to the SSH private key
output "example_ssh_key_path" {
  value = "~/.ssh/id_rsa"
}

// The username associated with the SSH private key
output "example_ssh_username" {
  value = "ubuntu"
}