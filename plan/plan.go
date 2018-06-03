package plan

import (
  "io"
  "encoding/json"
  "github.com/hashicorp/terraform/terraform"
  "github.com/hashicorp/terraform/command/format"
)

type Plan struct {
  Plan *terraform.Plan
  Formatter *format.Plan
}

// Read a terraform plan from a file
func ReadPlanFile(planReader io.Reader) (Plan, error) {
  plan := Plan{}
  var err error

  plan.Plan, err = terraform.ReadPlan(planReader)
  if err != nil {
    return Plan{}, err
  }

  plan.CreateFormatter()
  return plan, nil
}

// Create a formatter from a plan.
func (p *Plan) CreateFormatter() {
  p.Formatter = format.NewPlan(p.Plan)
}

// Implements the Marshaler interface to override the DiffChangeType output.
func (p *Plan) MarshalJSON() ([]byte, error) {
  outputDiff := []InstanceDiff{}
  for _, resource := range p.Formatter.Resources {
    instance := InstanceDiff{
      Addr: resource.Addr,
      Action: DiffChangeType(resource.Action),
      ActionRaw: resource.Action,
      Tainted: resource.Tainted,
      Deposed: resource.Deposed,
    }

    for _, attrib := range resource.Attributes {
      instance.Attributes = append(instance.Attributes, &AttributeDiff{
        Path: attrib.Path,
        Action: DiffChangeType(attrib.Action),
        ActionRaw: attrib.Action,
        OldValue: attrib.OldValue,
        NewValue: attrib.NewValue,
        NewComputed: attrib.NewComputed,
        Sensitive: attrib.Sensitive,
        ForcesNew: attrib.ForcesNew,
      })
    }

    outputDiff = append(outputDiff, instance)
  }

  return json.Marshal(&struct{
    Stats format.PlanStats
    Resources []InstanceDiff
  }{
    Stats: p.Formatter.Stats(),
    Resources: outputDiff,
  })
}

func (p *Plan) PlanJson() (string, error) {
  output, err := json.Marshal(p)
  if err != nil {
    return "", err
  }
  return string(output), err
}
