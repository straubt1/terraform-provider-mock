terraform {
  required_providers {
    mock = {
      source = "hashicorp.com/edu/mock"
    }
  }
}

provider "mock" {}

resource "mock_resource" "example" {
  # id = "tom"

  create {
    # failure         = true                            # Default false
    failure_message = "I am a resoure create failure" # Default nil
    failure_type    = "Forced"                        # Default nil
    # delay           = 3                               # Default to 0, seconds to delay before returning
  }
  read {
    # failure         = true
    failure_message = "I am a resource read failure"
    failure_type    = "Forced from Terraform *.tf"
    # delay           = 2
  }
  update {
    # failure         = true
    failure_message = "I am a resource update failure"
    failure_type    = "Forced from Terraform *.tf"
    # delay           = 2
  }
  delete {
    # failure         = true
    failure_message = "I am a resource delete failure"
    failure_type    = "Forced from Terraform *.tf"
    # delay           = 2
  }
  planmodify {
    failure         = true
    failure_message = "I am a resource plan modify failure"
    failure_type    = "Forced from Terraform *.tf"
    delay           = 2
  }
}
