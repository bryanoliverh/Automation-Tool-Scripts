package main

import (
    "context"
    "fmt"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/plugin"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
    "google.golang.org/api/compute/v1"
    "google.golang.org/api/option"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
                    "gcp_virtual_machine": resourceVirtualMachine(),
                },
            }
        },
    })
}

func resourceVirtualMachine() *schema.Resource {
    return &schema.Resource{
        Create: resourceVirtualMachineCreate,
        Read:   resourceVirtualMachineRead,
        Update: resourceVirtualMachineUpdate,
        Delete: resourceVirtualMachineDelete,

        Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "image": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "machine_type": &schema.Schema{
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

func resourceVirtualMachineCreate(d *schema.ResourceData, m interface{}) error {
    ctx := context.Background()

    computeService, err := compute.NewService(ctx, option.WithCredentialsFile("/path/to/credentials.json"))
    if err != nil {
        return fmt.Errorf("Error creating compute service: %s", err)
    }

    // Get the required parameters from the Terraform configuration
    name := d.Get("name").(string)
    image := d.Get("image").(string)
    machineType := d.Get("machine_type").(string)
    zone := d.Get("zone").(string)

    // Create the instance configuration object
    instance := &compute.Instance{
        Name:        name,
        MachineType: fmt.Sprintf("zones/%s/machineTypes/%s", zone, machineType),
        Disks: []*compute.AttachedDisk{
            &compute.AttachedDisk{
                AutoDelete: true,
                Boot:       true,
                Type:       "PERSISTENT",
                InitializeParams: &compute.AttachedDiskInitializeParams{
                    SourceImage: image,
                },
            },
        },
        NetworkInterfaces: []*compute.NetworkInterface{
            &compute.NetworkInterface{
                AccessConfigs: []*compute.AccessConfig{
                    &compute.AccessConfig{
                        Type: "ONE_TO_ONE_NAT",
                    },
                },
                Network: "global/networks/default",
            },
        },
    }

    // Insert the instance into GCP
    operation, err := computeService.Instances.Insert("<YOUR-PROJECT-ID>", zone, instance).Do()
    if err != nil {
        return fmt.Errorf("Error creating instance: %s", err)
    }

    // Store the instance ID in the Terraform state
    d.SetId(operation.TargetId)

    return resourceVirtualMachineRead(d, m)
}

func resourceVirtualMachineRead(d *schema.ResourceData, m interface{}) error {
    ctx := context.Background()

    // Get the GCP provider configuration
    providerConf := m.(*GCPConfig)

    // Create the GCP client using the provider configuration
    computeService, err := compute.NewService(ctx, option.WithCredentialsFile(providerConf.CredentialsFile))
    if err != nil {
        return fmt.Errorf("Error creating Compute service: %v", err)
    }

    // Get the virtual machine ID from the resource data
    vmID := d.Id()

    // Call the Compute API to get the virtual machine
    vm, err := computeService.Instances.Get(providerConf.ProjectID, providerConf.Zone, vmID).Do()
    if err != nil {
        return fmt.Errorf("Error getting virtual machine: %v", err)
    }

    // Update the resource data with the virtual machine properties
    d.Set("name", vm.Name)
    d.Set("machine_type", vm.MachineType)
    d.Set("zone", vm.Zone)
    d.Set("description", vm.Description)
    d.Set("network_interface", flattenNetworkInterface(vm.NetworkInterfaces))

    return nil
}

func flattenNetworkInterface(nis []*compute.NetworkInterface) []interface{} {
    var result []interface{}

    for _, ni := range nis {
        niMap := make(map[string]interface{})

        niMap["network"] = ni.Network
        niMap["subnetwork"] = ni.Subnetwork
        niMap["access_config"] = flattenAccessConfig(ni.AccessConfigs)

        result = append(result, niMap)
    }

    return result
}

func flattenAccessConfig(ac []*compute.AccessConfig) []interface{} {
    var result []interface{}

    for _, a := range ac {
        aMap := make(map[string]interface{})

        aMap["name"] = a.Name
        aMap["nat_ip"] = a.NatIP

        result = append(result, aMap)
    }

    return result
}