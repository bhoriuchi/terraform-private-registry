# terraform-private-registry
Private Terraform Registry

## .terraformrc override
```
host "registry.terraform.io" {
  services = {
    "modules.v1" = "http://localhost:3000/v1/modules/"
    "providers.v1" = "http://localhost:3000/v1/providers/"
  }
}
```