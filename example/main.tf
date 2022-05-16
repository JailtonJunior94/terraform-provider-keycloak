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
  realm        = "realm_terraform"
  display_name = "Realm criado via terraform"
}

resource "keycloak_client_scope" "client_scope_test" {
  realm_id    = keycloak_realm.realm_test.id
  name        = "client_scope_terraform"
  description = "Client Scope criado via Terraform"
  protocol    = "openid-connect"

  depends_on = [
    keycloak_realm.realm_test
  ]
}
