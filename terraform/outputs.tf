output "cloud_run_url" {
  description = "URL of the deployed Cloud Run frontend service"
  value       = google_cloud_run_v2_service.frontend.uri
}

output "artifact_registry_repository" {
  description = "Artifact Registry repository resource name"
  value       = google_artifact_registry_repository.frontend.id
}
