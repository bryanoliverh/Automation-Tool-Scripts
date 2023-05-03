variable "credentials_file" {
  type = "string"
}

variable "project_id" {
  type = "string"
}

variable "region" {
  type = "string"
}

variable "instance_name" {
  type        = "string"
  description = "The name of the instance to create"
}

variable "zone" {
  type        = "string"
  description = "The zone in which to create the instance"
}

variable "public_key_file" {
  type        = "string"
  description = "The path to the public key file to use for SSH access"
}
