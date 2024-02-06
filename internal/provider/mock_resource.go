// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MockResource{}

// TODO: add import and modify plan options
var _ resource.ResourceWithImportState = &MockResource{}
var _ resource.ResourceWithModifyPlan = &MockResource{}

func NewMockResource() resource.Resource {
	return &MockResource{}
}

// MockResource defines the resource implementation.
type MockResource struct {
	client *http.Client
}

// MockResourceModel describes the resource data model.
type MockResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Create     CRUDModel    `tfsdk:"create"`
	Read       CRUDModel    `tfsdk:"read"`
	Update     CRUDModel    `tfsdk:"update"`
	Delete     CRUDModel    `tfsdk:"delete"`
	PlanModify CRUDModel    `tfsdk:"planmodify"`
}

type CRUDModel struct {
	Failure        types.Bool   `tfsdk:"failure"`
	FailureMessage types.String `tfsdk:"failure_message"`
	FailureType    types.String `tfsdk:"failure_type"`
	Delay          types.Int64  `tfsdk:"delay"`
}

func (r *MockResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (r *MockResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Mock resource",

		Blocks: map[string]schema.Block{
			"create": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"failure": schema.BoolAttribute{
						Optional: true,
						// Default:  booldefault.StaticBool(false),
					},
					"failure_message": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"failure_type": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"delay": schema.Int64Attribute{
						Optional: true,
						// Default:  int64default.StaticInt64(0),
					},
				},
			},
			"read": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"failure": schema.BoolAttribute{
						Optional: true,
						// Default:  booldefault.StaticBool(false),
					},
					"failure_message": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"failure_type": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"delay": schema.Int64Attribute{
						Optional: true,
						// Default:  int64default.StaticInt64(0),
					},
				},
			},
			"update": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"failure": schema.BoolAttribute{
						Optional: true,
						// Default:  booldefault.StaticBool(false),
					},
					"failure_message": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"failure_type": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"delay": schema.Int64Attribute{
						Optional: true,
						// Default:  int64default.StaticInt64(0),
					},
				},
			},
			"delete": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"failure": schema.BoolAttribute{
						Optional: true,
						// Default:  booldefault.StaticBool(false),
					},
					"failure_message": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"failure_type": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"delay": schema.Int64Attribute{
						Optional: true,
						// Default:  int64default.StaticInt64(0),
					},
				},
			},
			"planmodify": schema.SingleNestedBlock{
				Attributes: map[string]schema.Attribute{
					"failure": schema.BoolAttribute{
						Optional: true,
						// Default:  booldefault.StaticBool(false),
					},
					"failure_message": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"failure_type": schema.StringAttribute{
						Optional: true,
						// Default:  stringdefault.StaticString(""),
					},
					"delay": schema.Int64Attribute{
						Optional: true,
						// Default:  int64default.StaticInt64(0),
					},
				},
			},
		},
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Mock identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *MockResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*http.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *MockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MockResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	tflog.Trace(ctx, "createtrace: "+plan.Create.FailureMessage.String())

	if plan.Create.Failure.ValueBool() {
		resp.Diagnostics.AddError(
			plan.Create.FailureMessage.String(),
			plan.Create.FailureType.String(),
		)
	}

	if plan.Id.IsUnknown() {
		plan.Id = types.StringValue("set by provider")
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// delaySeconds := 3

	time.Sleep(time.Duration(plan.Create.Delay.ValueInt64()) * time.Second)
	// call error at the end after delay
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MockResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	tflog.Trace(ctx, "readtrace: "+state.Read.FailureMessage.String())

	if state.Read.Failure.ValueBool() {
		resp.Diagnostics.AddError(
			state.Read.FailureMessage.String(),
			state.Read.FailureType.String(),
		)
	}

	if state.Id.IsUnknown() {
		state.Id = types.StringValue("set by provider")
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a resource")

	// delaySeconds := 3

	time.Sleep(time.Duration(state.Read.Delay.ValueInt64()) * time.Second)
	// call error at the end after delay
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MockResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	tflog.Trace(ctx, "updatetrace: "+plan.Update.FailureMessage.String())

	if plan.Update.Failure.ValueBool() {
		resp.Diagnostics.AddError(
			plan.Update.FailureMessage.String(),
			plan.Update.FailureType.String(),
		)
	}

	// if plan.Id.IsUnknown() {
	// 	plan.Id = types.StringValue("set by provider")
	// }

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "updated a resource")

	// delaySeconds := 3

	time.Sleep(time.Duration(plan.Update.Delay.ValueInt64()) * time.Second)
	// call error at the end after delay
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MockResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	tflog.Trace(ctx, "deletetrace: "+state.Delete.FailureMessage.String())

	if state.Delete.Failure.ValueBool() {
		resp.Diagnostics.AddError(
			state.Delete.FailureMessage.String(),
			state.Delete.FailureType.String(),
		)
	}

	// if state.Id.IsUnknown() {
	// 	state.Id = types.StringValue("set by provider")
	// }

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "delete a resource")

	// delaySeconds := 3

	time.Sleep(time.Duration(state.Delete.Delay.ValueInt64()) * time.Second)
	// call error at the end after delay
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r MockResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan MockResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	tflog.Trace(ctx, "planmodifytrace: "+plan.PlanModify.FailureMessage.String())

	if plan.PlanModify.Failure.ValueBool() {
		resp.Diagnostics.AddError(
			plan.PlanModify.FailureMessage.String(),
			plan.PlanModify.FailureType.String(),
		)
	}

	// if plan.Id.IsUnknown() {
	// 	plan.Id = types.StringValue("set by provider")
	// }

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "plan modify a resource")

	// delaySeconds := 3

	time.Sleep(time.Duration(plan.PlanModify.Delay.ValueInt64()) * time.Second)
	// call error at the end after delay
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Plan
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (r *MockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
