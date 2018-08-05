package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func policyRuleFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"api_groups": {
			Type:        schema.TypeList,
			Description: "APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.",
			Required:    true,
			MinItems:    1,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"non_resource_urls": {
			Type:        schema.TypeList,
			Description: `NonResourceURLs is a set of partial urls that a user should have access to. *s are allowed, but only as the full, final step in the path Since non-resource URLs are not namespaced, this field is only applicable for ClusterRoles referenced from a ClusterRoleBinding. Rules can either apply to API resources (such as "pods" or "secrets") or non-resource URL paths (such as "/api"), but not both.`,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"resource_names": {
			Type:        schema.TypeList,
			Description: "ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"resources": {
			Type:        schema.TypeList,
			Description: "Resources is a list of resources this rule applies to. ResourceAll represents all resources.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"verbs": {
			Type:        schema.TypeList,
			Description: "Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
	return s
}
