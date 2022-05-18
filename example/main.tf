terraform {
  required_providers {
    keycloak = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/jailtonjunior/keycloak"
    }
  }
}

provider "keycloak" {
  url      = var.url
  username = var.username
  password = var.password
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
}

resource "keycloak_client" "client_api_test" {
  realm_id                 = keycloak_realm.realm_test.id
  client_scope             = keycloak_client_scope.client_scope_test.name
  base_url                 = "http://localhost:9000"
  client_id                = "client_api_terraform"
  name                     = "client_api_terraform"
  description              = "Client criado via Terraform"
  protocol                 = "openid-connect"
  public_client            = false
  service_accounts_enabled = true
}

# resource "keycloak_client" "client_web_test" {
#   realm_id                 = keycloak_realm.realm_test.id
#   client_scope             = keycloak_client_scope.client_scope_test.name
#   base_url                 = "http://localhost:8000"
#   client_id                = "client_web_terraform"
#   name                     = "client_web_terraform"
#   description              = "Client Web criado via Terraform"
#   protocol                 = "openid-connect"
#   public_client            = true
#   service_accounts_enabled = false
# }
