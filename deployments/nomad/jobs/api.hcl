job "api" {
  type = "service"

  group "api" {
    count = 1

    network {
      mode = "host"

      port "api" {
        to = 6000
      }
    }

    service {
      name     = "api"
      port     = "api"
      provider = "nomad"
    }

    volume "roach-certs" {
        type      = "host"
        read_only = true
        source    = "roach-certs"
    }

    task "api" {
      driver = "docker"

        volume_mount {
            volume      = "roach-certs"
            destination = "/certs"
            read_only   = true
        }

      config {
        image = "registry.com.gecko/api"
        ports = ["api"]
      }

      env {
        ROACH_DATABASE = "go_demo"
        ROACH_HOST     = "${NOMAD_ADDR_roach}"
        ROACH_PASSWORD = "go_password"
        ROACH_PORT     = "${NOMAD_HOST_PORT_roach}"
        ROACH_USER     = "go_user"

        ROACH_SSL_CERTIFICATE = "/certs/client.go_user.crt"
        ROACH_SSL_KEY         = "/certs/client.go_user.key"
        ROACH_SSL_CA          = "/certs/ca.crt"
      }
    }
  }
}
