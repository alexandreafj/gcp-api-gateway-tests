terraform {
  backend "remote" {
    organization = "alexandre"

    workspaces {
      name = "api-gateway"
    }
  }
}

provider "google-beta" {
  project = "targaren-cli"
}

resource "google_api_gateway_api" "terraform_api_gateway" {
  api_id   = "terraform-api-gateway"
  provider = google-beta
}

resource "google_api_gateway_api_config" "terraform_api_gateway" {
  provider      = google-beta
  api           = google_api_gateway_api.terraform_api_gateway.api_id
  api_config_id = "terraform-api-gateway-cfg"

  openapi_documents {
    document {
      path     = "spec.yaml"
      contents = filebase64("./api-gateway.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "terraform_api_gateway" {
  provider   = google-beta
  api_config = google_api_gateway_api_config.terraform_api_gateway.id
  gateway_id = "terraform-api-gateway"
  region     = "us-central1"
}