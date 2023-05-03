package main

import (
    "context"
    "fmt"
    "os"

    "github.com/google/uuid"
    "github.com/hashicorp/terraform-exec/tfexec"
)

func main() {
    // Create a new Terraform executor
    tf, err := tfexec.NewTerraform(".", "")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Initialize Terraform
    err = tf.Init(context.Background(), tfexec.Upgrade(true), tfexec.Lock(false))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Create a unique identifier for the resources
    id := uuid.New().String()

    // Set the Terraform variables for the GCP instance
    gcpVars := map[string]interface{}{
        "instance_name": "instance-" + id,
        "machine_type":  "e2-micro",
        "zone":          "us-central1-a",
    }

    // Apply the Terraform configuration for the GCP instance
    err = tf.Apply(context.Background(), tfexec.VarFiles([]string{"gcp.tfvars"}), tfexec.Vars(gcpVars))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Set the Terraform variables for the SQL instance
    sqlVars := map[string]interface{}{
        "instance_name": "sql-" + id,
        "tier":          "db-n1-standard-1",
        "region":        "us-central1",
    }

    // Apply the Terraform configuration for the SQL instance
    err = tf.Apply(context.Background(), tfexec.VarFiles([]string{"sql.tfvars"}), tfexec.Vars(sqlVars))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Modify the Terraform configuration for the SQL instance
    modifiedSqlVars := map[string]interface{}{
        "tier": "db-n1-standard-2",
    }

    // Apply the modified Terraform configuration for the SQL instance
    err = tf.Apply(context.Background(), tfexec.VarFiles([]string{"sql.tfvars"}), tfexec.Vars(modifiedSqlVars))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Destroy the Terraform resources
    err = tf.Destroy(context.Background(), tfexec.VarFiles([]string{"gcp.tfvars", "sql.tfvars"}))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
