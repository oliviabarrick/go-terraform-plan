package plan

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPlanParsing(t *testing.T) {
	planFile, err := os.Open("terraform.plan")
	assert.Nil(t, err)

	plan, err := ReadPlanFile(planFile)
	assert.Nil(t, err)

	resource := plan.GetResource("digitalocean_droplet.web[1]")
	assert.NotNil(t, resource)

	attr := resource.GetAttribute("image")
	assert.NotNil(t, attr)
	assert.Equal(t, "ubuntu-14-04-x64", attr.NewValue)

	attr = resource.GetAttribute("size")
	assert.NotNil(t, attr)
	assert.Equal(t, "512mb", attr.NewValue)

	attr = resource.GetAttribute("region")
	assert.NotNil(t, attr)
	assert.Equal(t, "nyc2", attr.NewValue)
}
