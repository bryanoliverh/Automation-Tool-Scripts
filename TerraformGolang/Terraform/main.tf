terraform {
  backend "gcs" {
    # Bucket is passed in via cli arg. Eg, terraform init -reconfigure -backend-configuration=dev.tfbackend
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

provider "google-beta" {
  project = var.project_id
  region  = var.region
}

#provider "google" {
#  project = "my-project-id"
#  region  = "us-central1"
#}

resource "my_bucket" "example_bucket" {
  name     = "example-bucket"
  location = "US"
  project  = "my-project-id"
}
