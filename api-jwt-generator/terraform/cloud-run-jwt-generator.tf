terraform {
  backend "remote" {
    organization = "alexandre"

    workspaces {
      name = "api-jwt-generator"
    }
  }
}

provider "google-beta" {
  project = "targaren-cli"
  region  = "us-central1"

}

resource "google_project_service" "run_api" {
  service = "run.googleapis.com"
}

resource "google_secret_manager_secret" "secret" {
  provider  = google-beta
  
  secret_id = "secret"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret-version-data" {
  provider = google-beta

  secret      = google_secret_manager_secret.secret.name
  secret_data = "secret-data"
}

resource "google_secret_manager_secret_iam_member" "secret-access" {
  provider = google-beta

  secret_id  = google_secret_manager_secret.secret.id
  role       = "roles/secretmanager.secretAccessor"
  member     = "GCP_MEMBER"
  depends_on = [google_secret_manager_secret.secret]
}

resource "google_cloud_run_service" "default" {
  provider = google-beta
  name     = "api-jwt-generator"
  location = "us-central1"

  template {
    spec {
      containers {
        image = var.container_registry_docker_image
        env {
          name  = "JWT_AUDIENCE"
          value = var.jwt_audience
        }
        env {
          name  = "SERVICE_ACCOUNT_CLIENT_EMAIL"
          value = var.service_account_client_email
        }
        env {
          name  = "SERVICE_ACCOUNT_PRIVATE_KEY"
          value = var.service_account_private_key
        }
        env {
          name = "SECRET_ENV_TEST"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.secret.secret_id
              key  = "1"
            }
          }
        }
        ports {
          name           = "http1"
          container_port = "8080"
          protocol       = "TCP"
        }
        resources {
          limits = {
            "cpu"    = "1000m"
            "memory" = "256Mi"
          }
          requests = {
            "cpu"    = "1000m"
            "memory" = "128Mi"
          }
        }
      }
      timeout_seconds = 30
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "4"
      }
    }
  }

  metadata {
    annotations = {
      "run.googleapis.com/launch-stage" = "BETA"
      cloud-run                           = "jwt-generator"
      "run.googleapis.com/ingress"        = "all"
      "run.googleapis.com/ingress-status" = "all"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  depends_on                 = [google_project_service.run_api, google_secret_manager_secret_version.secret-version-data]
  autogenerate_revision_name = true

}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.default.location
  project  = google_cloud_run_service.default.project
  service  = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

# Return service URL
output "url" {
  value = google_cloud_run_service.default.status[0].url
}