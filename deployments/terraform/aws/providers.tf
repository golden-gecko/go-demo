terraform {
  cloud {
    organization = "golden-gecko"

    workspaces {
      name = "go-demo"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.40.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.26.0"
    }
  }
}

provider "aws" {
  region     = "eu-central-1"
  access_key = "AKIA3KK3GJL6YKC3QGHE"
  secret_key = "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"
}
