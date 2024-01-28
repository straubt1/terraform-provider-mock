# Terraform Interface Design


resource "mock_resource" "create_failure" {

  create {
    failure         = true
    failure_message = "I am a create failure"
    failure_type    = "Forced"
  }

  update {
    failure         = true
    failure_message = "I am a create failure"
    failure_type    = "Forced"
  }

  read {
    failure         = true
    failure_message = "I am a create failure"
    failure_type    = "Forced"
  }
  delete {
    failure         = true
    failure_message = "I am a create failure"
    failure_type    = "Forced"
  }

}
