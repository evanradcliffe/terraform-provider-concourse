package provider

import (
	"fmt"

	"github.com/concourse/concourse/go-concourse/concourse"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataJob() *schema.Resource {
	return &schema.Resource{
		Read: dataJobRead,

		Schema: map[string]*schema.Schema{
			"job_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pipeline_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"team_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobCreate,
		Read:   resourceJobRead,
		Update: resourceJobUpdate,
		Delete: resourceJobDelete,

		Schema: map[string]*schema.Schema{
			"job_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pipeline_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"team_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type jobHelper struct {
	JobName       string
	TeamName      string
	PipelineName  string
	IsExposed     bool
	IsPaused      bool
	JSON          string
	YAML          string
	ConfigVersion string
}

func jobID(teamName string, pipelineName string, jobName string) string {
	return fmt.Sprintf("%s:%s:%s", teamName, pipelineName, jobName)
}

func readJob(
	client concourse.Client,
	teamName string,
	pipelineName string,
	jobName string,
) (jobHelper, bool, error) {

	retVal := jobHelper{
		JobName:      jobName,
		PipelineName: pipelineName,
	}

	team := client.Team(teamName)

	job, jobFound, err := team.Job(pipelineName, jobName)
	// TODO: set job info in retVal
	fmt.Println(job)

	if err != nil {
		return retVal, false, err
	}

	if !jobFound {
		return retVal, false, nil
	}

	return retVal, true, nil
}

func dataJobRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*ProviderConfig).Client
	pipelineName := d.Get("pipeline_name").(string)
	teamName := d.Get("team_name").(string)
	jobName := d.Get("job_name").(string)

	_, wasFound, err := readJob(client, teamName, pipelineName, jobName)
	// TODO: use job to set d info

	if err != nil {
		return fmt.Errorf(
			"Error reading job %s in pipeline %s from team '%s': %s",
			jobName, pipelineName, teamName, err,
		)
	}

	if wasFound {
		d.SetId(jobID(teamName, pipelineName, jobName))
	} else {
		d.SetId("")
	}

	return nil
}

// CRUD callbacks

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {
	return resourceJobUpdate(d, m)
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	// client := m.(*ProviderConfig).Client
	// TODO: implement
	return nil
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	// client := m.(*ProviderConfig).Client
	// TODO: implement
	return nil
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	// client := m.(*ProviderConfig).Client
	// TODO: implement
	return nil
}
