package main

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/plugin"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
    "google.golang.org/api/option"
    "google.golang.org/api/compute/v1"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
                    "gcp_instance": resourceInstance(),
                },
            }
        },
    })
}

func resourceInstance() *schema.Resource {
    return &schema.Resource{
        Create: resourceInstanceCreate,
        Read:   resourceInstanceRead,
        Delete: resourceInstanceDelete,

        Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "zone": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceInstanceCreate(d *schema.ResourceData, m interface{}) error {
    ctx := context.Background()
    computeService, err := compute.NewService(ctx, option.WithCredentialsFile("/path/to/credentials.json"))
    if err != nil {
        return fmt.Errorf("Error creating compute service: %s", err)
    }

    name := d.Get("name").(string)
    zone := d.Get("zone").(string)

    instance := &compute.Instance{
        Name: name,
        Zone: "projects/<PROJECT>/zones/" + zone,
        MachineType: "projects/<PROJECT>/zones/" + zone + "/machineTypes/n1-standard-1",
        Disks: []*compute.AttachedDisk{
            {
                AutoDelete: true,
                Boot:       true,
                Type:       "PERSISTENT",
                InitializeParams: &compute.AttachedDiskInitializeParams{
                    DiskName:    name + "-disk",
                    SourceImage: "projects/debian-cloud/global/images/family/debian-10",
                },
            },
        },
        NetworkInterfaces: []*compute.NetworkInterface{
            {
                Network: "global/networks/default",
                AccessConfigs: []*compute.AccessConfig{
                    {
                        Type: "ONE_TO_ONE_NAT",
                        Name: "External NAT",
                    },
                },
            },
        },
    }

    operation, err := computeService.Instances.Insert("<PROJECT>", zone, instance).Do()
    if err != nil {
        return fmt.Errorf("Error creating instance: %s", err)
    }

    d.SetId(name)

    return nil
}

func resourceInstanceRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceInstanceDelete(d *schema.ResourceData, m interface{}) error {
    ctx := context.Background()
    computeService, err := compute.NewService(ctx, option.WithCredentialsFile("/path/to/credentials.json"))
    if err != nil {
        return fmt.Errorf("Error creating compute service: %s", err)
    }

    name := d.Id()
    zone := d.Get("zone").(string)

    operation, err := computeService.Instances.Delete("<PROJECT>", zone, name).Do()
    if err != nil {
        return fmt.Errorf("Error deleting instance: %s", err)
    }

    return nil
}
