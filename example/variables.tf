variable "url" {
  description = "BaseURL do Keycloak"
  default     = "https://qa-aks-conectcar-keycloak.conectcar.com"
}

variable "username" {
  description = "Usuário administrador do Keycloak"
  default     = "admin"
}

variable "password" {
  description = "Senha do usuário administrador do Keycloak"
  default     = "@C0nectC@r@2021"
}
