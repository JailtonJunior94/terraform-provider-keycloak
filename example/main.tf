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

resource "keycloak_realm" "conectcar_realm" {
  realm        = "ConectCar"
  display_name = "ConectCar"
}

resource "keycloak_client_scope" "pedidos_client_scope_pf" {
  realm_id    = keycloak_realm.conectcar_realm.id
  name        = "pedidos_scope_pf"
  description = "Client Scope do contexto de Pedidos PF"
  protocol    = "openid-connect"
}

resource "keycloak_client_scope" "pedidos_client_scope_pj" {
  realm_id    = keycloak_realm.conectcar_realm.id
  name        = "pedidos_scope_pj"
  description = "Client Scope do contexto de Pedidos PJ"
  protocol    = "openid-connect"
}

resource "keycloak_client" "pedidos_bff_api" {
  realm_id                 = keycloak_realm.conectcar_realm.id
  client_scope             = keycloak_client_scope.pedidos_client_scope_pf.name
  base_url                 = "http://localhost:9000"
  client_id                = "pedidos_bff_api"
  name                     = "Pedidos BFF API"
  description              = "Client da API BFF de Pedidos [PF e PJ]"
  protocol                 = "openid-connect"
  public_client            = false
  service_accounts_enabled = true
}

resource "keycloak_client" "pedidos_web_pf" {
  realm_id                 = keycloak_realm.conectcar_realm.id
  client_scope             = keycloak_client_scope.pedidos_client_scope_pf.name
  base_url                 = "http://localhost:8000"
  client_id                = "pedidos_web_pf"
  name                     = "Pedidos Web PF"
  description              = "Client do formulário de pedidos pessoa física"
  protocol                 = "openid-connect"
  public_client            = false
  service_accounts_enabled = true
}

resource "keycloak_client" "pedidos_web_pj" {
  realm_id                 = keycloak_realm.conectcar_realm.id
  client_scope             = keycloak_client_scope.pedidos_client_scope_pj.name
  base_url                 = "http://localhost:8000"
  client_id                = "pedidos_web_pj"
  name                     = "Pedidos Web PJ"
  description              = "Client do formulário de pedidos pessoa juridica"
  protocol                 = "openid-connect"
  public_client            = false
  service_accounts_enabled = true
}
