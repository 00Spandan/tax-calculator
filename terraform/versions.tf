terraform {
  required_version = ">= 1.6.0"
  
  backend "gcs" {
    bucket = "financial-tools-502613-tfstate"
    prefix = "tax-calculator"
  }

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.40"
    }
  }
}
