bind_addr = "0.0.0.0"

client {
    host_volume "certs" {
        path      = "/opt/nomad/data/certs"
        read_only = true
    }

    host_volume "roach-certs" {
        path      = "/opt/nomad/data/roach-certs"
        read_only = true
    }

    host_volume "roach-1" {
        path      = "/opt/nomad/data/roach-1"
        read_only = false
    }

    host_volume "roach-2" {
        path      = "/opt/nomad/data/roach-2"
        read_only = false
    }

    host_volume "roach-3" {
        path      = "/opt/nomad/data/roach-3"
        read_only = false
    }
}
