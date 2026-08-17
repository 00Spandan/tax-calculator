output "cloud_run_url" {
  description = "URL of the deployed Cloud Run frontend service"
  value       = google_cloud_run_v2_service.frontend.uri
}

output "frontend_url" {
  description = "URL of the deployed frontend service"
  value       = google_cloud_run_v2_service.frontend.uri
}

output "bff_url" {
  description = "URL of the deployed BFF service"
  value       = google_cloud_run_v2_service.bff.uri
}

output "tax_calculator_url" {
  description = "URL of the deployed tax calculator service"
  value       = google_cloud_run_v2_service.tax_calculator.uri
}

output "artifact_registry_repository" {
  description = "Artifact Registry repository resource name"
  value       = google_artifact_registry_repository.containers.id
}
