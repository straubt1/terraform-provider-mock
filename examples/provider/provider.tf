terraform {
  required_providers {
    mock = {
      source = "hashicorp.com/edu/mock"
    }
  }
}

provider "mock" {}

# data "hashicups_coffees" "example" {}


resource "mock_resource" "example" {
  configurable_attribute = "some-value"
  # create_failure         = true

  # create = {
  #   failure         = true                            # Default false
  #   failure_message = "I am a resoure create failure" # Default nil
  #   failure_type    = "Forced"                        # Default nil
  #   delay           = 5                               # Default to 0, seconds to delay before returning
  # }
}
