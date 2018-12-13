storage "file" {
  path = "/vault"
  redirect_addr = "https://127.0.0.1:8200"
}

listener "tcp" {
  address = "127.0.0.1:8200"
  tls_disable = 0
  tls_cert_file = "example.crt"
  tls_key_file = "example.key"
}

log_level = "debug"

plugin_directory = "/src"

disable_mlock = true

api_addr = "https://127.0.0.1:8200"
