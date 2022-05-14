build:
	go build -o terraform-provider-keycloak

build-example: build
	mkdir -p ~/.terraform.d/plugins/terraform-example.com/jailtonjunior/keycloak/1.0.0/linux_amd64
	cp  terraform-provider-keycloak ~/.terraform.d/plugins/terraform-example.com/jailtonjunior/keycloak/1.0.0/linux_amd64