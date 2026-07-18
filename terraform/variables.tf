variable "project_id" {
  description = "Google Cloud project ID"
  type        = string
}

variable "region" {
  description = "Google Cloud region"
  type        = string
  default     = "australia-southeast1"
}

variable "artifact_registry_repository_id" {
  description = "Artifact Registry repository ID for Docker images"
  type        = string
  default     = "tax-calculator"
}

variable "cloud_run_service_name" {
  description = "Cloud Run service name"
  type        = string
  default     = "tax-calculator-frontend"
}

variable "container_image" {
  description = "Container image URI to deploy to Cloud Run"
  type        = string
}
