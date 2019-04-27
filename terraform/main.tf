provider "google" { 
    credentials = "${file("../packer/laranerds-157820-f31e50997708.json")}"
    zone = "europe-west2-c"
}
data "google_compute_image" "packerlearn" {
  project = "laranerds-157820"
  name = "packerlearning1"
}

resource "google_compute_address" "external_ip" {
  project = "laranerds-157820"
  name = "externalip"
}

resource "google_compute_network" "default" {
  name = "packer-network"
  project = "laranerds-157820"
}

resource "google_compute_firewall" "default" {
  name= "packer-firewall"
  project = "laranerds-157820"
  network = "${google_compute_network.default.name}"
  allow {
    protocol = "tcp"
    ports = ["80", "22", "443", "8080"]
  }
}


resource "google_compute_instance" "new_instance" {
  name = "packerlearninginsta"
  machine_type = "n1-standard-1"
  network_interface {
    network = "${google_compute_network.default.name}"
    access_config {
      nat_ip = "${google_compute_address.external_ip.address}"
    }
  } 
  project = "laranerds-157820"
  boot_disk {
    initialize_params {
      image = "${data.google_compute_image.packerlearn.self_link}"
    }
    auto_delete = "true"
  }
  metadata {
    ssh-keys = "yuvaraj.gunisetti:${file("~/.ssh/id_rsa.pub")}"
  }
}