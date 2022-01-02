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
