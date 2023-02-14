package botservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// DirectLineClient is the azure Bot Service is a platform for creating smart conversational agents.
type DirectLineClient struct {
	BaseClient
}

// NewDirectLineClient creates an instance of the DirectLineClient client.
func NewDirectLineClient(subscriptionID string) DirectLineClient {
	return NewDirectLineClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDirectLineClientWithBaseURI creates an instance of the DirectLineClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDirectLineClientWithBaseURI(baseURI string, subscriptionID string) DirectLineClient {
	return DirectLineClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// RegenerateKeys regenerates secret keys and returns them for the DirectLine Channel of a particular BotService
// resource
// Parameters:
// resourceGroupName - the name of the Bot resource group in the user subscription.
// resourceName - the name of the Bot resource.
// channelName - the name of the Channel resource for which keys are to be regenerated.
// parameters - the parameters to provide for the created bot.
func (client DirectLineClient) RegenerateKeys(ctx context.Context, resourceGroupName string, resourceName string, channelName RegenerateKeysChannelName, parameters SiteInfo) (result BotChannel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DirectLineClient.RegenerateKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: resourceName,
			Constraints: []validation.Constraint{{Target: "resourceName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "resourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SiteName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("botservice.DirectLineClient", "RegenerateKeys", err.Error())
	}

	req, err := client.RegenerateKeysPreparer(ctx, resourceGroupName, resourceName, channelName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.DirectLineClient", "RegenerateKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "botservice.DirectLineClient", "RegenerateKeys", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "botservice.DirectLineClient", "RegenerateKeys", resp, "Failure responding to request")
		return
	}

	return
}

// RegenerateKeysPreparer prepares the RegenerateKeys request.
func (client DirectLineClient) RegenerateKeysPreparer(ctx context.Context, resourceGroupName string, resourceName string, channelName RegenerateKeysChannelName, parameters SiteInfo) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"channelName":       autorest.Encode("path", channelName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}/regeneratekeys", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeysSender sends the RegenerateKeys request. The method will close the
// http.Response Body if it receives an error.
func (client DirectLineClient) RegenerateKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateKeysResponder handles the response to the RegenerateKeys request. The method always
// closes the http.Response Body.
func (client DirectLineClient) RegenerateKeysResponder(resp *http.Response) (result BotChannel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}