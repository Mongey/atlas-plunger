## Atlas Plunger
### Unclog your AMI deployment pipeline ™
#### Artifact promotion for your [Atlas](https://atlas.hashicorp.com) artifacts
![http://i.imgur.com/YOU57nK.jpg](http://i.imgur.com/YOU57nK.jpg)
[![Build Status](https://travis-ci.org/Mongey/atlas-plunger.svg?branch=master)](https://travis-ci.org/Mongey/atlas-plunger)

### Motivation
It can be difficult to track what versions of AMIS are being put into different
environments. Atlas-Plunger aims to help create a more seamless pipeline for
promoting artifacts through your infrastructures environment.

AMI ID's and Atlas build version numbers are not useful for tracking what
version of code you are deploying.

## Goals
* Increase visibility for code is running in your environments.
* Decrease churn on terraform files for bumping AMI version

### Example
What environments are available
`atlas list environments`
```
ci
chaos
uat
staging
production
```

What is running in [Env]
`atlas find uat kafka`
```
sha     : e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855he
ami_id  : some_id
region  : some_id 
```

The packer postprocessor is intended to help link commits to ami
Packer PostProcessor

`packer/kafka.json`
```json
{
  "post-processors": [
    [
      {
        "type": "atlas",
        "artifact": "mongey/kafka",
        "artifact_type": "amazon.ami",
        "metadata": {
          "created_at": "{{timestamp}}",
          "git_commit": "{{user `commit`}}",
          "git_status": "{{user `git_status`}}",
          "branch": "{{ user `branch` }}",
        }
      }
    ]
  ]
}

```
Your CI should then create an artifact, promoting it to the first stage like so

```
atlas promote kafka chaos $git_commit
```


Terraform Example
```
terraform
└── providers
    └── aws
        ├── production-us-west-1
        │   └── kafka.tf
        └── staging-us-west-1
            └── kafka.tf
```

`staging-us-west-1`
```hcl
resource "atlas_artifact" "kafka" {
  name     = "mongey/kafka"
  type     = "amazon.ami"
  version  = "latest"
  metadata = {
    environment = "staging"
  }
}

```
`production-us-west-1`
```hcl
resource "atlas_artifact" "kafka" {
  name     = "mongey/kafka"
  type     = "amazon.ami"
  version  = "latest"
  metadata = {
    environment = "production"
  }
}
```


