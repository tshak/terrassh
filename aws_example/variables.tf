variable "public_key_path" {
  description = "Path to file containing public key"
  default     = "~/.ssh/id_rsa.pub"
}

variable "private_key_path" {
  description = "Path to file containing private key"
  default     = "~/.ssh/id_rsa"
}