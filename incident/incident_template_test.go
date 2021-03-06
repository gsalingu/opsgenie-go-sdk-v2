package incident

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &CreateIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name property cannot be empty.").Error())
	request.Name = "Incident Template Name-4"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message property cannot be empty.").Error())
	request.Message = "Incident Template Message"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Priority should be one of these: 'P1', 'P2', 'P3', 'P4' and 'P5'").Error())
	request.Priority = "P2"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of stakeholder property cannot be empty.").Error())
	request.StakeholderProperties = StakeholderProperties{Message: "Stakeholder Message"}
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates/")
	assert.Equal(t, request.Method(), http.MethodPost)
}

func TestUpdateIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &UpdateIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident Template Id property cannot be empty.").Error())
	request.IncidentTemplateId = "929fa6a4-ef29-4bda-8172-135335a9e8f2"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Name property cannot be empty.").Error())
	request.Name = "Incident Template Name-4"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message property cannot be empty.").Error())
	request.Message = "Incident Template Message"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Priority should be one of these: 'P1', 'P2', 'P3', 'P4' and 'P5'").Error())
	request.Priority = "P2"
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Message field of stakeholder property cannot be empty.").Error())
	request.StakeholderProperties = StakeholderProperties{Message: "Stakeholder Message"}
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates/929fa6a4-ef29-4bda-8172-135335a9e8f2")
	assert.Equal(t, request.Method(), http.MethodPut)
}

func TestDeleteIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &DeleteIncidentTemplateRequest{}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Incident Template Id property cannot be empty.").Error())
	request.IncidentTemplateId = "929fa6a4-ef29-4bda-8172-135335a9e8f2"
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates/929fa6a4-ef29-4bda-8172-135335a9e8f2")
	assert.Equal(t, request.Method(), http.MethodDelete)
}

func TestGetIncidentTemplateRequest_Validate(t *testing.T) {
	var err error
	request := &GetIncidentTemplateRequest{
		Limit:  20,
		Offset: 2,
		Order:  "desc",
	}
	err = request.Validate()
	assert.Nil(t, err)

	assert.Equal(t, request.ResourcePath(), "v1/incident-templates")
	assert.Equal(t, request.Method(), http.MethodGet)
}
