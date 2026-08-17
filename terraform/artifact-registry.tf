resource "google_artifact_registry_repository" "containers" {
  project       = var.project_id
  location      = var.region
  repository_id = var.artifact_registry_repository_id
  description   = "Docker repository for Financial Tools services"
  format        = "DOCKER"

  depends_on = [google_project_service.artifact_registry]
}
