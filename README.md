Converts a Terraform plan file to JSON. Also provides a Go [library](https://godoc.org/github.com/justinbarrick/go-terraform-plan) for reading Terraform plan files.

We use the actual [Terraform libraries](https://godoc.org/github.com/hashicorp/terraform) to ensure correctness.

# Installation

To install, run:

```
go get github.com/justinbarrick/go-terraform-plan
go install github.com/justinbarrick/go-terraform-plan
```

# Docker

There is also a docker image:

```
$ terraform plan -out=terraform.plan
$ docker run -w "$(pwd)" -v "$(pwd):$(pwd)" justinbarrick/terraform-plan -plan ./terraform.plan
$ cat terraform.plan |docker run -i justinbarrick/terraform-plan
$
```

# Usage

To use, write out a terraform plan and then parse it:

```
$ terraform plan -out=terraform.plan
$ go-terraform-plan ./terraform.plan |jq .
{
  "Stats": {
    "ToAdd": 0,
    "ToChange": 1,
    "ToDestroy": 0
  },
  "Resources": [
    {
      "Addr": {
        "Path": null,
        "Index": -1,
        "InstanceType": 1,
        "InstanceTypeSet": false,
        "Name": "kubernetes_subdomains",
        "Type": "cloudflare_record",
        "Mode": 0
      },
      "Action": "Update",
      "ActionRaw": 3,
      "Attributes": [
        {
          "Path": "ttl",
          "Action": "Update",
          "ActionRaw": 3,
          "OldValue": "1",
          "NewValue": "3600",
          "NewComputed": false,
          "Sensitive": false,
          "ForcesNew": false
        }
      ],
      "Tainted": false,
      "Deposed": false
    }
  ]
}
```
