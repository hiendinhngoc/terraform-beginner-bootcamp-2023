package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
    Schema: map[string]*schema.Schema {
      "name": {
        Type: schema.TypeString,
        Required: true,
        Description: "Name of home",
      },
      "description": {
        Type: schema.TypeString,
        Required: true,
        Description: "Description of home",
      },
      "domain_name": {
        Type: schema.TypeString,
        Required: true,
        Description: "Domain name of home",
      },
      "town": {
        Type: schema.TypeString,
        Required: true,
        Description: "The town to which the home will belong to",
      },
      "content_version": {
        Type: schema.TypeInt,
        Required: true,
        Description: "The content version",
      },
    },
  }
  log.Print("Resource:end")

  return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  log.Print("resourceHouseCreate:start")
  var diags diag.Diagnostics

  config := m.(*Config)

  payload := map[string]interface{} {
    "name": d.Get("name").(string),
    "description": d.Get("description").(string),
    "domain_name": d.Get("domain_name").(string),
    "town": d.Get("town").(string),
    "content_version": d.Get("content_version").(int),
  }

  payloadBytes, err := json.Marshal(payload)
  if err != nil {
    return diag.FromErr(err)
  }

  // Construct the HTTP Reauest
  req, err := http.NewRequest("POST", config.Endpoint+"/u/"+config.UserUuid+"/homes", bytes.NewBuffer(payloadBytes))
  if err != nil {
    return diag.FromErr(err)
  }

  // set headers
  req.Header.Set("Authorization", "Bearer "+config.Token)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Accept", "application/json")

  client := http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return diag.FromErr(err)
  }
  defer resp.Body.Close()

  // parse response JSON
  var responseData map[string]interface{}

  if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
    return diag.FromErr(err)
  }

  // StatusOk = 200 HTTP Response Code
  if resp.StatusCode != http.StatusOK {
    return diag.FromErr(fmt.Errorf("failed to create home resource, status_code: %d, status %s, body %s", resp.StatusCode, resp.Status, responseData))
  }

  
  homeUUID := responseData["uuid"].(string)
  d.SetId(homeUUID)
  
  log.Print("resourceHouseCreate:end")
  return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  log.Print("resourceHouseRead:start")
  var diags diag.Diagnostics
  
  config := m.(*Config)
  homeUUID := d.Id()

  // Construct the HTTP Reauest
  req, err := http.NewRequest("GET", config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID, nil)
  if err != nil {
    return diag.FromErr(err)
  }

  // set headers
  req.Header.Set("Authorization", "Bearer "+config.Token)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Accept", "application/json")

  client := http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return diag.FromErr(err)
  }
  defer resp.Body.Close()

  var responseData map[string]interface{}
  if resp.StatusCode == http.StatusOK {
    // parse response JSON
    if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
      return diag.FromErr(err)
    }
    d.Set("name", responseData["name"].(string))
    d.Set("description", responseData["description"].(string))
    d.Set("domain_name", responseData["domain_name"].(string))
    d.Set("content_version", responseData["content_version"].(float64))
  } else if resp.StatusCode == http.StatusNotFound{
    d.SetId("")
  } else if resp.StatusCode != http.StatusOK {
    // StatusOk = 200 HTTP Response Code
    return diag.FromErr(fmt.Errorf("failed to read home resource, status_code: %d, status %s, body %s", resp.StatusCode, resp.Status, responseData))
  }

  log.Print("resourceHouseRead:end")
  return diags
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  log.Print("resourceHouseUpdate:start")
  var diags diag.Diagnostics
  
  config := m.(*Config)
  homeUUID := d.Id()

  payload := map[string]interface{} {
    "name": d.Get("name").(string),
    "description": d.Get("description").(string),
    "content_version": d.Get("content_version").(int),
  }

  payloadBytes, err := json.Marshal(payload)
  if err != nil {
    return diag.FromErr(err)
  }

  // Construct the HTTP Reauest
  req, err := http.NewRequest("PUT", config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID, bytes.NewBuffer(payloadBytes))
  if err != nil {
    return diag.FromErr(err)
  }

  // set headers
  req.Header.Set("Authorization", "Bearer "+config.Token)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Accept", "application/json")

  client := http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return diag.FromErr(err)
  }
  defer resp.Body.Close()

  // parse response JSON
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData);  err != nil {
		return diag.FromErr(err)
	}
  // StatusOk = 200 HTTP Response Code
  if resp.StatusCode != http.StatusOK {
    return diag.FromErr(fmt.Errorf("failed to update home resource, status_code: %d, status %s, body %s", resp.StatusCode, resp.Status, responseData))
  }

  log.Print("resourceHouseUpdate:end")
  
  d.Set("name", payload["name"])
  d.Set("description", payload["description"])
  d.Set("content_version", payload["content_version"])
  return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  log.Print("resourceHouseDelete:start")
  var diags diag.Diagnostics
  
  config := m.(*Config)
  homeUUID := d.Id()

  // Construct the HTTP Reauest
  req, err := http.NewRequest("DELETE", config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID, nil)
  if err != nil {
    return diag.FromErr(err)
  }

  // set headers
  req.Header.Set("Authorization", "Bearer "+config.Token)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Accept", "application/json")

  client := http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return diag.FromErr(err)
  }
  defer resp.Body.Close()

  // StatusOk = 200 HTTP Response Code
  if resp.StatusCode != http.StatusOK {
    return diag.FromErr(fmt.Errorf("failed to delete home resource, status_code: %d, status %s", resp.StatusCode, resp.Status))
  }

  d.SetId("")
  
  log.Print("resourceHouseDelete:end")
  return diags
}