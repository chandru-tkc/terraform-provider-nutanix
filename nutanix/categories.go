package nutanix

import "github.com/hashicorp/terraform/helper/schema"

func categoriesSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
}

func expandCategories(categories []interface{}) map[string]string {
	output := make(map[string]string, len(categories))
	for _, v := range categories {
		var tempkey, tempvalue string
		for k, val := range v.(map[string]interface{}) {
			if k == "name" {
				tempkey = val.(string)
			} else {
				tempvalue = val.(string)
			}
		}
		output[tempkey] = tempvalue
	}
	return output
}

// func expandCategories(categories map[string]interface{}) map[string]string {
// 	output := make(map[string]string, len(categories))

// 	for i, v := range categories {
// 		output[i] = v.(string)
// 	}

// 	return output
// }
