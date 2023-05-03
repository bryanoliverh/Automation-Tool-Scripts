provider "my_provider" {
  # Add any required configuration for your provider here
}

resource "my_provider_my_bucket" "my_bucket" {
  name     = "my-bucket"
  location = "us-central1"
  project  = "my-project-id"
}
