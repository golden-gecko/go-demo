terraform {
  required_version = "1.7.5"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.40.0"
    }
  }
}

provider "aws" {
  region     = "eu-central-1"
  access_key = "AKIA3KK3GJL6YKC3QGHE"
  secret_key = "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"
}
