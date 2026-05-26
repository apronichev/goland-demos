package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// runProviderExample demonstrates the structure of a custom Terraform provider
func runProviderExample() {
	fmt.Println("=== Custom Terraform Provider Example ===")
	fmt.Println("Demonstrating the structure of a Terraform provider written in Go")
	fmt.Println()

	fmt.Println("A Terraform provider consists of:")
	fmt.Println("1. Provider Schema - defines provider-level configuration")
	fmt.Println("2. Resources - defines managed resources (create, read, update, delete)")
	fmt.Println("3. Data Sources - defines read-only data sources")
	fmt.Println()

	// Create an example provider
	provider := exampleProvider()

	fmt.Println("Example Provider Created:")
	fmt.Printf("  Resources: %d defined\n", len(provider.ResourcesMap))
	fmt.Printf("  Data Sources: %d defined\n", len(provider.DataSourcesMap))
	fmt.Println()

	fmt.Println("Provider Configuration Schema:")
	for key, schemaItem := range provider.Schema {
		fmt.Printf("  - %s (%s, required: %t)\n", key, schemaItem.Type, schemaItem.Required)
	}
	fmt.Println()

	fmt.Println("Resources:")
	for name := range provider.ResourcesMap {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()

	fmt.Println("Data Sources:")
	for name := range provider.DataSourcesMap {
		fmt.Printf("  - %s\n", name)
	}
	fmt.Println()

	fmt.Println("=== Provider Example Complete ===")
	fmt.Println("\nTo create a real provider, you would:")
	fmt.Println("1. Implement the provider schema with configuration")
	fmt.Println("2. Define resources with CRUD operations")
	fmt.Println("3. Define data sources with read operations")
	fmt.Println("4. Build and publish the provider binary")
	fmt.Println("5. Register with Terraform Registry (optional)")
}

// exampleProvider returns an example provider schema
func exampleProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API URL for the service",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "API key for authentication",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "Request timeout in seconds",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"example_server":   exampleServerResource(),
			"example_database": exampleDatabaseResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"example_image": exampleImageDataSource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// exampleServerResource defines a managed resource
func exampleServerResource() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a server resource",

		CreateContext: resourceServerCreate,
		ReadContext:   resourceServerRead,
		UpdateContext: resourceServerUpdate,
		DeleteContext: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the server",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the server",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance type (e.g., small, medium, large)",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "us-east-1",
				Description: "The region where the server will be created",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Tags to apply to the server",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP address of the server",
			},
		},
	}
}

// exampleDatabaseResource defines another managed resource
func exampleDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a database resource",

		CreateContext: resourceDatabaseCreate,
		ReadContext:   resourceDatabaseRead,
		UpdateContext: resourceDatabaseUpdate,
		DeleteContext: resourceDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the database",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the database",
			},
			"engine": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Database engine (postgres, mysql, etc.)",
			},
			"size_gb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
				Description: "Size of the database in GB",
			},
		},
	}
}

// exampleImageDataSource defines a data source
func exampleImageDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieves information about an image",

		ReadContext: dataSourceImageRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the image",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the image",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the image",
			},
		},
	}
}

// Provider configuration function
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// In a real provider, you would:
	// 1. Read configuration values
	// 2. Create an API client
	// 3. Validate credentials
	// 4. Return the client for use in resource operations

	fmt.Println("Provider configured (example)")
	return nil, nil
}

// Resource CRUD operations (simplified examples)
func resourceServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// In a real provider:
	// 1. Extract attributes from d.Get()
	// 2. Make API call to create resource
	// 3. Set the resource ID with d.SetId()
	// 4. Set computed attributes with d.Set()
	fmt.Println("Creating server resource (example)")
	return nil
}

func resourceServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Reading server resource (example)")
	return nil
}

func resourceServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Updating server resource (example)")
	return nil
}

func resourceServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Deleting server resource (example)")
	return nil
}

// Database resource operations
func resourceDatabaseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Creating database resource (example)")
	return nil
}

func resourceDatabaseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Reading database resource (example)")
	return nil
}

func resourceDatabaseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Updating database resource (example)")
	return nil
}

func resourceDatabaseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Deleting database resource (example)")
	return nil
}

// Data source read operation
func dataSourceImageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Reading image data source (example)")
	return nil
}
