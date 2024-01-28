terraform {
  required_providers {
    mock = {
      source = "hashicorp.com/edu/mock"
    }
  }
}

provider "mock" {}

# data "hashicups_coffees" "example" {}


resource "mock_example" "example" {
  configurable_attribute = "some-value"
  create_failure         = true
}
