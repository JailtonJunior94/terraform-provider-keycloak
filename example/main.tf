terraform {
  required_providers {
    keycloak = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/jailtonjunior/keycloak"
    }
  }
}

provider "keycloak" {
  url      = "http://localhost:8080/"
  username = "admin"
  password = "admin"
}
