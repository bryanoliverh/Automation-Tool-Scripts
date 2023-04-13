# Setup the root directory of where the source code will be stored.
locals {
  root_dir = abspath("../src")
}

# Zip up our code so that we can store it for deployment.
data "archive_file" "source" {
  type        = "zip"
  source_dir  = local.root_dir
  output_path = "/tmp/function.zip"
}

# This bucket will host the zipped file.
resource "google_storage_bucket" "bucket" {
  name     = "${var.project_id}-${var.function_name}"
  location = var.region
}

# Add the zipped file to the bucket.
resource "google_storage_bucket_object" "zip" {
  # Use an MD5 here. If there's no changes to the source code, this won't change either.
  # We can avoid unnecessary redeployments by validating the code is unchanged, and forcing
  # a redeployment when it has!
  name   = "${data.archive_file.source.output_md5}.zip"
  bucket = google_storage_bucket.bucket.name
  source = data.archive_file.source.output_path
}

# The cloud function resource.
resource "google_cloudfunctions_function" "function" {
  available_memory_mb = "128"
  entry_point         = var.entry_point
  ingress_settings    = "ALLOW_ALL"

  name                  = var.function_name
  project               = var.project_id
  region                = var.region
  runtime               = "go116"
  service_account_email = google_service_account.function-sa.email
  timeout               = 20
  trigger_http          = true
  source_archive_bucket = google_storage_bucket.bucket.name
  source_archive_object = "${data.archive_file.source.output_md5}.zip"
}

# IAM Configuration. This allows unauthenticated, public access to the function.
# Change this if you require more control here.
resource "google_cloudfunctions_function_iam_member" "invoker" {
  project        = google_cloudfunctions_function.function.project
  region         = google_cloudfunctions_function.function.region
  cloud_function = google_cloudfunctions_function.function.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers"
}

# This is the service account in which the function will act as.
resource "google_service_account" "function-sa" {
  account_id   = "function-sa"
  description  = "Controls the workflow for the cloud pipeline"
  display_name = "function-sa"
  project      = var.project_id
}