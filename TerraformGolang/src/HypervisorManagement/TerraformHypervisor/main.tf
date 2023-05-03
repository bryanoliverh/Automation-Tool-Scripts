provider "google" {
  credentials = file("/path/to/credentials.json")
  project     = "<YOUR-PROJECT-ID>"
  region      = "<YOUR-REGION>"
}

resource "gcp_virtual_machine" "example_vm" {
  name          = "example-vm"
  image         = "debian-cloud/debian-9"
  machine_type  = "n1-standard-1"
  zone          = "us-central1-a"
}
