// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	azuredevops "github.com/mcdafydd/go-azuredevops/azuredevops"
)

func AnyAzuredevopsEvent() azuredevops.Event {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(azuredevops.Event))(nil)).Elem()))
	var nullValue azuredevops.Event
	return nullValue
}

func EqAzuredevopsEvent(value azuredevops.Event) azuredevops.Event {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue azuredevops.Event
	return nullValue
}

func NotEqAzuredevopsEvent(value azuredevops.Event) azuredevops.Event {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue azuredevops.Event
	return nullValue
}

func AzuredevopsEventThat(matcher pegomock.ArgumentMatcher) azuredevops.Event {
	pegomock.RegisterMatcher(matcher)
	var nullValue azuredevops.Event
	return nullValue
}