resource "google_cloud_run_v2_service" "frontend" {
  name     = var.frontend_service_name
  location = var.region
  project  = var.project_id

  template {
    containers {
      image = var.frontend_container_image
      ports {
        container_port = 8080
      }
      env {
        name  = "BFF_UPSTREAM"
        value = google_cloud_run_v2_service.bff.uri
      }
    }
  }

  ingress = "INGRESS_TRAFFIC_ALL"

  depends_on = [
    google_project_service.cloud_run,
    google_artifact_registry_repository.containers,
  ]
}

resource "google_cloud_run_v2_service" "bff" {
  name     = var.bff_service_name
  location = var.region
  project  = var.project_id

  template {
    containers {
      image = var.bff_container_image
      ports {
        container_port = 8080
      }
      env {
        name  = "TAX_CALCULATOR_GRPC_TARGET"
        value = google_cloud_run_v2_service.tax_calculator.uri
      }
    }
  }

  ingress = "INGRESS_TRAFFIC_ALL"

  depends_on = [
    google_project_service.cloud_run,
    google_artifact_registry_repository.containers,
  ]
}

resource "google_cloud_run_v2_service" "tax_calculator" {
  name     = var.tax_calculator_service_name
  location = var.region
  project  = var.project_id

  template {
    containers {
      image = var.tax_calculator_container_image
      ports {
        name           = "h2c"
        container_port = 8080
      }
    }
  }

  ingress = "INGRESS_TRAFFIC_INTERNAL_ONLY"

  depends_on = [
    google_project_service.cloud_run,
    google_artifact_registry_repository.containers,
  ]
}

resource "google_cloud_run_v2_service_iam_member" "frontend_public_invoker" {
  name     = google_cloud_run_v2_service.frontend.name
  project  = var.project_id
  location = var.region
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_cloud_run_v2_service_iam_member" "bff_public_invoker" {
  name     = google_cloud_run_v2_service.bff.name
  project  = var.project_id
  location = var.region
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_cloud_run_v2_service_iam_member" "tax_calculator_internal_invoker" {
  name     = google_cloud_run_v2_service.tax_calculator.name
  project  = var.project_id
  location = var.region
  role     = "roles/run.invoker"
  member   = "allUsers"
}
