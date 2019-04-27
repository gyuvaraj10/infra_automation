package test

import (
	"fmt"
	"testing"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformTest(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("terratest-%s", random.UniqueId())
	awsRegion := "eu-west-2"

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../terraform",
		// // Variables to pass to our Terraform code using -var options
		// Vars: map[string]interface{}{
		// 	"instance_name": expectedName,
		// },
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	aws.AddTagsToResource(t, awsRegion, instanceID, map[string]string{"testing": "testing-tag-value"})

	// Look up the tags for the given Instance ID
	instanceTags := aws.GetTagsForEc2Instance(t, awsRegion, instanceID)
	aws.
	testingTag, containsTestingTag := instanceTags["testing"]
	assert.True(t, containsTestingTag)
	assert.Equal(t, "testing-tag-value", testingTag)

	// Verify that our expected name tag is one of the tags
	nameTag, containsNameTag := instanceTags["Name"]
	assert.True(t, containsNameTag)
	assert.Equal(t, expectedName, nameTag)

}
