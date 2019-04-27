Pre-conditions:
    keep the service account key with the compute engine permissions. example: laranerds-157820-f31e50997708.json in the packer directory. The same applied for terraform also.

Packer:
 To build an image in GCP with the packer run the following steps
    1. cd packer
    2. packer build build_centos.json

Terraform:
 To spin up an instance from the built packer image above, run the following steps.
    1. cd terraform
    2. terraform init (if you have not initialized the terraform)
    3. terraform validate (if you want to make sure that there are no errors in the terrafom template)
    4. terraform plan -out learnterraform (Now a binary file will be generated)
    5. terraform apply learnterraform (this will create an instance with the added configuration in the terraform template)



