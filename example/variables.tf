variable "url" {
  description = "BaseURL do Keycloak"
  default     = "http://localhost:8080"
}

variable "username" {
  description = "Usuário administrador do Keycloak"
  default     = "admin"
}

variable "password" {
  description = "Senha do usuário administrador do Keycloak"
  default     = "admin"
}