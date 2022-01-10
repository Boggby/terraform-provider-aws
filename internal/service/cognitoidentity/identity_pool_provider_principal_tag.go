package cognitoidentity

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/hashicorp/aws-sdk-go-base/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func ResourceIdentityProviderPrincipalTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityProviderPrincipalTagCreate,
		Read:   resourceIdentityProviderPrincipalTagRead,
		Update: resourceIdentityProviderPrincipalTagUpdate,
		Delete: resourceIdentityProviderPrincipalTagDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"identity_provider_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 128),
				),
			},
			"identity_pool_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 55),
					validation.StringMatch(regexp.MustCompile(`^[\w-]+:[0-9a-f-]+$`), "see https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_SetPrincipalTagAttributeMap.html#API_SetPrincipalTagAttributeMap_ResponseSyntax"),
				),
			},
			"principal_tags": tftags.TagsSchema(),
			"use_defaults": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceIdentityProviderPrincipalTagCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).CognitoIdentityConn
	log.Print("[DEBUG] Creating Cognito Identity Provider Principal Tags")

	providerName := d.Get("identity_provider_name").(string)
	poolId := d.Get("identity_pool_id").(string)

	params := &cognitoidentity.SetPrincipalTagAttributeMapInput{
		IdentityPoolId:       aws.String(poolId),
		IdentityProviderName: aws.String(providerName),
	}

	if v, ok := d.GetOk("principal_tags"); ok {
		params.PrincipalTags = flex.ExpandStringMap(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("use_defaults"); ok {
		params.UseDefaults = aws.Bool(v.(bool))
	}

	_, err := conn.SetPrincipalTagAttributeMap(params)
	if err != nil {
		return fmt.Errorf("Error creating Cognito Identity Provider Principal Tags: %w", err)
	}

	d.SetId(fmt.Sprintf("%s:%s", poolId, providerName))

	return resourceIdentityProviderPrincipalTagRead(d, meta)
}

func resourceIdentityProviderPrincipalTagRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).CognitoIdentityConn
	log.Printf("[DEBUG] Reading Cognito Identity Provider Principal Tags: %s", d.Id())

	poolId, providerName, err := DecodeIdentityProviderPrincipalTagsID(d.Id())
	if err != nil {
		return err
	}

	ret, err := conn.GetPrincipalTagAttributeMap(&cognitoidentity.GetPrincipalTagAttributeMapInput{
		IdentityProviderName: aws.String(providerName),
		IdentityPoolId:       aws.String(poolId),
	})

	if err != nil {
		if tfawserr.ErrMessageContains(err, cognitoidentity.ErrCodeResourceNotFoundException, "") {
			log.Printf("[WARN] Cognito Identity Provider %q not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("identity_pool_id", ret.IdentityPoolId)
	d.Set("identity_provider_name", ret.IdentityProviderName)
	d.Set("use_defaults", ret.UseDefaults)

	if err := d.Set("principal_tags", aws.StringValueMap(ret.PrincipalTags)); err != nil {
		return fmt.Errorf("error setting attribute_mapping error: %w", err)
	}

	return nil
}

func resourceIdentityProviderPrincipalTagUpdate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).CognitoIdentityConn
	log.Print("[DEBUG] Updating Cognito Identity Provider Principal Tags")

	poolId, providerName, err := DecodeIdentityProviderPrincipalTagsID(d.Id())
	if err != nil {
		return err
	}

	params := &cognitoidentity.SetPrincipalTagAttributeMapInput{
		IdentityPoolId:       aws.String(poolId),
		IdentityProviderName: aws.String(providerName),
	}

	if d.HasChanges("principal_tags", "use_defaults") {
		params.PrincipalTags = flex.ExpandStringMap(d.Get("principal_tags").(map[string]interface{}))
		params.UseDefaults = aws.Bool(d.Get("use_defaults").(bool))

		_, err = conn.SetPrincipalTagAttributeMap(params)
		if err != nil {
			return fmt.Errorf("Error updating Cognito Identity Provider: %w", err)
		}
	}

	return resourceIdentityProviderPrincipalTagRead(d, meta)
}

func resourceIdentityProviderPrincipalTagDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).CognitoIdentityConn
	log.Printf("[DEBUG] Deleting Cognito Identity Provider Principal Tags: %s", d.Id())

	poolId, providerName, err := DecodeIdentityProviderPrincipalTagsID(d.Id())
	if err != nil {
		return err
	}
	emptyList := make(map[string]string)
	params := &cognitoidentity.SetPrincipalTagAttributeMapInput{
		IdentityPoolId:       aws.String(poolId),
		IdentityProviderName: aws.String(providerName),
		UseDefaults:          aws.Bool(true),
		PrincipalTags:        aws.StringMap(emptyList),
	}

	_, err = conn.SetPrincipalTagAttributeMap(params)

	if err != nil {
		if tfawserr.ErrMessageContains(err, cognitoidentityprovider.ErrCodeResourceNotFoundException, "") {
			return nil
		}
		return err
	}
	return nil
}

func DecodeIdentityProviderPrincipalTagsID(id string) (string, string, error) {
	idParts := strings.Split(id, ":")
	if len(idParts) <= 2 {
		return "", "", fmt.Errorf("expected ID in format UserPoolID:ProviderName, received: %s", id)
	}
	providerName := idParts[len(idParts)-1:]
	userPoolId := idParts[:len(idParts)-1]
	return strings.Join(userPoolId, ":"), strings.Join(providerName, ""), nil
}
