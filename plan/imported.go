package plan

import (
  "encoding/json"
  "github.com/hashicorp/terraform/terraform"
)

// These all come from https://godoc.org/github.com/hashicorp/terraform/command/format
// We override them so that we can implement our own MarshalJSON function on the 
// DiffChangeType so that it it is not numerical.

type AttributeDiff struct {
  Path string

  Action    DiffChangeType
  ActionRaw terraform.DiffChangeType

  OldValue string
  NewValue string

  NewComputed bool
  Sensitive   bool
  ForcesNew   bool
}

type InstanceDiff struct {
  Addr      *terraform.ResourceAddress
  Action    DiffChangeType
  ActionRaw terraform.DiffChangeType

  Attributes []*AttributeDiff

  Tainted bool
  Deposed bool
}

type DiffChangeType terraform.DiffChangeType

const (
  DiffInvalid DiffChangeType = iota
  DiffNone
  DiffCreate
  DiffUpdate
  DiffDestroy
  DiffDestroyCreate
  DiffRefresh
)

func (d DiffChangeType) String() string {
  asStr := ""
  switch d {
    case DiffInvalid:
      asStr = "Invalid"
    case DiffNone:
      asStr = "None"
    case DiffCreate:
      asStr = "Create"
    case DiffUpdate:
      asStr = "Update"
    case DiffDestroy:
      asStr = "Destroy"
    case DiffDestroyCreate:
      asStr = "DestroyCreate"
    case DiffRefresh:
      asStr = "efresh"
    default:
      asStr = "Unknown"
  }
  return asStr
}

func (d DiffChangeType) MarshalJSON() ([]byte, error) {
  return json.Marshal(d.String())
}
