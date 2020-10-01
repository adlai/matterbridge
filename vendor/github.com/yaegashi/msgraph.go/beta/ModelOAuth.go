// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// OAuth2PermissionGrant undocumented
type OAuth2PermissionGrant struct {
	// Entity is the base model of OAuth2PermissionGrant
	Entity
	// ClientID undocumented
	ClientID *string `json:"clientId,omitempty"`
	// ConsentType undocumented
	ConsentType *string `json:"consentType,omitempty"`
	// ExpiryTime undocumented
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	// PrincipalID undocumented
	PrincipalID *string `json:"principalId,omitempty"`
	// ResourceID undocumented
	ResourceID *string `json:"resourceId,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// StartTime undocumented
	StartTime *time.Time `json:"startTime,omitempty"`
}