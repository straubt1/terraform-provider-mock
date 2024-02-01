# Terraform Interface Design

provider "mock" {
  configure {
    failure         = true                                    # Default false
    failure_message = "I am a provider configuration failure" # Default nil
    failure_type    = "Forced"                                # Default nil
  }
}


resource "mock_resource" "create_failure" {
  id = "mr-12345" # auto generated if not provided, saved to state

  # Settings are saved to state, can this be a dynamic object without `jsonencode`?
  settings = {
    set_number = 1
    set_string = "a"
    set_bool   = false
    # Insert all types (list, map, tuple)
  }

  planmodify {
    failure         = true                                       # Default false
    failure_message = "I am a resoure plan modification failure" # Default nil
    failure_type    = "Forced"                                   # Default nil
    delay           = 5                                          # Default to 0, seconds to delay before returning
  }

  create {
    failure         = true                            # Default false
    failure_message = "I am a resoure create failure" # Default nil
    failure_type    = "Forced"                        # Default nil
    delay           = 5                               # Default to 0, seconds to delay before returning
  }

  update {
    failure         = true                           # Default false
    failure_message = "I am a resoure upate failure" # Default nil
    failure_type    = "Forced"                       # Default nil
    delay           = 5                              # Default to 0, seconds to delay before returning
  }

  read {
    failure         = true                          # Default false
    failure_message = "I am a resoure read failure" # Default nil
    failure_type    = "Forced"                      # Default nil
    save_state      = false                         # Save to state if there is an error
    delay           = 5                             # Default to 0, seconds to delay before returning

    # if set, the value to return during a read operation, gives ability to simulate bad read
    override_id       = "mr-12345"
    override_settings = {}
  }
  delete {
    failure         = true                            # Default false
    failure_message = "I am a resoure delete failure" # Default nil
    failure_type    = "Forced"                        # Default nil
    delay           = 5                               # Default to 0, seconds to delay before returning
  }
}
