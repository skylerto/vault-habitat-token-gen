# Vault Habitat Token Generation

This is a vault plugin backend used for dynamic generation of habitat authentication tokens.


## Developer Notes

* Plugin Backend Interface (https://github.com/hashicorp/vault/blob/master/logical/framework/backend.go)
* SSH Plugin Implementation (https://github.com/hashicorp/vault/blob/master/builtin/logical/ssh/backend.go)
* Various Backend Types (https://github.com/hashicorp/vault/blob/df18871704fe869e9be45b542a6b1eb2fe46c293/logical/logical.go#L14)
* Generate random SSL
```
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout example.key -out example.crt -extensions san -config <(echo '[req]'; echo 'distinguished_name=req'; echo '[san]'; echo 'subjectAltName=DNS:example.com,DNS:example.net') -subj '/CN=example.com'
```
