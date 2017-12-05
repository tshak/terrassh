output "web_ssh_hosts" {
  value =  ["${aws_instance.web.*.public_ip}"]
}

output "web_ssh_key_path" {
  value = "${var.private_key_path}"
}

output "web_ssh_username" {
  value = "ubuntu"
}