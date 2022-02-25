# infrastructure
Setup for AWS infrastructure using IaC


### Prerequisites:
1) aws-cli 2.*
2) configured IAM credentials


### How to run:
1) Clone the repository
2) Run the command <code>export AWS_PROFILE={profile-name}</code> in the terminal
3) Run the command <code>export AWS_REGION={region}</code> in the terminal
4) To use default values run the command <code>aws cloudformation create-stack --stack-name {stack-name} --template-body file://csye6225-infra.yml</code>
5) To use your parameters run the command <code>aws cloudformation create-stack --stack-name {stack-name} --template-body file://csye6225-infra.yml --parameters ParameterKey=VPCCidrBlock,ParameterValue="{vpc-cidr-value}" ParameterKey=Subnet1CidrBlock,ParameterValue="{subnet1-cidr-value}" ParameterKey=Subnet2CidrBlock,ParameterValue="{subnet2-cidr-value}" ParameterKey=Subnet3CidrBlock,ParameterValue="{subnet3-cidr-value}" ParameterKey=AvailabilityZone1,ParameterValue="{availability-zone1-value}" ParameterKey=AvailabilityZone2,ParameterValue="{availability-zone2-value}" ParameterKey=AvailabilityZone3,ParameterValue="{availability-zone3-value}" ParameterKey=PublicRouteCidrBlock,ParameterValue="{public-cidr-value}"</code>
6) To clean up and delete stack run <code>aws cloudformation delete-stack --stack-name {stack-name}</code>
