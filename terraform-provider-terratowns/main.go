package main

import (
	"context"
	"fmt"
	"log"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: Provider,
  })

  fmt.Println("Hello, worlrd!")
}

type Config struct {
  Endpoint string
  Token string
  UserUuid string
}

func Provider() *schema.Provider{
  var p *schema.Provider

  p = &schema.Provider{
    ResourcesMap: map[string]*schema.Resource {
      "terratowns_home": Resource(),
    },
    DataSourcesMap: map[string]*schema.Resource {

    },
    Schema: map[string]*schema.Schema {
      "endpoint": {
        Type: schema.TypeString,
        Required: true,
        Description: "The endpoint for the external service",
      },
      "token": {
        Type: schema.TypeString,
        Sensitive: true,
        Required: true,
        Description: "Bearer token for authorization",
      },
      "user_uuid": {
        Type: schema.TypeString,
        Required: true,
        Description: "UUID for configuration",
        ValidateFunc: validateUUID,
      },
    },
  }
  p.ConfigureContextFunc = providerConfigure(p)

  return p
}

func validateUUID(val interface{}, k string) (ws []string, errors []error) {
  log.Print("validateUUID:start")
  value := val.(string)

  if _, err := uuid.Parse(value); err != nil {
    errors = append(errors, fmt.Errorf("invalid UUID format"))
  }

  log.Print("validateUUID:end")

  return
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
  return func(ctx context.Context, d *schema.ResourceData) (interface {}, diag.Diagnostics) {
    log.Print("providerConfigure:start")
    config := Config{
      Endpoint: d.Get("endpoint").(string),
      Token: d.Get("token").(string),
      UserUuid: d.Get("user_uuid").(string),
    }
    log.Print("providerConfigure:end")
    return &config, nil
  }
}

func Resource() *schema.Resource {
  log.Print("Resource:start")
  resource := &schema.Resource{
    CreateContext: resourceHouseCreate,
    ReadContext: resourceHouseRead,
    UpdateContext: resourceHouseUpdate,
    DeleteContext: resourceHouseDelete,
  }
  log.Print("Resource:end")

  return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
  return diags
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
  return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
  return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics
  return diags
}