package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSLogsSubscriptionFilter AWS CloudFormation Resource (AWS::Logs::SubscriptionFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
type AWSLogsSubscriptionFilter struct {

	// DestinationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-destinationarn
	DestinationArn *Value `json:"DestinationArn,omitempty"`

	// FilterPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-filterpattern
	FilterPattern *Value `json:"FilterPattern,omitempty"`

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-loggroupname
	LogGroupName *Value `json:"LogGroupName,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-rolearn
	RoleArn *Value `json:"RoleArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsSubscriptionFilter) AWSCloudFormationType() string {
	return "AWS::Logs::SubscriptionFilter"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSLogsSubscriptionFilter) MarshalJSON() ([]byte, error) {
	type Properties AWSLogsSubscriptionFilter
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSLogsSubscriptionFilter) UnmarshalJSON(b []byte) error {
	type Properties AWSLogsSubscriptionFilter
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSLogsSubscriptionFilter(*res.Properties)
	}

	return nil
}

// GetAllAWSLogsSubscriptionFilterResources retrieves all AWSLogsSubscriptionFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsSubscriptionFilterResources() map[string]AWSLogsSubscriptionFilter {
	results := map[string]AWSLogsSubscriptionFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSLogsSubscriptionFilter:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Logs::SubscriptionFilter" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSLogsSubscriptionFilter{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSLogsSubscriptionFilterWithName retrieves all AWSLogsSubscriptionFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsSubscriptionFilterWithName(name string) (AWSLogsSubscriptionFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSLogsSubscriptionFilter:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Logs::SubscriptionFilter" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSLogsSubscriptionFilter{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSLogsSubscriptionFilter{}, errors.New("resource not found")
}
