terraform {
  required_providers {
    keycloak = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/jailtonjunior/keycloak"
    }
  }
}

provider "keycloak" {
  url      = "http://localhost:8080"
  username = "admin"
  password = "admin"
}

resource "keycloak_realm" "realm_test" {
  realm = "realm_com_terraform"
  display_name = "Automatizando a criação do realm com keycloak [Editado]"
}
