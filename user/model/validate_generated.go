// Code generated by gen; DO NOT EDIT.

package model

import (
	"fmt"
	"strings"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
)

func ValidateUpdateProvider(oldVal, newVal Provider) error {
	fields := []string{}
	if newVal.ID != nil && !apiequality.Semantic.DeepEqual(newVal.ID, oldVal.ID){
		fields = append(fields, "id")
	}
	
	if !apiequality.Semantic.DeepEqual(newVal.Provider, oldVal.Provider){
		fields = append(fields, "provider")
	}

	if len(fields) > 0 {
		return fmt.Errorf("update field: [%s] not allowed", strings.Join(fields, ","))
	}
	return nil
}

func ValidateUpdateService(oldVal, newVal Service) error {
	fields := []string{}
	if newVal.ID != nil && !apiequality.Semantic.DeepEqual(newVal.ID, oldVal.ID){
		fields = append(fields, "id")
	}
	
	if !apiequality.Semantic.DeepEqual(newVal.Name, oldVal.Name){
		fields = append(fields, "name")
	}

	if len(fields) > 0 {
		return fmt.Errorf("update field: [%s] not allowed", strings.Join(fields, ","))
	}
	return nil
}
