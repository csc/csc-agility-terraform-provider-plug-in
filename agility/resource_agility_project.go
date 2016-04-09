package agility

import (
	"log"
	"os"

	"github.com/pogo61/terraform-provider-agility/agility/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAgilityProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAgilityProjectCreate,
		Read:   resourceAgilityProjectRead,
		Delete: resourceAgilityProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
		},
	}
}

func resourceAgilityProjectCreate(d *schema.ResourceData, meta interface{}) error {
	// set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

	//get the ID of the Project from the resource schema
	projectName := d.Get("name").(string)
	log.Println("the Project name is: ", projectName)

	// call the Agility API to get the ID of the Project being asked to deploy into
	response, err := api.GetProjectId(string(projectName))
	if err != nil {
		log.Println("Error in getting ProjectId: ", err)
		return err
	}

	//set the ID as the ID of this resource
	d.SetId(string(response))

	return nil
}

func resourceAgilityProjectRead(d *schema.ResourceData, meta interface{}) error {
	// there is nothing to do here

	return nil
}

func resourceAgilityProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	// call the Agility API to get the ID of the Project being asked to deploy into
	response, err := api.GetProjectId(d.Get("name").(string))
	if err != nil {
		return err
	}

	//set the ID as the ID of this resource
	d.SetId(string(response))

	return nil
}

func resourceAgilityProjectDelete(d *schema.ResourceData, meta interface{}) error {
	// we don't delete the project in agility, so just remove the resource from Terraform
	// by removing the ID
	d.SetId("")

	return nil
}
