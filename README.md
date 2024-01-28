# Mock Provider

The purpose of this Terraform Provider is to allow for testing, debugging, and exploration of the Terraform Workflow



## Target Features

The follow features are desired as an output of this Provider.

- [Resource] Ability to force a failure at:
  - Create()
  - Read()
  - Update()
  - Delete()
- Ability to use resource [PlanModification](https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#resource-plan-modification)

```
gh repo create --add-readme terraform --public terraform-provider-mock 
```

## References

* https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework
* https://github.com/hashicorp/terraform-provider-scaffolding-framework
* [Random Integer](https://github.com/hashicorp/terraform-provider-random/blob/main/internal/provider/resource_integer.go)