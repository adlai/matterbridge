// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ConditionalAccessApplications undocumented
type ConditionalAccessApplications struct {
	// Object is the base model of ConditionalAccessApplications
	Object
	// IncludeApplications undocumented
	IncludeApplications []string `json:"includeApplications,omitempty"`
	// ExcludeApplications undocumented
	ExcludeApplications []string `json:"excludeApplications,omitempty"`
	// IncludeUserActions undocumented
	IncludeUserActions []string `json:"includeUserActions,omitempty"`
}

// ConditionalAccessConditionSet undocumented
type ConditionalAccessConditionSet struct {
	// Object is the base model of ConditionalAccessConditionSet
	Object
	// Applications undocumented
	Applications *ConditionalAccessApplications `json:"applications,omitempty"`
	// Users undocumented
	Users *ConditionalAccessUsers `json:"users,omitempty"`
	// SignInRiskLevels undocumented
	SignInRiskLevels []RiskLevel `json:"signInRiskLevels,omitempty"`
	// Platforms undocumented
	Platforms *ConditionalAccessPlatforms `json:"platforms,omitempty"`
	// Locations undocumented
	Locations *ConditionalAccessLocations `json:"locations,omitempty"`
	// ClientAppTypes undocumented
	ClientAppTypes []ConditionalAccessClientApp `json:"clientAppTypes,omitempty"`
	// DeviceStates undocumented
	DeviceStates *ConditionalAccessDeviceStates `json:"deviceStates,omitempty"`
}

// ConditionalAccessDeviceStates undocumented
type ConditionalAccessDeviceStates struct {
	// Object is the base model of ConditionalAccessDeviceStates
	Object
	// IncludeStates undocumented
	IncludeStates []string `json:"includeStates,omitempty"`
	// ExcludeStates undocumented
	ExcludeStates []string `json:"excludeStates,omitempty"`
}

// ConditionalAccessGrantControls undocumented
type ConditionalAccessGrantControls struct {
	// Object is the base model of ConditionalAccessGrantControls
	Object
	// Operator undocumented
	Operator *string `json:"operator,omitempty"`
	// BuiltInControls undocumented
	BuiltInControls []ConditionalAccessGrantControl `json:"builtInControls,omitempty"`
	// CustomAuthenticationFactors undocumented
	CustomAuthenticationFactors []string `json:"customAuthenticationFactors,omitempty"`
	// TermsOfUse undocumented
	TermsOfUse []string `json:"termsOfUse,omitempty"`
}

// ConditionalAccessLocations undocumented
type ConditionalAccessLocations struct {
	// Object is the base model of ConditionalAccessLocations
	Object
	// IncludeLocations undocumented
	IncludeLocations []string `json:"includeLocations,omitempty"`
	// ExcludeLocations undocumented
	ExcludeLocations []string `json:"excludeLocations,omitempty"`
}

// ConditionalAccessPlatforms undocumented
type ConditionalAccessPlatforms struct {
	// Object is the base model of ConditionalAccessPlatforms
	Object
	// IncludePlatforms undocumented
	IncludePlatforms []ConditionalAccessDevicePlatform `json:"includePlatforms,omitempty"`
	// ExcludePlatforms undocumented
	ExcludePlatforms []ConditionalAccessDevicePlatform `json:"excludePlatforms,omitempty"`
}

// ConditionalAccessPolicy undocumented
type ConditionalAccessPolicy struct {
	// Entity is the base model of ConditionalAccessPolicy
	Entity
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// State undocumented
	State *ConditionalAccessPolicyState `json:"state,omitempty"`
	// Conditions undocumented
	Conditions *ConditionalAccessConditionSet `json:"conditions,omitempty"`
	// GrantControls undocumented
	GrantControls *ConditionalAccessGrantControls `json:"grantControls,omitempty"`
	// SessionControls undocumented
	SessionControls *ConditionalAccessSessionControls `json:"sessionControls,omitempty"`
}

// ConditionalAccessRoot undocumented
type ConditionalAccessRoot struct {
	// Entity is the base model of ConditionalAccessRoot
	Entity
	// Policies undocumented
	Policies []ConditionalAccessPolicy `json:"policies,omitempty"`
	// NamedLocations undocumented
	NamedLocations []NamedLocation `json:"namedLocations,omitempty"`
}

// ConditionalAccessSessionControl undocumented
type ConditionalAccessSessionControl struct {
	// Object is the base model of ConditionalAccessSessionControl
	Object
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// ConditionalAccessSessionControls undocumented
type ConditionalAccessSessionControls struct {
	// Object is the base model of ConditionalAccessSessionControls
	Object
	// ApplicationEnforcedRestrictions undocumented
	ApplicationEnforcedRestrictions *ApplicationEnforcedRestrictionsSessionControl `json:"applicationEnforcedRestrictions,omitempty"`
	// CloudAppSecurity undocumented
	CloudAppSecurity *CloudAppSecuritySessionControl `json:"cloudAppSecurity,omitempty"`
	// SignInFrequency undocumented
	SignInFrequency *SignInFrequencySessionControl `json:"signInFrequency,omitempty"`
	// PersistentBrowser undocumented
	PersistentBrowser *PersistentBrowserSessionControl `json:"persistentBrowser,omitempty"`
}

// ConditionalAccessUsers undocumented
type ConditionalAccessUsers struct {
	// Object is the base model of ConditionalAccessUsers
	Object
	// IncludeUsers undocumented
	IncludeUsers []string `json:"includeUsers,omitempty"`
	// ExcludeUsers undocumented
	ExcludeUsers []string `json:"excludeUsers,omitempty"`
	// IncludeGroups undocumented
	IncludeGroups []string `json:"includeGroups,omitempty"`
	// ExcludeGroups undocumented
	ExcludeGroups []string `json:"excludeGroups,omitempty"`
	// IncludeRoles undocumented
	IncludeRoles []string `json:"includeRoles,omitempty"`
	// ExcludeRoles undocumented
	ExcludeRoles []string `json:"excludeRoles,omitempty"`
}