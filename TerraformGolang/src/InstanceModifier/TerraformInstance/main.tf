provider "google" {
  credentials = "${file(var.credentials_file)}"
  project     = "${var.project_id}"
  region      = "${var.region}"
}

variable "credentials_file" {
  type = "string"
}

variable "project_id" {
  type = "string"
}

variable "region" {
  type = "string"
}

resource "google_compute_instance" "example_instance" {
  name         = "${var.instance_name}"
  machine_type = "f1-micro"
  zone         = "${var.zone}"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
  }

  metadata = {
    ssh-keys = "terraform:${file(var.public_key_file)}"
  }

  tags = ["terraform", "example"]
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
