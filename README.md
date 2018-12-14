# Vault Habitat Token Generation

This is a vault plugin backend used for dynamic generation of habitat authentication tokens.

The purpose is to provide a way to authenticate with vault, then retrieve and rotate habitat depot
auth tokens for both public and on-prem.


## Features

The following is the first release feature set:

* Login to any Depot (requires the first auth token)
* Get Auth Token
* Refresh Auth Token

Each of these features should be implemented for any (user account, depot) pair.
Which shall be governed by the paths that you write to within the plugin.

## Examples

```
vault write habitattoken/<unique-id>/login auth_token=${current} hab_bldr_url=https://bldr.habitat.sh
vault write habitattoken/<unique-id>/get
vault write habitattoken/<unique-id>/renew
```

In this case we've enabled the plugin at the path `habitattoken`.

## Developer Notes

* Plugin Backend Interface (https://github.com/hashicorp/vault/blob/master/logical/framework/backend.go)
* SSH Plugin Implementation (https://github.com/hashicorp/vault/blob/master/builtin/logical/ssh/backend.go)
* Various Backend Types (https://github.com/hashicorp/vault/blob/df18871704fe869e9be45b542a6b1eb2fe46c293/logical/logical.go#L14)
* Generate random SSL
```
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout example.key -out example.crt -extensions san -config <(echo '[req]'; echo 'distinguished_name=req'; echo '[san]'; echo 'subjectAltName=DNS:example.com,DNS:example.net') -subj '/CN=example.com'
```
