package gmail

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func expandStringList(configured []interface{}) []string {
	list := make([]string, 0, len(configured))
	for _, str := range configured {
		list = append(list, str.(string))
	}
	return list
}

func expandStringSet(configured *schema.Set) []string {
	return expandStringList(configured.List())
}

func flattenStringList(list []string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, v)
	}
	return vs
}

func flattenStringSet(list []string) *schema.Set {
	return schema.NewSet(schema.HashString, flattenStringList(list))
}
