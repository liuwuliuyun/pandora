package generator

import (
	"fmt"
	"testing"

	"github.com/hashicorp/pandora/tools/sdk/resourcemanager"
)

func TestTemplateMethodsPandoraGet(t *testing.T) {
	input := ServiceGeneratorData{
		packageName:       "skinnyPandas",
		serviceClientName: "pandaClient",
		source:            AccTestLicenceType,
		resourceIds: map[string]resourcemanager.ResourceIdDefinition{
			"PandaPop": {
				Id: "LingLing",
			},
		},
	}

	actual, err := methodsPandoraTemplater{
		operation: resourcemanager.ApiOperation{
			ExpectedStatusCodes: []int{200},
			Method:              "GET",
			ResourceIdName:      stringPointer("PandaPop"),
			ResponseObject: &resourcemanager.ApiObjectDefinition{
				Type: resourcemanager.StringApiObjectDefinitionType,
			},
		},
		operationName: "Get",
	}.immediateOperationTemplate(input)

	if err != nil {
		t.Fatalf("err %+v", err)
	}

	expected := `
type GetOperationResponse struct {
	HttpResponse *http.Response
	OData *odata.OData
	Model *string
}

// Get ...
func (c pandaClient) Get(ctx context.Context , id PandaPop) (result GetOperationResponse, err error) {
	req, err := c.Client.NewGetRequest(ctx , id.ID(), defaultApiVersion, nil)
	if err != nil {
		return
	}

	req.ValidStatusCodes = []int{http.StatusOK}

	var resp *client.Response
	resp, err = req.Execute()
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
`
	assertTemplatedCodeMatches(t, expected, *actual)
}

func TestTemplateMethodsPandoraLROCreate(t *testing.T) {
	input := ServiceGeneratorData{
		packageName:       "skinnyPandas",
		serviceClientName: "pandaClient",
		source:            AccTestLicenceType,
		resourceIds: map[string]resourcemanager.ResourceIdDefinition{
			"PandaPop": {
				Id: "LingLing",
			},
		},
	}

	actual, err := methodsPandoraTemplater{
		operation: resourcemanager.ApiOperation{
			ExpectedStatusCodes: []int{201, 202},
			LongRunning:         true,
			Method:              "PUT",
			RequestObject: &resourcemanager.ApiObjectDefinition{
				Type: resourcemanager.StringApiObjectDefinitionType,
			},
			ResourceIdName: stringPointer("PandaPop"),
		},
		operationName: "Create",
	}.longRunningOperationTemplate(input)

	if err != nil {
		t.Fatalf("err %+v", err)
	}

	expected := `
type CreateOperationResponse struct {
	Poller *resourcemanager.LongRunningPoller
	HttpResponse *http.Response
	OData *odata.OData
}

// Create ...
func (c pandaClient) Create(ctx context.Context , id PandaPop, input string) (result CreateOperationResponse, err error) {
	req, err := c.Client.NewPutRequest(ctx , id.ID(), defaultApiVersion, nil, input)
	if err != nil {
		return
	}

	req.ValidStatusCodes = []int{http.StatusAccepted, http.StatusCreated}

	var resp *client.Response
	resp, err = req.Execute()
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.NewPollerFromResponse(ctx, resp, c.Client)
	if err != nil {
		return
	}

	return
}

// CreateThenPoll performs Create then polls until it's completed
func (c pandaClient) CreateThenPoll(ctx context.Context , id PandaPop, input string) error {
	result, err := c.Create(ctx , id, input)
	if err != nil {
		return fmt.Errorf("performing Create: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Create: %+v", err)
	}

	return nil
}
`
	assertTemplatedCodeMatches(t, expected, *actual)
}

func TestTemplateMethodsPandoraLROReboot(t *testing.T) {
	input := ServiceGeneratorData{
		packageName:       "skinnyPandas",
		serviceClientName: "pandaClient",
		source:            AccTestLicenceType,
		resourceIds: map[string]resourcemanager.ResourceIdDefinition{
			"PandaPop": {
				Id: "LingLing",
			},
		},
	}

	actual, err := methodsPandoraTemplater{
		operation: resourcemanager.ApiOperation{
			ExpectedStatusCodes: []int{200},
			LongRunning:         true,
			Method:              "POST",
			ResourceIdName:      stringPointer("PandaPop"),
		},
		operationName: "Reboot",
	}.longRunningOperationTemplate(input)

	if err != nil {
		t.Fatalf("err %+v", err)
	}

	expected := `
type RebootOperationResponse struct {
	Poller *resourcemanager.LongRunningPoller
	HttpResponse *http.Response
	OData *odata.OData
}

// Reboot ...
func (c pandaClient) Reboot(ctx context.Context , id PandaPop) (result RebootOperationResponse, err error) {
	req, err := c.Client.NewPostRequest(ctx , id.ID(), defaultApiVersion, nil, nil)
	if err != nil {
		return
	}

	req.ValidStatusCodes = []int{http.StatusOK}

	var resp *client.Response
	resp, err = req.Execute()
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.NewPollerFromResponse(ctx, resp, c.Client)
	if err != nil {
		return
	}

	return
}

// RebootThenPoll performs Reboot then polls until it's completed
func (c pandaClient) RebootThenPoll(ctx context.Context , id PandaPop) error {
	result, err := c.Reboot(ctx , id)
	if err != nil {
		return fmt.Errorf("performing Reboot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Reboot: %+v", err)
	}

	return nil
}
`
	assertTemplatedCodeMatches(t, expected, *actual)
}

func TestTemplateMethodsPandoraNonBaseTypePredicates(t *testing.T) {
	input := ServiceGeneratorData{
		packageName:       "chubbyPandas",
		serviceClientName: "pandaClient",
		source:            AccTestLicenceType,
		resourceIds: map[string]resourcemanager.ResourceIdDefinition{
			"PandaPop": {
				Id: "LingLing",
			},
		},
	}

	actual, err := methodsPandoraTemplater{
		operation: resourcemanager.ApiOperation{
			ExpectedStatusCodes:              []int{200},
			FieldContainingPaginationDetails: stringPointer("nextLink"),
			Method:                           "GET",
			ResponseObject: &resourcemanager.ApiObjectDefinition{
				Type:          resourcemanager.ReferenceApiObjectDefinitionType,
				ReferenceName: stringPointer("LingLing"),
			},
			ResourceIdName: stringPointer("PandaPop"),
			UriSuffix:      stringPointer("/pandas"),
		},
		operationName: "List",
	}.listOperationTemplate(input)

	if err != nil {
		t.Fatalf("err %+v", err)
	}

	expected := fmt.Sprintf(`
type ListOperationResponse struct {
	HttpResponse *http.Response
    OData *odata.OData
	Model *[]LingLing
}

type ListCompleteResult struct {
	Items []LingLing
}

// List ...
func (c pandaClient) List(ctx context.Context , id PandaPop) (result ListOperationResponse, err error) {
	req, err := c.Client.NewGetRequest(ctx , fmt.Sprintf("%%s/pandas", id.ID()), defaultApiVersion, nil)
	if err != nil {
		return
	}

	req.ValidStatusCodes = []int{http.StatusOK}
	var resp *client.Response
	resp, err = req.ExecutePaged()
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]LingLing %[1]s
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c pandaClient) ListComplete(ctx context.Context, id PandaPop) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, LingLingOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c pandaClient) ListCompleteMatchingPredicate(ctx context.Context, id PandaPop, predicate LingLingOperationPredicate) (resp ListCompleteResult, err error) {
	items := make([]LingLing, 0)

	result, err := c.List(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %%+v", err)
		return
	}
	if result.Model != nil {
		for _, v := range *result.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	out := ListCompleteResult{
		Items: items,
	}
	return out, nil
}
`, "`json:\"values\"`")

	assertTemplatedCodeMatches(t, expected, *actual)

}
