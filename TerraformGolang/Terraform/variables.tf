variable "project_id" {
  description = "The project ID in Google Cloud to use for these resources."
}

variable "region" {
  description = "The region in Google Cloud where the resources will be deployed."
}

variable "function_name" {
  description = "The name of the function to be deployed"
}

variable "entry_point" {
  description = "The entrypoint where the function is called"
  default     = "HelloWorld"
}