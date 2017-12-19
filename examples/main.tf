provider "kong" {
  endpoint = "localhost:8001"
}

resource "kong_api" "terraform" {
  name         = "terraform2"
  uris         = ["/terraform"]
  upstream_url = "http://httpbin.org"
}
