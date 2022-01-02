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

I'm generally trying to add a comment anywhere I modify things from the auto-generated defaults, to make updates easier. I start the comment with `SWAGGERMOD:` for easy searching. I do my best to leave things alone, but sometimes there isn't an option.

## API Mapping Notes

I tried to keep as close as possible to the terminology and semantics of the API.

I skipped the app "failed" and "installed" endpoints. They seem to be useful for running a script which installs application code after the app itself is installed. I can't think of any way this makes sense with Terraform. If you think of one, let me know.

Resource, mail, and schema endpoints aren't implemented. The former two only return access denied errors for me, and the latter one doesn't make sense in this context.

I'm skipping quarantinedmail because I'm not sure how to validate that I did it right, and I don't see a use with Terraform. File a ticket if you have a use case for this.

Mailuser has an update_public api endpoint, for which I don't know the use. I'll look at implementing it when I understand its purpose.

## Development

Not sure if this is the best way to do this, but here's how I go about it.

Make a folder for Terraform to cache the locally built version of your plugin. Modify `darwin_amd64` to reflect the OS you are using. This is `GOOS` and `GOARCH` from the [list in the Golang documentation](https://go.dev/doc/install/source#environment). You can also get it from the output of `terraform -version`.
```bash
# Example for macOS
mkdir -p ~/.terraform.d/plugins/github.com/will-in-wi/opalstack/0.0.1/darwin_amd64/
```

Vendor the Golang modules. You may want to do this whenever you add a dependency.
```bash
go mod vendor
```

Whenever you want to test a change, compile the module and copy the binary to the folder you created above.
```bash
go build -o terraform-provider-opalstack && mv terraform-provider-opalstack ~/.terraform.d/plugins/github.com/will-in-wi/opalstack/0.0.1/darwin_amd64/
```

In another folder, create a Terraform file (I call it `main.tf`), and stick the following in it:
```hcl
terraform {
  required_providers {
    opalstack = {
      version = "~> 0.0.1"
      source  = "github.com/will-in-wi/opalstack"
    }
  }
}

provider "opalstack" {
  # Generate a token in the Opalstack admin and add it here. Or prepend all your Terraform commands with OPALSTACK_TOKEN=tokenfromopalstack
  token = "tokenfromopalstack"
}
```

In that folder, you can run `rm .terraform.lock.hcl && terraform init && terraform plan` to pick up the newly compiled version and run a plan. Technically the `rm` and `terraform init` are only required when the provider binary changes.
