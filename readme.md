# Opalstack Terraform Provider

This is a Terraform provider for the [Opalstack hosting](https://www.opalstack.com/) api.

## Swagger Codegen

The API client itself is built using Swagger Codegen from the [OpenAPI documentation](https://my.opalstack.com/api/v1/doc/).

To regenerate the client run the following:

```bash
rm -rf swagger/
mkdir swagger/
cd swagger/
swagger-codegen generate -i https://my.opalstack.com/api/v1/schema/ -l go
```

As of version 3.0.30, `swagger/model_ip_address_type_enum.go` generates invalid Golang enums due to starting with an integer. You need to modify that by prepending a `v`. I [reported this upstream](https://github.com/swagger-api/swagger-codegen/issues/11615)

Also, JSON object types need to be converted to either `interface{}` or some other format depending on the documentation. I used `string` in some cases that just looked like key/value pairs.
