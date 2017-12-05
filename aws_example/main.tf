provider "aws" {
  region = "us-east-2"
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"]
}

resource "aws_key_pair" "terrassh" {
  key_name   = "terrassh"
  public_key = "${file(var.public_key_path)}"
}

resource "aws_instance" "web" {
  ami           = "${data.aws_ami.ubuntu.id}"
  instance_type = "t2.nano"
  count         = 2
  subnet_id     = "${aws_subnet.default.id}"
  key_name      = "terrassh"

  vpc_security_group_ids = [
    "${aws_security_group.default.id}"
  ]

  tags {
    Name = "terrassh-example-${count.index}"
  }
}