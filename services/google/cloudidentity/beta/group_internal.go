// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl/operations"
)

func (r *Group) validate() error {

	if err := dcl.Required(r, "groupKey"); err != nil {
		return err
	}
	if err := dcl.Required(r, "parent"); err != nil {
		return err
	}
	if err := dcl.Required(r, "labels"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.GroupKey) {
		if err := r.GroupKey.validate(); err != nil {
			return err
		}
	}
	if !dcl.IsEmptyValueIndirect(r.DynamicGroupMetadata) {
		if err := r.DynamicGroupMetadata.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GroupGoogleappscloudidentitygroupsvxentitykey) validate() error {
	if err := dcl.Required(r, "id"); err != nil {
		return err
	}
	return nil
}
func (r *GroupDynamicGroupMetadata) validate() error {
	if !dcl.IsEmptyValueIndirect(r.Status) {
		if err := r.Status.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *GroupDynamicGroupMetadataQueries) validate() error {
	return nil
}
func (r *GroupDynamicGroupMetadataStatus) validate() error {
	return nil
}
func (r *GroupPosixGroups) validate() error {
	return nil
}
func (r *Group) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://cloudidentity.googleapis.com/v1beta1/", params)
}

func (r *Group) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{name}}", nr.basePath(), userBasePath, params), nil
}

func (r *Group) listURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"parent": dcl.ValueOrEmptyString(nr.Parent),
	}
	return dcl.URL("groups?view=BASIC&parent=customers/{{parent}}", nr.basePath(), userBasePath, params), nil

}

func (r *Group) createURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"initialGroupConfig": dcl.ValueOrEmptyString(nr.InitialGroupConfig),
	}
	return dcl.URL("groups?initialGroupConfig={{initialGroupConfig}}", nr.basePath(), userBasePath, params), nil

}

func (r *Group) deleteURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"name": dcl.ValueOrEmptyString(nr.Name),
	}
	return dcl.URL("groups/{{name}}", nr.basePath(), userBasePath, params), nil
}

// groupApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type groupApiOperation interface {
	do(context.Context, *Group, *Client) error
}

// newUpdateGroupUpdateGroupRequest creates a request for an
// Group resource's UpdateGroup update type by filling in the update
// fields based on the intended state of the resource.
func newUpdateGroupUpdateGroupRequest(ctx context.Context, f *Group, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := dcl.DeriveField("groups/%s", f.Name, dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["name"] = v
	}
	if v := f.DisplayName; !dcl.IsEmptyValueIndirect(v) {
		req["displayName"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.Labels; !dcl.IsEmptyValueIndirect(v) {
		req["labels"] = v
	}
	if v, err := expandGroupDynamicGroupMetadata(c, f.DynamicGroupMetadata); err != nil {
		return nil, fmt.Errorf("error expanding DynamicGroupMetadata into dynamicGroupMetadata: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["dynamicGroupMetadata"] = v
	}
	if v, err := expandGroupPosixGroupsSlice(c, f.PosixGroups); err != nil {
		return nil, fmt.Errorf("error expanding PosixGroups into posixGroups: %w", err)
	} else if v != nil {
		req["posixGroups"] = v
	}
	return req, nil
}

// marshalUpdateGroupUpdateGroupRequest converts the update into
// the final JSON request body.
func marshalUpdateGroupUpdateGroupRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updateGroupUpdateGroupOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updateGroupUpdateGroupOperation) do(ctx context.Context, r *Group, c *Client) error {
	_, err := c.GetGroup(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdateGroup")
	if err != nil {
		return err
	}
	mask := dcl.UpdateMask(op.FieldDiffs)
	u, err = dcl.AddQueryParams(u, map[string]string{"updateMask": mask})
	if err != nil {
		return err
	}

	req, err := newUpdateGroupUpdateGroupRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdateGroupUpdateGroupRequest(c, req)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "PATCH", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	err = o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET")

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) listGroupRaw(ctx context.Context, r *Group, pageToken string, pageSize int32) ([]byte, error) {
	u, err := r.urlNormalized().listURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != GroupMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listGroupOperation struct {
	Groups []map[string]interface{} `json:"groups"`
	Token  string                   `json:"nextPageToken"`
}

func (c *Client) listGroup(ctx context.Context, r *Group, pageToken string, pageSize int32) ([]*Group, string, error) {
	b, err := c.listGroupRaw(ctx, r, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listGroupOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Group
	for _, v := range m.Groups {
		res, err := unmarshalMapGroup(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Parent = r.Parent
		l = append(l, res)
	}

	return l, m.Token, nil
}

func (c *Client) deleteAllGroup(ctx context.Context, f func(*Group) bool, resources []*Group) error {
	var errors []string
	for _, res := range resources {
		if f(res) {
			// We do not want deleteAll to fail on a deletion or else it will stop deleting other resources.
			err := c.DeleteGroup(ctx, res)
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf("%v", strings.Join(errors, "\n"))
	} else {
		return nil
	}
}

type deleteGroupOperation struct{}

func (op *deleteGroupOperation) do(ctx context.Context, r *Group, c *Client) error {
	r, err := c.GetGroup(ctx, r)
	if err != nil {
		if dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.InfoWithContextf(ctx, "Group not found, returning. Original error: %v", err)
			return nil
		}
		c.Config.Logger.WarningWithContextf(ctx, "GetGroup checking for existence. error: %v", err)
		return err
	}

	u, err := r.deleteURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	// Delete should never have a body
	body := &bytes.Buffer{}
	resp, err := dcl.SendRequest(ctx, c.Config, "DELETE", u, body, c.Config.RetryProvider)
	if err != nil {
		return err
	}

	// wait for object to be deleted.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		return err
	}

	// we saw a race condition where for some successful delete operation, the Get calls returned resources for a short duration.
	// this is the reason we are adding retry to handle that case.
	maxRetry := 10
	for i := 1; i <= maxRetry; i++ {
		_, err = c.GetGroup(ctx, r)
		if !dcl.IsNotFoundOrCode(err, 403) {
			if i == maxRetry {
				return dcl.NotDeletedError{ExistingResource: r}
			}
			time.Sleep(1000 * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createGroupOperation struct {
	response map[string]interface{}
}

func (op *createGroupOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createGroupOperation) do(ctx context.Context, r *Group, c *Client) error {
	c.Config.Logger.InfoWithContextf(ctx, "Attempting to create %v", r)
	u, err := r.createURL(c.Config.BasePath)
	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}
	// wait for object to be created.
	var o operations.StandardGCPOperation
	if err := dcl.ParseResponse(resp.Response, &o); err != nil {
		return err
	}
	if err := o.Wait(context.WithValue(ctx, dcl.DoNotLogRequestsKey, true), c.Config, r.basePath(), "GET"); err != nil {
		c.Config.Logger.Warningf("Creation failed after waiting for operation: %v", err)
		return err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Successfully waited for operation")
	op.response, _ = o.FirstResponse()

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string in %v, was %T", op.response, op.response["name"])
	}
	r.Name = &name

	if _, err := c.GetGroup(ctx, r); err != nil {
		c.Config.Logger.WarningWithContextf(ctx, "get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getGroupRaw(ctx context.Context, r *Group) ([]byte, error) {

	u, err := r.getURL(c.Config.BasePath)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) groupDiffsForRawDesired(ctx context.Context, rawDesired *Group, opts ...dcl.ApplyOption) (initial, desired *Group, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Group
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Group); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Group, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeGroupDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetGroup(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFoundOrCode(err, 403) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Group resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Group resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Group resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeGroupDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Group: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Group: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeGroupInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Group: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeGroupDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Group: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffGroup(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeGroupInitialState(rawInitial, rawDesired *Group) (*Group, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeGroupDesiredState(rawDesired, rawInitial *Group, opts ...dcl.ApplyOption) (*Group, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.GroupKey = canonicalizeGroupGoogleappscloudidentitygroupsvxentitykey(rawDesired.GroupKey, nil, opts...)
		rawDesired.DynamicGroupMetadata = canonicalizeGroupDynamicGroupMetadata(rawDesired.DynamicGroupMetadata, nil, opts...)

		return rawDesired, nil
	}
	canonicalDesired := &Group{}
	if dcl.IsZeroValue(rawDesired.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	canonicalDesired.GroupKey = canonicalizeGroupGoogleappscloudidentitygroupsvxentitykey(rawDesired.GroupKey, rawInitial.GroupKey, opts...)
	if dcl.IsZeroValue(rawDesired.AdditionalGroupKeys) {
		canonicalDesired.AdditionalGroupKeys = rawInitial.AdditionalGroupKeys
	} else {
		canonicalDesired.AdditionalGroupKeys = rawDesired.AdditionalGroupKeys
	}
	if dcl.StringCanonicalize(rawDesired.Parent, rawInitial.Parent) {
		canonicalDesired.Parent = rawInitial.Parent
	} else {
		canonicalDesired.Parent = rawDesired.Parent
	}
	if dcl.StringCanonicalize(rawDesired.DisplayName, rawInitial.DisplayName) {
		canonicalDesired.DisplayName = rawInitial.DisplayName
	} else {
		canonicalDesired.DisplayName = rawDesired.DisplayName
	}
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.Labels) {
		canonicalDesired.Labels = rawInitial.Labels
	} else {
		canonicalDesired.Labels = rawDesired.Labels
	}
	canonicalDesired.DynamicGroupMetadata = canonicalizeGroupDynamicGroupMetadata(rawDesired.DynamicGroupMetadata, rawInitial.DynamicGroupMetadata, opts...)
	canonicalDesired.PosixGroups = canonicalizeGroupPosixGroupsSlice(rawDesired.PosixGroups, rawInitial.PosixGroups, opts...)
	if dcl.IsZeroValue(rawDesired.InitialGroupConfig) {
		canonicalDesired.InitialGroupConfig = rawInitial.InitialGroupConfig
	} else {
		canonicalDesired.InitialGroupConfig = rawDesired.InitialGroupConfig
	}

	return canonicalDesired, nil
}

func canonicalizeGroupNewState(c *Client, rawNew, rawDesired *Group) (*Group, error) {

	if dcl.IsNotReturnedByServer(rawNew.Name) && dcl.IsNotReturnedByServer(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.GroupKey) && dcl.IsNotReturnedByServer(rawDesired.GroupKey) {
		rawNew.GroupKey = rawDesired.GroupKey
	} else {
		rawNew.GroupKey = canonicalizeNewGroupGoogleappscloudidentitygroupsvxentitykey(c, rawDesired.GroupKey, rawNew.GroupKey)
	}

	if dcl.IsNotReturnedByServer(rawNew.AdditionalGroupKeys) && dcl.IsNotReturnedByServer(rawDesired.AdditionalGroupKeys) {
		rawNew.AdditionalGroupKeys = rawDesired.AdditionalGroupKeys
	} else {
		rawNew.AdditionalGroupKeys = rawDesired.AdditionalGroupKeys
	}

	if dcl.IsNotReturnedByServer(rawNew.Parent) && dcl.IsNotReturnedByServer(rawDesired.Parent) {
		rawNew.Parent = rawDesired.Parent
	} else {
		if dcl.StringCanonicalize(rawDesired.Parent, rawNew.Parent) {
			rawNew.Parent = rawDesired.Parent
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.DisplayName) && dcl.IsNotReturnedByServer(rawDesired.DisplayName) {
		rawNew.DisplayName = rawDesired.DisplayName
	} else {
		if dcl.StringCanonicalize(rawDesired.DisplayName, rawNew.DisplayName) {
			rawNew.DisplayName = rawDesired.DisplayName
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.CreateTime) && dcl.IsNotReturnedByServer(rawDesired.CreateTime) {
		rawNew.CreateTime = rawDesired.CreateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.Labels) && dcl.IsNotReturnedByServer(rawDesired.Labels) {
		rawNew.Labels = rawDesired.Labels
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.DynamicGroupMetadata) && dcl.IsNotReturnedByServer(rawDesired.DynamicGroupMetadata) {
		rawNew.DynamicGroupMetadata = rawDesired.DynamicGroupMetadata
	} else {
		rawNew.DynamicGroupMetadata = canonicalizeNewGroupDynamicGroupMetadata(c, rawDesired.DynamicGroupMetadata, rawNew.DynamicGroupMetadata)
	}

	if dcl.IsNotReturnedByServer(rawNew.PosixGroups) && dcl.IsNotReturnedByServer(rawDesired.PosixGroups) {
		rawNew.PosixGroups = rawDesired.PosixGroups
	} else {
		rawNew.PosixGroups = canonicalizeNewGroupPosixGroupsSlice(c, rawDesired.PosixGroups, rawNew.PosixGroups)
	}

	rawNew.InitialGroupConfig = rawDesired.InitialGroupConfig

	return rawNew, nil
}

func canonicalizeGroupGoogleappscloudidentitygroupsvxentitykey(des, initial *GroupGoogleappscloudidentitygroupsvxentitykey, opts ...dcl.ApplyOption) *GroupGoogleappscloudidentitygroupsvxentitykey {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupGoogleappscloudidentitygroupsvxentitykey{}

	if dcl.StringCanonicalize(des.Id, initial.Id) || dcl.IsZeroValue(des.Id) {
		cDes.Id = initial.Id
	} else {
		cDes.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, initial.Namespace) || dcl.IsZeroValue(des.Namespace) {
		cDes.Namespace = initial.Namespace
	} else {
		cDes.Namespace = des.Namespace
	}

	return cDes
}

func canonicalizeGroupGoogleappscloudidentitygroupsvxentitykeySlice(des, initial []GroupGoogleappscloudidentitygroupsvxentitykey, opts ...dcl.ApplyOption) []GroupGoogleappscloudidentitygroupsvxentitykey {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupGoogleappscloudidentitygroupsvxentitykey, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupGoogleappscloudidentitygroupsvxentitykey(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupGoogleappscloudidentitygroupsvxentitykey, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupGoogleappscloudidentitygroupsvxentitykey(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupGoogleappscloudidentitygroupsvxentitykey(c *Client, des, nw *GroupGoogleappscloudidentitygroupsvxentitykey) *GroupGoogleappscloudidentitygroupsvxentitykey {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupGoogleappscloudidentitygroupsvxentitykey while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Id, nw.Id) {
		nw.Id = des.Id
	}
	if dcl.StringCanonicalize(des.Namespace, nw.Namespace) {
		nw.Namespace = des.Namespace
	}

	return nw
}

func canonicalizeNewGroupGoogleappscloudidentitygroupsvxentitykeySet(c *Client, des, nw []GroupGoogleappscloudidentitygroupsvxentitykey) []GroupGoogleappscloudidentitygroupsvxentitykey {
	if des == nil {
		return nw
	}
	var reorderedNew []GroupGoogleappscloudidentitygroupsvxentitykey
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareGroupGoogleappscloudidentitygroupsvxentitykeyNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewGroupGoogleappscloudidentitygroupsvxentitykeySlice(c *Client, des, nw []GroupGoogleappscloudidentitygroupsvxentitykey) []GroupGoogleappscloudidentitygroupsvxentitykey {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupGoogleappscloudidentitygroupsvxentitykey
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupGoogleappscloudidentitygroupsvxentitykey(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDynamicGroupMetadata(des, initial *GroupDynamicGroupMetadata, opts ...dcl.ApplyOption) *GroupDynamicGroupMetadata {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDynamicGroupMetadata{}

	cDes.Queries = canonicalizeGroupDynamicGroupMetadataQueriesSlice(des.Queries, initial.Queries, opts...)

	return cDes
}

func canonicalizeGroupDynamicGroupMetadataSlice(des, initial []GroupDynamicGroupMetadata, opts ...dcl.ApplyOption) []GroupDynamicGroupMetadata {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDynamicGroupMetadata, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDynamicGroupMetadata(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDynamicGroupMetadata, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDynamicGroupMetadata(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDynamicGroupMetadata(c *Client, des, nw *GroupDynamicGroupMetadata) *GroupDynamicGroupMetadata {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDynamicGroupMetadata while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	nw.Queries = canonicalizeNewGroupDynamicGroupMetadataQueriesSlice(c, des.Queries, nw.Queries)
	nw.Status = canonicalizeNewGroupDynamicGroupMetadataStatus(c, des.Status, nw.Status)

	return nw
}

func canonicalizeNewGroupDynamicGroupMetadataSet(c *Client, des, nw []GroupDynamicGroupMetadata) []GroupDynamicGroupMetadata {
	if des == nil {
		return nw
	}
	var reorderedNew []GroupDynamicGroupMetadata
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareGroupDynamicGroupMetadataNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewGroupDynamicGroupMetadataSlice(c *Client, des, nw []GroupDynamicGroupMetadata) []GroupDynamicGroupMetadata {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDynamicGroupMetadata
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDynamicGroupMetadata(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDynamicGroupMetadataQueries(des, initial *GroupDynamicGroupMetadataQueries, opts ...dcl.ApplyOption) *GroupDynamicGroupMetadataQueries {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDynamicGroupMetadataQueries{}

	if dcl.IsZeroValue(des.ResourceType) {
		cDes.ResourceType = initial.ResourceType
	} else {
		cDes.ResourceType = des.ResourceType
	}
	if dcl.StringCanonicalize(des.Query, initial.Query) || dcl.IsZeroValue(des.Query) {
		cDes.Query = initial.Query
	} else {
		cDes.Query = des.Query
	}

	return cDes
}

func canonicalizeGroupDynamicGroupMetadataQueriesSlice(des, initial []GroupDynamicGroupMetadataQueries, opts ...dcl.ApplyOption) []GroupDynamicGroupMetadataQueries {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDynamicGroupMetadataQueries, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDynamicGroupMetadataQueries(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDynamicGroupMetadataQueries, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDynamicGroupMetadataQueries(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDynamicGroupMetadataQueries(c *Client, des, nw *GroupDynamicGroupMetadataQueries) *GroupDynamicGroupMetadataQueries {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDynamicGroupMetadataQueries while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Query, nw.Query) {
		nw.Query = des.Query
	}

	return nw
}

func canonicalizeNewGroupDynamicGroupMetadataQueriesSet(c *Client, des, nw []GroupDynamicGroupMetadataQueries) []GroupDynamicGroupMetadataQueries {
	if des == nil {
		return nw
	}
	var reorderedNew []GroupDynamicGroupMetadataQueries
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareGroupDynamicGroupMetadataQueriesNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewGroupDynamicGroupMetadataQueriesSlice(c *Client, des, nw []GroupDynamicGroupMetadataQueries) []GroupDynamicGroupMetadataQueries {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDynamicGroupMetadataQueries
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDynamicGroupMetadataQueries(c, &d, &n))
	}

	return items
}

func canonicalizeGroupDynamicGroupMetadataStatus(des, initial *GroupDynamicGroupMetadataStatus, opts ...dcl.ApplyOption) *GroupDynamicGroupMetadataStatus {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupDynamicGroupMetadataStatus{}

	if dcl.IsZeroValue(des.Status) {
		cDes.Status = initial.Status
	} else {
		cDes.Status = des.Status
	}
	if dcl.IsZeroValue(des.StatusTime) {
		cDes.StatusTime = initial.StatusTime
	} else {
		cDes.StatusTime = des.StatusTime
	}

	return cDes
}

func canonicalizeGroupDynamicGroupMetadataStatusSlice(des, initial []GroupDynamicGroupMetadataStatus, opts ...dcl.ApplyOption) []GroupDynamicGroupMetadataStatus {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupDynamicGroupMetadataStatus, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupDynamicGroupMetadataStatus(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupDynamicGroupMetadataStatus, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupDynamicGroupMetadataStatus(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupDynamicGroupMetadataStatus(c *Client, des, nw *GroupDynamicGroupMetadataStatus) *GroupDynamicGroupMetadataStatus {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupDynamicGroupMetadataStatus while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	return nw
}

func canonicalizeNewGroupDynamicGroupMetadataStatusSet(c *Client, des, nw []GroupDynamicGroupMetadataStatus) []GroupDynamicGroupMetadataStatus {
	if des == nil {
		return nw
	}
	var reorderedNew []GroupDynamicGroupMetadataStatus
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareGroupDynamicGroupMetadataStatusNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewGroupDynamicGroupMetadataStatusSlice(c *Client, des, nw []GroupDynamicGroupMetadataStatus) []GroupDynamicGroupMetadataStatus {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupDynamicGroupMetadataStatus
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupDynamicGroupMetadataStatus(c, &d, &n))
	}

	return items
}

func canonicalizeGroupPosixGroups(des, initial *GroupPosixGroups, opts ...dcl.ApplyOption) *GroupPosixGroups {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &GroupPosixGroups{}

	if dcl.StringCanonicalize(des.Name, initial.Name) || dcl.IsZeroValue(des.Name) {
		cDes.Name = initial.Name
	} else {
		cDes.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Gid, initial.Gid) || dcl.IsZeroValue(des.Gid) {
		cDes.Gid = initial.Gid
	} else {
		cDes.Gid = des.Gid
	}
	if dcl.StringCanonicalize(des.SystemId, initial.SystemId) || dcl.IsZeroValue(des.SystemId) {
		cDes.SystemId = initial.SystemId
	} else {
		cDes.SystemId = des.SystemId
	}

	return cDes
}

func canonicalizeGroupPosixGroupsSlice(des, initial []GroupPosixGroups, opts ...dcl.ApplyOption) []GroupPosixGroups {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]GroupPosixGroups, 0, len(des))
		for _, d := range des {
			cd := canonicalizeGroupPosixGroups(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]GroupPosixGroups, 0, len(des))
	for i, d := range des {
		cd := canonicalizeGroupPosixGroups(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewGroupPosixGroups(c *Client, des, nw *GroupPosixGroups) *GroupPosixGroups {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for GroupPosixGroups while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.Name, nw.Name) {
		nw.Name = des.Name
	}
	if dcl.StringCanonicalize(des.Gid, nw.Gid) {
		nw.Gid = des.Gid
	}
	if dcl.StringCanonicalize(des.SystemId, nw.SystemId) {
		nw.SystemId = des.SystemId
	}

	return nw
}

func canonicalizeNewGroupPosixGroupsSet(c *Client, des, nw []GroupPosixGroups) []GroupPosixGroups {
	if des == nil {
		return nw
	}
	var reorderedNew []GroupPosixGroups
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := compareGroupPosixGroupsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
				matchedNew = idx
				break
			}
		}
		if matchedNew != -1 {
			reorderedNew = append(reorderedNew, nw[matchedNew])
			nw = append(nw[:matchedNew], nw[matchedNew+1:]...)
		}
	}
	reorderedNew = append(reorderedNew, nw...)

	return reorderedNew
}

func canonicalizeNewGroupPosixGroupsSlice(c *Client, des, nw []GroupPosixGroups) []GroupPosixGroups {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []GroupPosixGroups
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewGroupPosixGroups(c, &d, &n))
	}

	return items
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffGroup(c *Client, desired, actual *Group, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GroupKey, actual.GroupKey, dcl.Info{ObjectFunction: compareGroupGoogleappscloudidentitygroupsvxentitykeyNewStyle, EmptyObject: EmptyGroupGoogleappscloudidentitygroupsvxentitykey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("GroupKey")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.AdditionalGroupKeys, actual.AdditionalGroupKeys, dcl.Info{ObjectFunction: compareGroupGoogleappscloudidentitygroupsvxentitykeyNewStyle, EmptyObject: EmptyGroupGoogleappscloudidentitygroupsvxentitykey, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("AdditionalGroupKeys")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Parent, actual.Parent, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Parent")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DisplayName, actual.DisplayName, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("DisplayName")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.CreateTime, actual.CreateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("CreateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.UpdateTime, actual.UpdateTime, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("UpdateTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Labels, actual.Labels, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Labels")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DynamicGroupMetadata, actual.DynamicGroupMetadata, dcl.Info{ObjectFunction: compareGroupDynamicGroupMetadataNewStyle, EmptyObject: EmptyGroupDynamicGroupMetadata, OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("DynamicGroupMetadata")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.PosixGroups, actual.PosixGroups, dcl.Info{ObjectFunction: compareGroupPosixGroupsNewStyle, EmptyObject: EmptyGroupPosixGroups, OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("PosixGroups")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.InitialGroupConfig, actual.InitialGroupConfig, dcl.Info{Type: "EnumType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("InitialGroupConfig")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}
func compareGroupGoogleappscloudidentitygroupsvxentitykeyNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupGoogleappscloudidentitygroupsvxentitykey)
	if !ok {
		desiredNotPointer, ok := d.(GroupGoogleappscloudidentitygroupsvxentitykey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupGoogleappscloudidentitygroupsvxentitykey or *GroupGoogleappscloudidentitygroupsvxentitykey", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupGoogleappscloudidentitygroupsvxentitykey)
	if !ok {
		actualNotPointer, ok := a.(GroupGoogleappscloudidentitygroupsvxentitykey)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupGoogleappscloudidentitygroupsvxentitykey", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Id, actual.Id, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Id")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Namespace, actual.Namespace, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Namespace")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDynamicGroupMetadataNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDynamicGroupMetadata)
	if !ok {
		desiredNotPointer, ok := d.(GroupDynamicGroupMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadata or *GroupDynamicGroupMetadata", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDynamicGroupMetadata)
	if !ok {
		actualNotPointer, ok := a.(GroupDynamicGroupMetadata)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadata", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Queries, actual.Queries, dcl.Info{ObjectFunction: compareGroupDynamicGroupMetadataQueriesNewStyle, EmptyObject: EmptyGroupDynamicGroupMetadataQueries, OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Queries")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{OutputOnly: true, ObjectFunction: compareGroupDynamicGroupMetadataStatusNewStyle, EmptyObject: EmptyGroupDynamicGroupMetadataStatus, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDynamicGroupMetadataQueriesNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDynamicGroupMetadataQueries)
	if !ok {
		desiredNotPointer, ok := d.(GroupDynamicGroupMetadataQueries)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataQueries or *GroupDynamicGroupMetadataQueries", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDynamicGroupMetadataQueries)
	if !ok {
		actualNotPointer, ok := a.(GroupDynamicGroupMetadataQueries)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataQueries", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.ResourceType, actual.ResourceType, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("ResourceType")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Query, actual.Query, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Query")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupDynamicGroupMetadataStatusNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupDynamicGroupMetadataStatus)
	if !ok {
		desiredNotPointer, ok := d.(GroupDynamicGroupMetadataStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataStatus or *GroupDynamicGroupMetadataStatus", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupDynamicGroupMetadataStatus)
	if !ok {
		actualNotPointer, ok := a.(GroupDynamicGroupMetadataStatus)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupDynamicGroupMetadataStatus", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Status, actual.Status, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("Status")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.StatusTime, actual.StatusTime, dcl.Info{OperationSelector: dcl.TriggersOperation("updateGroupUpdateGroupOperation")}, fn.AddNest("StatusTime")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func compareGroupPosixGroupsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*GroupPosixGroups)
	if !ok {
		desiredNotPointer, ok := d.(GroupPosixGroups)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupPosixGroups or *GroupPosixGroups", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*GroupPosixGroups)
	if !ok {
		actualNotPointer, ok := a.(GroupPosixGroups)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a GroupPosixGroups", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Gid, actual.Gid, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Gid")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SystemId, actual.SystemId, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SystemId")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

// urlNormalized returns a copy of the resource struct with values normalized
// for URL substitutions. For instance, it converts long-form self-links to
// short-form so they can be substituted in.
func (r *Group) urlNormalized() *Group {
	normalized := dcl.Copy(*r).(Group)
	normalized.Name = dcl.SelfLinkToName(r.Name)
	normalized.Parent = dcl.SelfLinkToName(r.Parent)
	normalized.DisplayName = dcl.SelfLinkToName(r.DisplayName)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.InitialGroupConfig = r.InitialGroupConfig
	return &normalized
}

func (r *Group) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdateGroup" {
		fields := map[string]interface{}{
			"name": dcl.ValueOrEmptyString(nr.Name),
		}
		return dcl.URL("groups/{{name}}", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Group resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Group) marshal(c *Client) ([]byte, error) {
	m, err := expandGroup(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Group: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalGroup decodes JSON responses into the Group resource schema.
func unmarshalGroup(b []byte, c *Client) (*Group, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapGroup(m, c)
}

func unmarshalMapGroup(m map[string]interface{}, c *Client) (*Group, error) {

	flattened := flattenGroup(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandGroup expands Group into a JSON request object.
func expandGroup(c *Client, f *Group) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := dcl.DeriveField("groups/%s", f.Name, dcl.SelfLinkToName(f.Name)); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v, err := expandGroupGoogleappscloudidentitygroupsvxentitykey(c, f.GroupKey); err != nil {
		return nil, fmt.Errorf("error expanding GroupKey into groupKey: %w", err)
	} else if v != nil {
		m["groupKey"] = v
	}
	m["additionalGroupKeys"] = f.AdditionalGroupKeys
	if v := f.Parent; dcl.ValueShouldBeSent(v) {
		m["parent"] = v
	}
	if v := f.DisplayName; dcl.ValueShouldBeSent(v) {
		m["displayName"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.Labels; dcl.ValueShouldBeSent(v) {
		m["labels"] = v
	}
	if v, err := expandGroupDynamicGroupMetadata(c, f.DynamicGroupMetadata); err != nil {
		return nil, fmt.Errorf("error expanding DynamicGroupMetadata into dynamicGroupMetadata: %w", err)
	} else if v != nil {
		m["dynamicGroupMetadata"] = v
	}
	if v, err := expandGroupPosixGroupsSlice(c, f.PosixGroups); err != nil {
		return nil, fmt.Errorf("error expanding PosixGroups into posixGroups: %w", err)
	} else {
		m["posixGroups"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding InitialGroupConfig into initialGroupConfig: %w", err)
	} else if v != nil {
		m["initialGroupConfig"] = v
	}

	return m, nil
}

// flattenGroup flattens Group from a JSON request object into the
// Group type.
func flattenGroup(c *Client, i interface{}) *Group {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Group{}
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.GroupKey = flattenGroupGoogleappscloudidentitygroupsvxentitykey(c, m["groupKey"])
	res.AdditionalGroupKeys = flattenGroupGoogleappscloudidentitygroupsvxentitykeySlice(c, m["additionalGroupKeys"])
	res.Parent = dcl.FlattenString(m["parent"])
	res.DisplayName = dcl.FlattenString(m["displayName"])
	res.Description = dcl.FlattenString(m["description"])
	res.CreateTime = dcl.FlattenString(m["createTime"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])
	res.Labels = dcl.FlattenKeyValuePairs(m["labels"])
	res.DynamicGroupMetadata = flattenGroupDynamicGroupMetadata(c, m["dynamicGroupMetadata"])
	res.PosixGroups = flattenGroupPosixGroupsSlice(c, m["posixGroups"])
	res.InitialGroupConfig = flattenGroupInitialGroupConfigEnum(m["initialGroupConfig"])

	return res
}

// expandGroupGoogleappscloudidentitygroupsvxentitykeyMap expands the contents of GroupGoogleappscloudidentitygroupsvxentitykey into a JSON
// request object.
func expandGroupGoogleappscloudidentitygroupsvxentitykeyMap(c *Client, f map[string]GroupGoogleappscloudidentitygroupsvxentitykey) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupGoogleappscloudidentitygroupsvxentitykey(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupGoogleappscloudidentitygroupsvxentitykeySlice expands the contents of GroupGoogleappscloudidentitygroupsvxentitykey into a JSON
// request object.
func expandGroupGoogleappscloudidentitygroupsvxentitykeySlice(c *Client, f []GroupGoogleappscloudidentitygroupsvxentitykey) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupGoogleappscloudidentitygroupsvxentitykey(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupGoogleappscloudidentitygroupsvxentitykeyMap flattens the contents of GroupGoogleappscloudidentitygroupsvxentitykey from a JSON
// response object.
func flattenGroupGoogleappscloudidentitygroupsvxentitykeyMap(c *Client, i interface{}) map[string]GroupGoogleappscloudidentitygroupsvxentitykey {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupGoogleappscloudidentitygroupsvxentitykey{}
	}

	if len(a) == 0 {
		return map[string]GroupGoogleappscloudidentitygroupsvxentitykey{}
	}

	items := make(map[string]GroupGoogleappscloudidentitygroupsvxentitykey)
	for k, item := range a {
		items[k] = *flattenGroupGoogleappscloudidentitygroupsvxentitykey(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGroupGoogleappscloudidentitygroupsvxentitykeySlice flattens the contents of GroupGoogleappscloudidentitygroupsvxentitykey from a JSON
// response object.
func flattenGroupGoogleappscloudidentitygroupsvxentitykeySlice(c *Client, i interface{}) []GroupGoogleappscloudidentitygroupsvxentitykey {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupGoogleappscloudidentitygroupsvxentitykey{}
	}

	if len(a) == 0 {
		return []GroupGoogleappscloudidentitygroupsvxentitykey{}
	}

	items := make([]GroupGoogleappscloudidentitygroupsvxentitykey, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupGoogleappscloudidentitygroupsvxentitykey(c, item.(map[string]interface{})))
	}

	return items
}

// expandGroupGoogleappscloudidentitygroupsvxentitykey expands an instance of GroupGoogleappscloudidentitygroupsvxentitykey into a JSON
// request object.
func expandGroupGoogleappscloudidentitygroupsvxentitykey(c *Client, f *GroupGoogleappscloudidentitygroupsvxentitykey) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Id; !dcl.IsEmptyValueIndirect(v) {
		m["id"] = v
	}
	if v := f.Namespace; !dcl.IsEmptyValueIndirect(v) {
		m["namespace"] = v
	}

	return m, nil
}

// flattenGroupGoogleappscloudidentitygroupsvxentitykey flattens an instance of GroupGoogleappscloudidentitygroupsvxentitykey from a JSON
// response object.
func flattenGroupGoogleappscloudidentitygroupsvxentitykey(c *Client, i interface{}) *GroupGoogleappscloudidentitygroupsvxentitykey {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupGoogleappscloudidentitygroupsvxentitykey{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupGoogleappscloudidentitygroupsvxentitykey
	}
	r.Id = dcl.FlattenString(m["id"])
	r.Namespace = dcl.FlattenString(m["namespace"])

	return r
}

// expandGroupDynamicGroupMetadataMap expands the contents of GroupDynamicGroupMetadata into a JSON
// request object.
func expandGroupDynamicGroupMetadataMap(c *Client, f map[string]GroupDynamicGroupMetadata) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDynamicGroupMetadata(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDynamicGroupMetadataSlice expands the contents of GroupDynamicGroupMetadata into a JSON
// request object.
func expandGroupDynamicGroupMetadataSlice(c *Client, f []GroupDynamicGroupMetadata) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDynamicGroupMetadata(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDynamicGroupMetadataMap flattens the contents of GroupDynamicGroupMetadata from a JSON
// response object.
func flattenGroupDynamicGroupMetadataMap(c *Client, i interface{}) map[string]GroupDynamicGroupMetadata {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadata{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadata{}
	}

	items := make(map[string]GroupDynamicGroupMetadata)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadata(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGroupDynamicGroupMetadataSlice flattens the contents of GroupDynamicGroupMetadata from a JSON
// response object.
func flattenGroupDynamicGroupMetadataSlice(c *Client, i interface{}) []GroupDynamicGroupMetadata {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadata{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadata{}
	}

	items := make([]GroupDynamicGroupMetadata, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadata(c, item.(map[string]interface{})))
	}

	return items
}

// expandGroupDynamicGroupMetadata expands an instance of GroupDynamicGroupMetadata into a JSON
// request object.
func expandGroupDynamicGroupMetadata(c *Client, f *GroupDynamicGroupMetadata) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v, err := expandGroupDynamicGroupMetadataQueriesSlice(c, f.Queries); err != nil {
		return nil, fmt.Errorf("error expanding Queries into queries: %w", err)
	} else if v != nil {
		m["queries"] = v
	}

	return m, nil
}

// flattenGroupDynamicGroupMetadata flattens an instance of GroupDynamicGroupMetadata from a JSON
// response object.
func flattenGroupDynamicGroupMetadata(c *Client, i interface{}) *GroupDynamicGroupMetadata {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDynamicGroupMetadata{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDynamicGroupMetadata
	}
	r.Queries = flattenGroupDynamicGroupMetadataQueriesSlice(c, m["queries"])
	r.Status = flattenGroupDynamicGroupMetadataStatus(c, m["status"])

	return r
}

// expandGroupDynamicGroupMetadataQueriesMap expands the contents of GroupDynamicGroupMetadataQueries into a JSON
// request object.
func expandGroupDynamicGroupMetadataQueriesMap(c *Client, f map[string]GroupDynamicGroupMetadataQueries) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDynamicGroupMetadataQueries(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDynamicGroupMetadataQueriesSlice expands the contents of GroupDynamicGroupMetadataQueries into a JSON
// request object.
func expandGroupDynamicGroupMetadataQueriesSlice(c *Client, f []GroupDynamicGroupMetadataQueries) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDynamicGroupMetadataQueries(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDynamicGroupMetadataQueriesMap flattens the contents of GroupDynamicGroupMetadataQueries from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesMap(c *Client, i interface{}) map[string]GroupDynamicGroupMetadataQueries {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataQueries{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataQueries{}
	}

	items := make(map[string]GroupDynamicGroupMetadataQueries)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataQueries(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGroupDynamicGroupMetadataQueriesSlice flattens the contents of GroupDynamicGroupMetadataQueries from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesSlice(c *Client, i interface{}) []GroupDynamicGroupMetadataQueries {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataQueries{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataQueries{}
	}

	items := make([]GroupDynamicGroupMetadataQueries, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataQueries(c, item.(map[string]interface{})))
	}

	return items
}

// expandGroupDynamicGroupMetadataQueries expands an instance of GroupDynamicGroupMetadataQueries into a JSON
// request object.
func expandGroupDynamicGroupMetadataQueries(c *Client, f *GroupDynamicGroupMetadataQueries) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.ResourceType; !dcl.IsEmptyValueIndirect(v) {
		m["resourceType"] = v
	}
	if v := f.Query; !dcl.IsEmptyValueIndirect(v) {
		m["query"] = v
	}

	return m, nil
}

// flattenGroupDynamicGroupMetadataQueries flattens an instance of GroupDynamicGroupMetadataQueries from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueries(c *Client, i interface{}) *GroupDynamicGroupMetadataQueries {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDynamicGroupMetadataQueries{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDynamicGroupMetadataQueries
	}
	r.ResourceType = flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(m["resourceType"])
	r.Query = dcl.FlattenString(m["query"])

	return r
}

// expandGroupDynamicGroupMetadataStatusMap expands the contents of GroupDynamicGroupMetadataStatus into a JSON
// request object.
func expandGroupDynamicGroupMetadataStatusMap(c *Client, f map[string]GroupDynamicGroupMetadataStatus) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupDynamicGroupMetadataStatus(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupDynamicGroupMetadataStatusSlice expands the contents of GroupDynamicGroupMetadataStatus into a JSON
// request object.
func expandGroupDynamicGroupMetadataStatusSlice(c *Client, f []GroupDynamicGroupMetadataStatus) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupDynamicGroupMetadataStatus(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupDynamicGroupMetadataStatusMap flattens the contents of GroupDynamicGroupMetadataStatus from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusMap(c *Client, i interface{}) map[string]GroupDynamicGroupMetadataStatus {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataStatus{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataStatus{}
	}

	items := make(map[string]GroupDynamicGroupMetadataStatus)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataStatus(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGroupDynamicGroupMetadataStatusSlice flattens the contents of GroupDynamicGroupMetadataStatus from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusSlice(c *Client, i interface{}) []GroupDynamicGroupMetadataStatus {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataStatus{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataStatus{}
	}

	items := make([]GroupDynamicGroupMetadataStatus, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataStatus(c, item.(map[string]interface{})))
	}

	return items
}

// expandGroupDynamicGroupMetadataStatus expands an instance of GroupDynamicGroupMetadataStatus into a JSON
// request object.
func expandGroupDynamicGroupMetadataStatus(c *Client, f *GroupDynamicGroupMetadataStatus) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Status; !dcl.IsEmptyValueIndirect(v) {
		m["status"] = v
	}
	if v := f.StatusTime; !dcl.IsEmptyValueIndirect(v) {
		m["statusTime"] = v
	}

	return m, nil
}

// flattenGroupDynamicGroupMetadataStatus flattens an instance of GroupDynamicGroupMetadataStatus from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatus(c *Client, i interface{}) *GroupDynamicGroupMetadataStatus {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupDynamicGroupMetadataStatus{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupDynamicGroupMetadataStatus
	}
	r.Status = flattenGroupDynamicGroupMetadataStatusStatusEnum(m["status"])
	r.StatusTime = dcl.FlattenString(m["statusTime"])

	return r
}

// expandGroupPosixGroupsMap expands the contents of GroupPosixGroups into a JSON
// request object.
func expandGroupPosixGroupsMap(c *Client, f map[string]GroupPosixGroups) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandGroupPosixGroups(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandGroupPosixGroupsSlice expands the contents of GroupPosixGroups into a JSON
// request object.
func expandGroupPosixGroupsSlice(c *Client, f []GroupPosixGroups) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandGroupPosixGroups(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenGroupPosixGroupsMap flattens the contents of GroupPosixGroups from a JSON
// response object.
func flattenGroupPosixGroupsMap(c *Client, i interface{}) map[string]GroupPosixGroups {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupPosixGroups{}
	}

	if len(a) == 0 {
		return map[string]GroupPosixGroups{}
	}

	items := make(map[string]GroupPosixGroups)
	for k, item := range a {
		items[k] = *flattenGroupPosixGroups(c, item.(map[string]interface{}))
	}

	return items
}

// flattenGroupPosixGroupsSlice flattens the contents of GroupPosixGroups from a JSON
// response object.
func flattenGroupPosixGroupsSlice(c *Client, i interface{}) []GroupPosixGroups {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupPosixGroups{}
	}

	if len(a) == 0 {
		return []GroupPosixGroups{}
	}

	items := make([]GroupPosixGroups, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupPosixGroups(c, item.(map[string]interface{})))
	}

	return items
}

// expandGroupPosixGroups expands an instance of GroupPosixGroups into a JSON
// request object.
func expandGroupPosixGroups(c *Client, f *GroupPosixGroups) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.Name; !dcl.IsEmptyValueIndirect(v) {
		m["name"] = v
	}
	if v := f.Gid; !dcl.IsEmptyValueIndirect(v) {
		m["gid"] = v
	}
	if v := f.SystemId; !dcl.IsEmptyValueIndirect(v) {
		m["systemId"] = v
	}

	return m, nil
}

// flattenGroupPosixGroups flattens an instance of GroupPosixGroups from a JSON
// response object.
func flattenGroupPosixGroups(c *Client, i interface{}) *GroupPosixGroups {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &GroupPosixGroups{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyGroupPosixGroups
	}
	r.Name = dcl.FlattenString(m["name"])
	r.Gid = dcl.FlattenString(m["gid"])
	r.SystemId = dcl.FlattenString(m["systemId"])

	return r
}

// flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumMap flattens the contents of GroupDynamicGroupMetadataQueriesResourceTypeEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumMap(c *Client, i interface{}) map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	items := make(map[string]GroupDynamicGroupMetadataQueriesResourceTypeEnum)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(item.(interface{}))
	}

	return items
}

// flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumSlice flattens the contents of GroupDynamicGroupMetadataQueriesResourceTypeEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataQueriesResourceTypeEnumSlice(c *Client, i interface{}) []GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataQueriesResourceTypeEnum{}
	}

	items := make([]GroupDynamicGroupMetadataQueriesResourceTypeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(item.(interface{})))
	}

	return items
}

// flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum asserts that an interface is a string, and returns a
// pointer to a *GroupDynamicGroupMetadataQueriesResourceTypeEnum with the same value as that string.
func flattenGroupDynamicGroupMetadataQueriesResourceTypeEnum(i interface{}) *GroupDynamicGroupMetadataQueriesResourceTypeEnum {
	s, ok := i.(string)
	if !ok {
		return GroupDynamicGroupMetadataQueriesResourceTypeEnumRef("")
	}

	return GroupDynamicGroupMetadataQueriesResourceTypeEnumRef(s)
}

// flattenGroupDynamicGroupMetadataStatusStatusEnumMap flattens the contents of GroupDynamicGroupMetadataStatusStatusEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusStatusEnumMap(c *Client, i interface{}) map[string]GroupDynamicGroupMetadataStatusStatusEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	if len(a) == 0 {
		return map[string]GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	items := make(map[string]GroupDynamicGroupMetadataStatusStatusEnum)
	for k, item := range a {
		items[k] = *flattenGroupDynamicGroupMetadataStatusStatusEnum(item.(interface{}))
	}

	return items
}

// flattenGroupDynamicGroupMetadataStatusStatusEnumSlice flattens the contents of GroupDynamicGroupMetadataStatusStatusEnum from a JSON
// response object.
func flattenGroupDynamicGroupMetadataStatusStatusEnumSlice(c *Client, i interface{}) []GroupDynamicGroupMetadataStatusStatusEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	if len(a) == 0 {
		return []GroupDynamicGroupMetadataStatusStatusEnum{}
	}

	items := make([]GroupDynamicGroupMetadataStatusStatusEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupDynamicGroupMetadataStatusStatusEnum(item.(interface{})))
	}

	return items
}

// flattenGroupDynamicGroupMetadataStatusStatusEnum asserts that an interface is a string, and returns a
// pointer to a *GroupDynamicGroupMetadataStatusStatusEnum with the same value as that string.
func flattenGroupDynamicGroupMetadataStatusStatusEnum(i interface{}) *GroupDynamicGroupMetadataStatusStatusEnum {
	s, ok := i.(string)
	if !ok {
		return GroupDynamicGroupMetadataStatusStatusEnumRef("")
	}

	return GroupDynamicGroupMetadataStatusStatusEnumRef(s)
}

// flattenGroupInitialGroupConfigEnumMap flattens the contents of GroupInitialGroupConfigEnum from a JSON
// response object.
func flattenGroupInitialGroupConfigEnumMap(c *Client, i interface{}) map[string]GroupInitialGroupConfigEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]GroupInitialGroupConfigEnum{}
	}

	if len(a) == 0 {
		return map[string]GroupInitialGroupConfigEnum{}
	}

	items := make(map[string]GroupInitialGroupConfigEnum)
	for k, item := range a {
		items[k] = *flattenGroupInitialGroupConfigEnum(item.(interface{}))
	}

	return items
}

// flattenGroupInitialGroupConfigEnumSlice flattens the contents of GroupInitialGroupConfigEnum from a JSON
// response object.
func flattenGroupInitialGroupConfigEnumSlice(c *Client, i interface{}) []GroupInitialGroupConfigEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []GroupInitialGroupConfigEnum{}
	}

	if len(a) == 0 {
		return []GroupInitialGroupConfigEnum{}
	}

	items := make([]GroupInitialGroupConfigEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenGroupInitialGroupConfigEnum(item.(interface{})))
	}

	return items
}

// flattenGroupInitialGroupConfigEnum asserts that an interface is a string, and returns a
// pointer to a *GroupInitialGroupConfigEnum with the same value as that string.
func flattenGroupInitialGroupConfigEnum(i interface{}) *GroupInitialGroupConfigEnum {
	s, ok := i.(string)
	if !ok {
		return GroupInitialGroupConfigEnumRef("")
	}

	return GroupInitialGroupConfigEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Group) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalGroup(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type groupDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         groupApiOperation
}

func convertFieldDiffsToGroupDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]groupDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff: %v", ro, fd)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []groupDiff
	// For each operation name, create a groupDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := groupDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToGroupApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToGroupApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (groupApiOperation, error) {
	switch opName {

	case "updateGroupUpdateGroupOperation":
		return &updateGroupUpdateGroupOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractGroupFields(r *Group) error {
	// *GroupGoogleappscloudidentitygroupsvxentitykey is a reused type - that's not compatible with function extractors.

	vDynamicGroupMetadata := r.DynamicGroupMetadata
	if vDynamicGroupMetadata == nil {
		// note: explicitly not the empty object.
		vDynamicGroupMetadata = &GroupDynamicGroupMetadata{}
	}
	if err := extractGroupDynamicGroupMetadataFields(r, vDynamicGroupMetadata); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDynamicGroupMetadata) {
		r.DynamicGroupMetadata = vDynamicGroupMetadata
	}
	return nil
}
func extractGroupGoogleappscloudidentitygroupsvxentitykeyFields(r *Group, o *GroupGoogleappscloudidentitygroupsvxentitykey) error {
	return nil
}
func extractGroupDynamicGroupMetadataFields(r *Group, o *GroupDynamicGroupMetadata) error {
	vStatus := o.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &GroupDynamicGroupMetadataStatus{}
	}
	if err := extractGroupDynamicGroupMetadataStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vStatus) {
		o.Status = vStatus
	}
	return nil
}
func extractGroupDynamicGroupMetadataQueriesFields(r *Group, o *GroupDynamicGroupMetadataQueries) error {
	return nil
}
func extractGroupDynamicGroupMetadataStatusFields(r *Group, o *GroupDynamicGroupMetadataStatus) error {
	return nil
}
func extractGroupPosixGroupsFields(r *Group, o *GroupPosixGroups) error {
	return nil
}

func postReadExtractGroupFields(r *Group) error {
	// *GroupGoogleappscloudidentitygroupsvxentitykey is a reused type - that's not compatible with function extractors.

	vDynamicGroupMetadata := r.DynamicGroupMetadata
	if vDynamicGroupMetadata == nil {
		// note: explicitly not the empty object.
		vDynamicGroupMetadata = &GroupDynamicGroupMetadata{}
	}
	if err := postReadExtractGroupDynamicGroupMetadataFields(r, vDynamicGroupMetadata); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vDynamicGroupMetadata) {
		r.DynamicGroupMetadata = vDynamicGroupMetadata
	}
	return nil
}
func postReadExtractGroupGoogleappscloudidentitygroupsvxentitykeyFields(r *Group, o *GroupGoogleappscloudidentitygroupsvxentitykey) error {
	return nil
}
func postReadExtractGroupDynamicGroupMetadataFields(r *Group, o *GroupDynamicGroupMetadata) error {
	vStatus := o.Status
	if vStatus == nil {
		// note: explicitly not the empty object.
		vStatus = &GroupDynamicGroupMetadataStatus{}
	}
	if err := extractGroupDynamicGroupMetadataStatusFields(r, vStatus); err != nil {
		return err
	}
	if !dcl.IsNotReturnedByServer(vStatus) {
		o.Status = vStatus
	}
	return nil
}
func postReadExtractGroupDynamicGroupMetadataQueriesFields(r *Group, o *GroupDynamicGroupMetadataQueries) error {
	return nil
}
func postReadExtractGroupDynamicGroupMetadataStatusFields(r *Group, o *GroupDynamicGroupMetadataStatus) error {
	return nil
}
func postReadExtractGroupPosixGroupsFields(r *Group, o *GroupPosixGroups) error {
	return nil
}
