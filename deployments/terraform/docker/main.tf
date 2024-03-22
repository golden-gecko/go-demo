terraform {
  required_version = "1.7.5"

  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

provider "docker" {
}

resource "docker_image" "node" {
  name = "registry.com.gecko/node"
}

resource "docker_image" "python" {
  name = "registry.com.gecko/python"
}

resource "docker_container" "node" {
  image = docker_image.node.image_id
  name  = "node"
}

resource "docker_container" "python" {
  image = docker_image.python.image_id
  name  = "python"
}
