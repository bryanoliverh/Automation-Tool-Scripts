package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/plugin"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
    "google.golang.org/api/option"
    "google.golang.org/api/storage/v1"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
                    "my_bucket": resourceBucket(),
                },
            }
        },
    })
}

func resourceBucket() *schema.Resource {
    return &schema.Resource{
        Create: resourceBucketCreate,
        Read:   resourceBucketRead,
        Delete: resourceBucketDelete,

        Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "location": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "project": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceBucketCreate(d *schema.ResourceData, m interface{}) error {
    ctx := context.Background()
    client, err := storage.NewService(ctx, option.WithCredentialsFile("/path/to/credentials.json"))
    if err != nil {
        return err
    }

    name := d.Get("name").(string)
    location := d.Get("location").(string)
    project := d.Get("project").(string)

    bucket := &storage.Bucket{
        Name:     name,
        Location: location,
    }

    _, err = client.Buckets.Insert(project, bucket).Do()
    if err != nil {
        return err
    }

    d.SetId(name)

    return nil
}

func resourceBucketRead(d *schema.ResourceData, m interface{}) error {
    ctx := context.Background()
    client, err := storage.NewService(ctx, option.WithCredentialsFile("/path/to/credentials.json"))
    if err != nil {
        return err
    }

    name := d.Id()

    bucket, err := client.Buckets.Get(name).Do()
    if err != nil {
        return err
    }

    d.Set("location", bucket.Location)
    d.Set("project", bucket.ProjectNumber)

    return nil
}

func resourceBucketDelete(d *schema.ResourceData, m interface{}) error {
    ctx := context.Background()
    client, err := storage.NewService(ctx, option.WithCredentialsFile("/path/to/credentials.json"))
    if err != nil {
        return err
    }

    name := d.Id()

    err = client.Buckets.Delete(name).Do()
    if err != nil {
        return err
    }

    return nil
}
