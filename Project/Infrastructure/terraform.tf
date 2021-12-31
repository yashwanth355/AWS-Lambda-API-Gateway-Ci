terraform {
  required_version = "<= 1.0.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.46.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

terraform {
  backend "s3" {
    bucket = "terraform-b34"
    key    = "sample/terraform.tfstate"
    region = "us-east-1"
    dynamodb_table = "terraform"
  }
}