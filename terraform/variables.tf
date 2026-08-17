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
  default     = "financial-tools"
}

variable "frontend_service_name" {
  description = "Cloud Run frontend service name"
  type        = string
  default     = "financial-tools-frontend"
}

variable "bff_service_name" {
  description = "Cloud Run BFF service name"
  type        = string
  default     = "financial-tools-bff"
}

variable "tax_calculator_service_name" {
  description = "Cloud Run tax calculator service name"
  type        = string
  default     = "tax-calculator"
}

variable "frontend_container_image" {
  description = "Container image URI for frontend"
  type        = string
}

variable "bff_container_image" {
  description = "Container image URI for BFF"
  type        = string
}

variable "tax_calculator_container_image" {
  description = "Container image URI for tax calculator"
  type        = string
}
