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
		outputDiff = append(outputDiff, FromInstanceDiff(resource))
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

// Retrieve a resource from the plan by path.
func (p *Plan) GetResource(name string) *InstanceDiff {
	addr, _ := terraform.ParseResourceAddress(name)

	for _, resource := range p.Formatter.Resources {
		if ! resource.Addr.Equals(addr) {
			continue
		}

		r := FromInstanceDiff(resource)
		return &r
	}

	return nil
}
