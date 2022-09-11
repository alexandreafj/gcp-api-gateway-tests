variable "jwt_audience" {
  description = "Token destination, represents the application how its going to use."
  type        = string
  sensitive   = true
}

variable "service_account_client_email" {
  description = "SAE its the email from IAM SA"
  type        = string
  sensitive   = true
}

variable "service_account_private_key" {
  description = "SAPK its the private key generate from IAM SA"
  type        = string
  sensitive   = true
}

variable "container_registry_docker_image" {
  description = "Docker image path and tag generate by the github actions."
  type        = string
  sensitive   = true
}

variable "project_id" {
  description = "Project id to create the services on TF."
  type        = string
  sensitive   = true
}

variable "gcp_member" {
  description = "gcp member key."
  type        = string
  sensitive   = true
}