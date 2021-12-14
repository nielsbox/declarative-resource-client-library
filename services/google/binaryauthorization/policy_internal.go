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
package binaryauthorization

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Policy) validate() error {

	if err := dcl.ValidateAtMostOneOfFieldsSet([]string{"KubernetesNamespaceAdmissionRules", "KubernetesServiceAccountAdmissionRules", "IstioServiceIdentityAdmissionRules", "ClusterAdmissionRules"}, r.KubernetesNamespaceAdmissionRules, r.KubernetesServiceAccountAdmissionRules, r.IstioServiceIdentityAdmissionRules, r.ClusterAdmissionRules); err != nil {
		return err
	}
	if err := dcl.Required(r, "defaultAdmissionRule"); err != nil {
		return err
	}
	if !dcl.IsEmptyValueIndirect(r.DefaultAdmissionRule) {
		if err := r.DefaultAdmissionRule.validate(); err != nil {
			return err
		}
	}
	return nil
}
func (r *PolicyAdmissionWhitelistPatterns) validate() error {
	return nil
}
func (r *PolicyAdmissionRule) validate() error {
	if err := dcl.Required(r, "evaluationMode"); err != nil {
		return err
	}
	if err := dcl.Required(r, "enforcementMode"); err != nil {
		return err
	}
	return nil
}
func (r *Policy) basePath() string {
	params := map[string]interface{}{}
	return dcl.Nprintf("https://binaryauthorization.googleapis.com/v1", params)
}

func (r *Policy) getURL(userBasePath string) (string, error) {
	nr := r.urlNormalized()
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(nr.Project),
	}
	return dcl.URL("projects/{{project}}/policy", nr.basePath(), userBasePath, params), nil
}

func (r *Policy) SetPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *nr.Project,
	}
	return dcl.URL("projects/{{project}}/policy:setIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Policy) SetPolicyVerb() string {
	return "POST"
}

func (r *Policy) getPolicyURL(userBasePath string) string {
	nr := r.urlNormalized()
	fields := map[string]interface{}{
		"project": *nr.Project,
	}
	return dcl.URL("projects/{{project}}/policy:getIamPolicy", nr.basePath(), userBasePath, fields)
}

func (r *Policy) IAMPolicyVersion() int {
	return 0
}

// policyApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type policyApiOperation interface {
	do(context.Context, *Policy, *Client) error
}

// newUpdatePolicyUpdatePolicyRequest creates a request for an
// Policy resource's UpdatePolicy update type by filling in the update
// fields based on the intended state of the resource.
func newUpdatePolicyUpdatePolicyRequest(ctx context.Context, f *Policy, c *Client) (map[string]interface{}, error) {
	req := map[string]interface{}{}

	if v, err := expandPolicyAdmissionWhitelistPatternsSlice(c, f.AdmissionWhitelistPatterns); err != nil {
		return nil, fmt.Errorf("error expanding AdmissionWhitelistPatterns into admissionWhitelistPatterns: %w", err)
	} else if v != nil {
		req["admissionWhitelistPatterns"] = v
	}
	if v, err := expandPolicyAdmissionRuleMap(c, f.ClusterAdmissionRules); err != nil {
		return nil, fmt.Errorf("error expanding ClusterAdmissionRules into clusterAdmissionRules: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["clusterAdmissionRules"] = v
	}
	if v := f.KubernetesNamespaceAdmissionRules; !dcl.IsEmptyValueIndirect(v) {
		req["kubernetesNamespaceAdmissionRules"] = v
	}
	if v := f.KubernetesServiceAccountAdmissionRules; !dcl.IsEmptyValueIndirect(v) {
		req["kubernetesServiceAccountAdmissionRules"] = v
	}
	if v := f.IstioServiceIdentityAdmissionRules; !dcl.IsEmptyValueIndirect(v) {
		req["istioServiceIdentityAdmissionRules"] = v
	}
	if v, err := expandPolicyAdmissionRule(c, f.DefaultAdmissionRule); err != nil {
		return nil, fmt.Errorf("error expanding DefaultAdmissionRule into defaultAdmissionRule: %w", err)
	} else if !dcl.IsEmptyValueIndirect(v) {
		req["defaultAdmissionRule"] = v
	}
	if v := f.Description; !dcl.IsEmptyValueIndirect(v) {
		req["description"] = v
	}
	if v := f.GlobalPolicyEvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		req["globalPolicyEvaluationMode"] = v
	}
	return req, nil
}

// marshalUpdatePolicyUpdatePolicyRequest converts the update into
// the final JSON request body.
func marshalUpdatePolicyUpdatePolicyRequest(c *Client, m map[string]interface{}) ([]byte, error) {

	return json.Marshal(m)
}

type updatePolicyUpdatePolicyOperation struct {
	// If the update operation has the REQUIRES_APPLY_OPTIONS trait, this will be populated.
	// Usually it will be nil - this is to prevent us from accidentally depending on apply
	// options, which should usually be unnecessary.
	ApplyOptions []dcl.ApplyOption
	FieldDiffs   []*dcl.FieldDiff
}

// do creates a request and sends it to the appropriate URL. In most operations,
// do will transcribe a subset of the resource into a request object and send a
// PUT request to a single URL.

func (op *updatePolicyUpdatePolicyOperation) do(ctx context.Context, r *Policy, c *Client) error {
	_, err := c.GetPolicy(ctx, r)
	if err != nil {
		return err
	}

	u, err := r.updateURL(c.Config.BasePath, "UpdatePolicy")
	if err != nil {
		return err
	}

	req, err := newUpdatePolicyUpdatePolicyRequest(ctx, r, c)
	if err != nil {
		return err
	}

	c.Config.Logger.InfoWithContextf(ctx, "Created update: %#v", req)
	body, err := marshalUpdatePolicyUpdatePolicyRequest(c, req)
	if err != nil {
		return err
	}
	_, err = dcl.SendRequest(ctx, c.Config, "PUT", u, bytes.NewBuffer(body), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	return nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createPolicyOperation struct {
	response map[string]interface{}
}

func (op *createPolicyOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (c *Client) getPolicyRaw(ctx context.Context, r *Policy) ([]byte, error) {

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

func (c *Client) policyDiffsForRawDesired(ctx context.Context, rawDesired *Policy, opts ...dcl.ApplyOption) (initial, desired *Policy, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.InfoWithContext(ctx, "Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Policy
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Policy); !ok {
			c.Config.Logger.WarningWithContextf(ctx, "Initial state hint was of the wrong type; expected Policy, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetPolicy(ctx, fetchState)
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.WarningWithContextf(ctx, "Failed to retrieve whether a Policy resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Policy resource: %v", err)
		}
		c.Config.Logger.InfoWithContext(ctx, "Found that Policy resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizePolicyDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Found initial state for Policy: %v", rawInitial)
	c.Config.Logger.InfoWithContextf(ctx, "Initial desired state for Policy: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizePolicyInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized initial state for Policy: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizePolicyDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.InfoWithContextf(ctx, "Canonicalized desired state for Policy: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffPolicy(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizePolicyInitialState(rawInitial, rawDesired *Policy) (*Policy, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.

	if !dcl.IsZeroValue(rawInitial.KubernetesNamespaceAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesServiceAccountAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules, rawInitial.ClusterAdmissionRules) {
			rawInitial.KubernetesNamespaceAdmissionRules = map[string]PolicyAdmissionRule{}
		}
	}

	if !dcl.IsZeroValue(rawInitial.KubernetesServiceAccountAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesNamespaceAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules, rawInitial.ClusterAdmissionRules) {
			rawInitial.KubernetesServiceAccountAdmissionRules = map[string]PolicyAdmissionRule{}
		}
	}

	if !dcl.IsZeroValue(rawInitial.IstioServiceIdentityAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesNamespaceAdmissionRules, rawInitial.KubernetesServiceAccountAdmissionRules, rawInitial.ClusterAdmissionRules) {
			rawInitial.IstioServiceIdentityAdmissionRules = map[string]PolicyAdmissionRule{}
		}
	}

	if !dcl.IsZeroValue(rawInitial.ClusterAdmissionRules) {
		// Check if anything else is set.
		if dcl.AnySet(rawInitial.KubernetesNamespaceAdmissionRules, rawInitial.KubernetesServiceAccountAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules) {
			rawInitial.ClusterAdmissionRules = map[string]PolicyAdmissionRule{}
		}
	}

	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizePolicyDesiredState(rawDesired, rawInitial *Policy, opts ...dcl.ApplyOption) (*Policy, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.
		rawDesired.DefaultAdmissionRule = canonicalizePolicyAdmissionRule(rawDesired.DefaultAdmissionRule, nil, opts...)

		return rawDesired, nil
	}

	if rawDesired.KubernetesNamespaceAdmissionRules != nil || rawInitial.KubernetesNamespaceAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesServiceAccountAdmissionRules, rawDesired.IstioServiceIdentityAdmissionRules, rawDesired.ClusterAdmissionRules) {
			rawDesired.KubernetesNamespaceAdmissionRules = nil
			rawInitial.KubernetesNamespaceAdmissionRules = nil
		}
	}

	if rawDesired.KubernetesServiceAccountAdmissionRules != nil || rawInitial.KubernetesServiceAccountAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesNamespaceAdmissionRules, rawDesired.IstioServiceIdentityAdmissionRules, rawDesired.ClusterAdmissionRules) {
			rawDesired.KubernetesServiceAccountAdmissionRules = nil
			rawInitial.KubernetesServiceAccountAdmissionRules = nil
		}
	}

	if rawDesired.IstioServiceIdentityAdmissionRules != nil || rawInitial.IstioServiceIdentityAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesNamespaceAdmissionRules, rawDesired.KubernetesServiceAccountAdmissionRules, rawDesired.ClusterAdmissionRules) {
			rawDesired.IstioServiceIdentityAdmissionRules = nil
			rawInitial.IstioServiceIdentityAdmissionRules = nil
		}
	}

	if rawDesired.ClusterAdmissionRules != nil || rawInitial.ClusterAdmissionRules != nil {
		// Check if anything else is set.
		if dcl.AnySet(rawDesired.KubernetesNamespaceAdmissionRules, rawDesired.KubernetesServiceAccountAdmissionRules, rawDesired.IstioServiceIdentityAdmissionRules) {
			rawDesired.ClusterAdmissionRules = nil
			rawInitial.ClusterAdmissionRules = nil
		}
	}

	canonicalDesired := &Policy{}
	canonicalDesired.AdmissionWhitelistPatterns = canonicalizePolicyAdmissionWhitelistPatternsSlice(rawDesired.AdmissionWhitelistPatterns, rawInitial.AdmissionWhitelistPatterns, opts...)
	if dcl.IsZeroValue(rawDesired.ClusterAdmissionRules) {
		canonicalDesired.ClusterAdmissionRules = rawInitial.ClusterAdmissionRules
	} else {
		canonicalDesired.ClusterAdmissionRules = rawDesired.ClusterAdmissionRules
	}
	if dcl.IsZeroValue(rawDesired.KubernetesNamespaceAdmissionRules) {
		canonicalDesired.KubernetesNamespaceAdmissionRules = rawInitial.KubernetesNamespaceAdmissionRules
	} else {
		canonicalDesired.KubernetesNamespaceAdmissionRules = rawDesired.KubernetesNamespaceAdmissionRules
	}
	if dcl.IsZeroValue(rawDesired.KubernetesServiceAccountAdmissionRules) {
		canonicalDesired.KubernetesServiceAccountAdmissionRules = rawInitial.KubernetesServiceAccountAdmissionRules
	} else {
		canonicalDesired.KubernetesServiceAccountAdmissionRules = rawDesired.KubernetesServiceAccountAdmissionRules
	}
	if canonicalizePolicyISIAR(rawDesired.IstioServiceIdentityAdmissionRules, rawInitial.IstioServiceIdentityAdmissionRules) {
		canonicalDesired.IstioServiceIdentityAdmissionRules = rawInitial.IstioServiceIdentityAdmissionRules
	} else {
		canonicalDesired.IstioServiceIdentityAdmissionRules = rawDesired.IstioServiceIdentityAdmissionRules
	}
	canonicalDesired.DefaultAdmissionRule = canonicalizePolicyAdmissionRule(rawDesired.DefaultAdmissionRule, rawInitial.DefaultAdmissionRule, opts...)
	if dcl.StringCanonicalize(rawDesired.Description, rawInitial.Description) {
		canonicalDesired.Description = rawInitial.Description
	} else {
		canonicalDesired.Description = rawDesired.Description
	}
	if dcl.IsZeroValue(rawDesired.GlobalPolicyEvaluationMode) {
		canonicalDesired.GlobalPolicyEvaluationMode = rawInitial.GlobalPolicyEvaluationMode
	} else {
		canonicalDesired.GlobalPolicyEvaluationMode = rawDesired.GlobalPolicyEvaluationMode
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizePolicyNewState(c *Client, rawNew, rawDesired *Policy) (*Policy, error) {

	if dcl.IsNotReturnedByServer(rawNew.AdmissionWhitelistPatterns) && dcl.IsNotReturnedByServer(rawDesired.AdmissionWhitelistPatterns) {
		rawNew.AdmissionWhitelistPatterns = rawDesired.AdmissionWhitelistPatterns
	} else {
		rawNew.AdmissionWhitelistPatterns = canonicalizeNewPolicyAdmissionWhitelistPatternsSlice(c, rawDesired.AdmissionWhitelistPatterns, rawNew.AdmissionWhitelistPatterns)
	}

	if dcl.IsNotReturnedByServer(rawNew.ClusterAdmissionRules) && dcl.IsNotReturnedByServer(rawDesired.ClusterAdmissionRules) {
		rawNew.ClusterAdmissionRules = rawDesired.ClusterAdmissionRules
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.KubernetesNamespaceAdmissionRules) && dcl.IsNotReturnedByServer(rawDesired.KubernetesNamespaceAdmissionRules) {
		rawNew.KubernetesNamespaceAdmissionRules = rawDesired.KubernetesNamespaceAdmissionRules
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.KubernetesServiceAccountAdmissionRules) && dcl.IsNotReturnedByServer(rawDesired.KubernetesServiceAccountAdmissionRules) {
		rawNew.KubernetesServiceAccountAdmissionRules = rawDesired.KubernetesServiceAccountAdmissionRules
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.IstioServiceIdentityAdmissionRules) && dcl.IsNotReturnedByServer(rawDesired.IstioServiceIdentityAdmissionRules) {
		rawNew.IstioServiceIdentityAdmissionRules = rawDesired.IstioServiceIdentityAdmissionRules
	} else {
		if canonicalizePolicyISIAR(rawDesired.IstioServiceIdentityAdmissionRules, rawNew.IstioServiceIdentityAdmissionRules) {
			rawNew.IstioServiceIdentityAdmissionRules = rawDesired.IstioServiceIdentityAdmissionRules
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.DefaultAdmissionRule) && dcl.IsNotReturnedByServer(rawDesired.DefaultAdmissionRule) {
		rawNew.DefaultAdmissionRule = rawDesired.DefaultAdmissionRule
	} else {
		rawNew.DefaultAdmissionRule = canonicalizeNewPolicyAdmissionRule(c, rawDesired.DefaultAdmissionRule, rawNew.DefaultAdmissionRule)
	}

	if dcl.IsNotReturnedByServer(rawNew.Description) && dcl.IsNotReturnedByServer(rawDesired.Description) {
		rawNew.Description = rawDesired.Description
	} else {
		if dcl.StringCanonicalize(rawDesired.Description, rawNew.Description) {
			rawNew.Description = rawDesired.Description
		}
	}

	if dcl.IsNotReturnedByServer(rawNew.GlobalPolicyEvaluationMode) && dcl.IsNotReturnedByServer(rawDesired.GlobalPolicyEvaluationMode) {
		rawNew.GlobalPolicyEvaluationMode = rawDesired.GlobalPolicyEvaluationMode
	} else {
	}

	if dcl.IsNotReturnedByServer(rawNew.SelfLink) && dcl.IsNotReturnedByServer(rawDesired.SelfLink) {
		rawNew.SelfLink = rawDesired.SelfLink
	} else {
		if dcl.PartialSelfLinkToSelfLink(rawDesired.SelfLink, rawNew.SelfLink) {
			rawNew.SelfLink = rawDesired.SelfLink
		}
	}

	rawNew.Project = rawDesired.Project

	if dcl.IsNotReturnedByServer(rawNew.UpdateTime) && dcl.IsNotReturnedByServer(rawDesired.UpdateTime) {
		rawNew.UpdateTime = rawDesired.UpdateTime
	} else {
	}

	return rawNew, nil
}

func canonicalizePolicyAdmissionWhitelistPatterns(des, initial *PolicyAdmissionWhitelistPatterns, opts ...dcl.ApplyOption) *PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyAdmissionWhitelistPatterns{}

	if dcl.StringCanonicalize(des.NamePattern, initial.NamePattern) || dcl.IsZeroValue(des.NamePattern) {
		cDes.NamePattern = initial.NamePattern
	} else {
		cDes.NamePattern = des.NamePattern
	}

	return cDes
}

func canonicalizePolicyAdmissionWhitelistPatternsSlice(des, initial []PolicyAdmissionWhitelistPatterns, opts ...dcl.ApplyOption) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyAdmissionWhitelistPatterns, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyAdmissionWhitelistPatterns(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyAdmissionWhitelistPatterns, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyAdmissionWhitelistPatterns(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyAdmissionWhitelistPatterns(c *Client, des, nw *PolicyAdmissionWhitelistPatterns) *PolicyAdmissionWhitelistPatterns {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyAdmissionWhitelistPatterns while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringCanonicalize(des.NamePattern, nw.NamePattern) {
		nw.NamePattern = des.NamePattern
	}

	return nw
}

func canonicalizeNewPolicyAdmissionWhitelistPatternsSet(c *Client, des, nw []PolicyAdmissionWhitelistPatterns) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return nw
	}
	var reorderedNew []PolicyAdmissionWhitelistPatterns
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePolicyAdmissionWhitelistPatternsNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewPolicyAdmissionWhitelistPatternsSlice(c *Client, des, nw []PolicyAdmissionWhitelistPatterns) []PolicyAdmissionWhitelistPatterns {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyAdmissionWhitelistPatterns
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyAdmissionWhitelistPatterns(c, &d, &n))
	}

	return items
}

func canonicalizePolicyAdmissionRule(des, initial *PolicyAdmissionRule, opts ...dcl.ApplyOption) *PolicyAdmissionRule {
	if des == nil {
		return initial
	}
	if des.empty {
		return des
	}

	if initial == nil {
		return des
	}

	cDes := &PolicyAdmissionRule{}

	if dcl.IsZeroValue(des.EvaluationMode) {
		cDes.EvaluationMode = initial.EvaluationMode
	} else {
		cDes.EvaluationMode = des.EvaluationMode
	}
	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, initial.RequireAttestationsBy) || dcl.IsZeroValue(des.RequireAttestationsBy) {
		cDes.RequireAttestationsBy = initial.RequireAttestationsBy
	} else {
		cDes.RequireAttestationsBy = des.RequireAttestationsBy
	}
	if dcl.IsZeroValue(des.EnforcementMode) {
		cDes.EnforcementMode = initial.EnforcementMode
	} else {
		cDes.EnforcementMode = des.EnforcementMode
	}

	return cDes
}

func canonicalizePolicyAdmissionRuleSlice(des, initial []PolicyAdmissionRule, opts ...dcl.ApplyOption) []PolicyAdmissionRule {
	if des == nil {
		return initial
	}

	if len(des) != len(initial) {

		items := make([]PolicyAdmissionRule, 0, len(des))
		for _, d := range des {
			cd := canonicalizePolicyAdmissionRule(&d, nil, opts...)
			if cd != nil {
				items = append(items, *cd)
			}
		}
		return items
	}

	items := make([]PolicyAdmissionRule, 0, len(des))
	for i, d := range des {
		cd := canonicalizePolicyAdmissionRule(&d, &initial[i], opts...)
		if cd != nil {
			items = append(items, *cd)
		}
	}
	return items

}

func canonicalizeNewPolicyAdmissionRule(c *Client, des, nw *PolicyAdmissionRule) *PolicyAdmissionRule {

	if des == nil {
		return nw
	}

	if nw == nil {
		if dcl.IsNotReturnedByServer(des) {
			c.Config.Logger.Info("Found explicitly empty value for PolicyAdmissionRule while comparing non-nil desired to nil actual.  Returning desired object.")
			return des
		}
		return nil
	}

	if dcl.StringArrayCanonicalize(des.RequireAttestationsBy, nw.RequireAttestationsBy) {
		nw.RequireAttestationsBy = des.RequireAttestationsBy
	}

	return nw
}

func canonicalizeNewPolicyAdmissionRuleSet(c *Client, des, nw []PolicyAdmissionRule) []PolicyAdmissionRule {
	if des == nil {
		return nw
	}
	var reorderedNew []PolicyAdmissionRule
	for _, d := range des {
		matchedNew := -1
		for idx, n := range nw {
			if diffs, _ := comparePolicyAdmissionRuleNewStyle(&d, &n, dcl.FieldName{}); len(diffs) == 0 {
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

func canonicalizeNewPolicyAdmissionRuleSlice(c *Client, des, nw []PolicyAdmissionRule) []PolicyAdmissionRule {
	if des == nil {
		return nw
	}

	// Lengths are unequal. A diff will occur later, so we shouldn't canonicalize.
	// Return the original array.
	if len(des) != len(nw) {
		return nw
	}

	var items []PolicyAdmissionRule
	for i, d := range des {
		n := nw[i]
		items = append(items, *canonicalizeNewPolicyAdmissionRule(c, &d, &n))
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
func diffPolicy(c *Client, desired, actual *Policy, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	c.Config.Logger.Infof("Diff function called with desired state: %v", desired)
	c.Config.Logger.Infof("Diff function called with actual state: %v", actual)

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.AdmissionWhitelistPatterns, actual.AdmissionWhitelistPatterns, dcl.Info{ObjectFunction: comparePolicyAdmissionWhitelistPatternsNewStyle, EmptyObject: EmptyPolicyAdmissionWhitelistPatterns, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("AdmissionWhitelistPatterns")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.ClusterAdmissionRules, actual.ClusterAdmissionRules, dcl.Info{ObjectFunction: comparePolicyAdmissionRuleNewStyle, EmptyObject: EmptyPolicyAdmissionRule, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("ClusterAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesNamespaceAdmissionRules, actual.KubernetesNamespaceAdmissionRules, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("KubernetesNamespaceAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.KubernetesServiceAccountAdmissionRules, actual.KubernetesServiceAccountAdmissionRules, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("KubernetesServiceAccountAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.IstioServiceIdentityAdmissionRules, actual.IstioServiceIdentityAdmissionRules, dcl.Info{CustomDiff: canonicalizePolicyISIAR, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("IstioServiceIdentityAdmissionRules")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.DefaultAdmissionRule, actual.DefaultAdmissionRule, dcl.Info{ObjectFunction: comparePolicyAdmissionRuleNewStyle, EmptyObject: EmptyPolicyAdmissionRule, OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("DefaultAdmissionRule")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Description, actual.Description, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("Description")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.GlobalPolicyEvaluationMode, actual.GlobalPolicyEvaluationMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("GlobalPolicyEvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SelfLink, actual.SelfLink, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
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

	return newDiffs, nil
}
func comparePolicyAdmissionWhitelistPatternsNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyAdmissionWhitelistPatterns)
	if !ok {
		desiredNotPointer, ok := d.(PolicyAdmissionWhitelistPatterns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionWhitelistPatterns or *PolicyAdmissionWhitelistPatterns", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyAdmissionWhitelistPatterns)
	if !ok {
		actualNotPointer, ok := a.(PolicyAdmissionWhitelistPatterns)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionWhitelistPatterns", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.NamePattern, actual.NamePattern, dcl.Info{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("NamePattern")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}
	return diffs, nil
}

func comparePolicyAdmissionRuleNewStyle(d, a interface{}, fn dcl.FieldName) ([]*dcl.FieldDiff, error) {
	var diffs []*dcl.FieldDiff

	desired, ok := d.(*PolicyAdmissionRule)
	if !ok {
		desiredNotPointer, ok := d.(PolicyAdmissionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionRule or *PolicyAdmissionRule", d)
		}
		desired = &desiredNotPointer
	}
	actual, ok := a.(*PolicyAdmissionRule)
	if !ok {
		actualNotPointer, ok := a.(PolicyAdmissionRule)
		if !ok {
			return nil, fmt.Errorf("obj %v is not a PolicyAdmissionRule", a)
		}
		actual = &actualNotPointer
	}

	if ds, err := dcl.Diff(desired.EvaluationMode, actual.EvaluationMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EvaluationMode")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.RequireAttestationsBy, actual.RequireAttestationsBy, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("RequireAttestationsBy")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, ds...)
	}

	if ds, err := dcl.Diff(desired.EnforcementMode, actual.EnforcementMode, dcl.Info{Type: "EnumType", OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, fn.AddNest("EnforcementMode")); len(ds) != 0 || err != nil {
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
func (r *Policy) urlNormalized() *Policy {
	normalized := dcl.Copy(*r).(Policy)
	normalized.Description = dcl.SelfLinkToName(r.Description)
	normalized.SelfLink = dcl.SelfLinkToName(r.SelfLink)
	normalized.Project = dcl.SelfLinkToName(r.Project)
	return &normalized
}

func (r *Policy) updateURL(userBasePath, updateName string) (string, error) {
	nr := r.urlNormalized()
	if updateName == "UpdatePolicy" {
		fields := map[string]interface{}{
			"project": dcl.ValueOrEmptyString(nr.Project),
		}
		return dcl.URL("projects/{{project}}/policy", nr.basePath(), userBasePath, fields), nil

	}

	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Policy resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Policy) marshal(c *Client) ([]byte, error) {
	m, err := expandPolicy(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Policy: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalPolicy decodes JSON responses into the Policy resource schema.
func unmarshalPolicy(b []byte, c *Client) (*Policy, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapPolicy(m, c)
}

func unmarshalMapPolicy(m map[string]interface{}, c *Client) (*Policy, error) {

	flattened := flattenPolicy(c, m)
	if flattened == nil {
		return nil, fmt.Errorf("attempted to flatten empty json object")
	}
	return flattened, nil
}

// expandPolicy expands Policy into a JSON request object.
func expandPolicy(c *Client, f *Policy) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v, err := expandPolicyAdmissionWhitelistPatternsSlice(c, f.AdmissionWhitelistPatterns); err != nil {
		return nil, fmt.Errorf("error expanding AdmissionWhitelistPatterns into admissionWhitelistPatterns: %w", err)
	} else {
		m["admissionWhitelistPatterns"] = v
	}
	if v, err := expandPolicyAdmissionRuleMap(c, f.ClusterAdmissionRules); err != nil {
		return nil, fmt.Errorf("error expanding ClusterAdmissionRules into clusterAdmissionRules: %w", err)
	} else if v != nil {
		m["clusterAdmissionRules"] = v
	}
	if v := f.KubernetesNamespaceAdmissionRules; dcl.ValueShouldBeSent(v) {
		m["kubernetesNamespaceAdmissionRules"] = v
	}
	if v := f.KubernetesServiceAccountAdmissionRules; dcl.ValueShouldBeSent(v) {
		m["kubernetesServiceAccountAdmissionRules"] = v
	}
	if v := f.IstioServiceIdentityAdmissionRules; dcl.ValueShouldBeSent(v) {
		m["istioServiceIdentityAdmissionRules"] = v
	}
	if v, err := expandPolicyAdmissionRule(c, f.DefaultAdmissionRule); err != nil {
		return nil, fmt.Errorf("error expanding DefaultAdmissionRule into defaultAdmissionRule: %w", err)
	} else if v != nil {
		m["defaultAdmissionRule"] = v
	}
	if v := f.Description; dcl.ValueShouldBeSent(v) {
		m["description"] = v
	}
	if v := f.GlobalPolicyEvaluationMode; dcl.ValueShouldBeSent(v) {
		m["globalPolicyEvaluationMode"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenPolicy flattens Policy from a JSON request object into the
// Policy type.
func flattenPolicy(c *Client, i interface{}) *Policy {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Policy{}
	res.AdmissionWhitelistPatterns = flattenPolicyAdmissionWhitelistPatternsSlice(c, m["admissionWhitelistPatterns"])
	res.ClusterAdmissionRules = flattenPolicyAdmissionRuleMap(c, m["clusterAdmissionRules"])
	res.KubernetesNamespaceAdmissionRules = flattenPolicyAdmissionRuleMap(c, m["kubernetesNamespaceAdmissionRules"])
	res.KubernetesServiceAccountAdmissionRules = flattenPolicyAdmissionRuleMap(c, m["kubernetesServiceAccountAdmissionRules"])
	res.IstioServiceIdentityAdmissionRules = flattenPolicyAdmissionRuleMap(c, m["istioServiceIdentityAdmissionRules"])
	res.DefaultAdmissionRule = flattenPolicyAdmissionRule(c, m["defaultAdmissionRule"])
	res.Description = dcl.FlattenString(m["description"])
	res.GlobalPolicyEvaluationMode = flattenPolicyGlobalPolicyEvaluationModeEnum(m["globalPolicyEvaluationMode"])
	res.SelfLink = dcl.FlattenString(m["name"])
	res.Project = dcl.FlattenString(m["project"])
	res.UpdateTime = dcl.FlattenString(m["updateTime"])

	return res
}

// expandPolicyAdmissionWhitelistPatternsMap expands the contents of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatternsMap(c *Client, f map[string]PolicyAdmissionWhitelistPatterns) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyAdmissionWhitelistPatterns(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyAdmissionWhitelistPatternsSlice expands the contents of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatternsSlice(c *Client, f []PolicyAdmissionWhitelistPatterns) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyAdmissionWhitelistPatterns(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyAdmissionWhitelistPatternsMap flattens the contents of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatternsMap(c *Client, i interface{}) map[string]PolicyAdmissionWhitelistPatterns {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyAdmissionWhitelistPatterns{}
	}

	if len(a) == 0 {
		return map[string]PolicyAdmissionWhitelistPatterns{}
	}

	items := make(map[string]PolicyAdmissionWhitelistPatterns)
	for k, item := range a {
		items[k] = *flattenPolicyAdmissionWhitelistPatterns(c, item.(map[string]interface{}))
	}

	return items
}

// flattenPolicyAdmissionWhitelistPatternsSlice flattens the contents of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatternsSlice(c *Client, i interface{}) []PolicyAdmissionWhitelistPatterns {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyAdmissionWhitelistPatterns{}
	}

	if len(a) == 0 {
		return []PolicyAdmissionWhitelistPatterns{}
	}

	items := make([]PolicyAdmissionWhitelistPatterns, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyAdmissionWhitelistPatterns(c, item.(map[string]interface{})))
	}

	return items
}

// expandPolicyAdmissionWhitelistPatterns expands an instance of PolicyAdmissionWhitelistPatterns into a JSON
// request object.
func expandPolicyAdmissionWhitelistPatterns(c *Client, f *PolicyAdmissionWhitelistPatterns) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.NamePattern; !dcl.IsEmptyValueIndirect(v) {
		m["namePattern"] = v
	}

	return m, nil
}

// flattenPolicyAdmissionWhitelistPatterns flattens an instance of PolicyAdmissionWhitelistPatterns from a JSON
// response object.
func flattenPolicyAdmissionWhitelistPatterns(c *Client, i interface{}) *PolicyAdmissionWhitelistPatterns {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyAdmissionWhitelistPatterns{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyAdmissionWhitelistPatterns
	}
	r.NamePattern = dcl.FlattenString(m["namePattern"])

	return r
}

// expandPolicyAdmissionRuleMap expands the contents of PolicyAdmissionRule into a JSON
// request object.
func expandPolicyAdmissionRuleMap(c *Client, f map[string]PolicyAdmissionRule) (map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := make(map[string]interface{})
	for k, item := range f {
		i, err := expandPolicyAdmissionRule(c, &item)
		if err != nil {
			return nil, err
		}
		if i != nil {
			items[k] = i
		}
	}

	return items, nil
}

// expandPolicyAdmissionRuleSlice expands the contents of PolicyAdmissionRule into a JSON
// request object.
func expandPolicyAdmissionRuleSlice(c *Client, f []PolicyAdmissionRule) ([]map[string]interface{}, error) {
	if f == nil {
		return nil, nil
	}

	items := []map[string]interface{}{}
	for _, item := range f {
		i, err := expandPolicyAdmissionRule(c, &item)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

// flattenPolicyAdmissionRuleMap flattens the contents of PolicyAdmissionRule from a JSON
// response object.
func flattenPolicyAdmissionRuleMap(c *Client, i interface{}) map[string]PolicyAdmissionRule {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyAdmissionRule{}
	}

	if len(a) == 0 {
		return map[string]PolicyAdmissionRule{}
	}

	items := make(map[string]PolicyAdmissionRule)
	for k, item := range a {
		items[k] = *flattenPolicyAdmissionRule(c, item.(map[string]interface{}))
	}

	return items
}

// flattenPolicyAdmissionRuleSlice flattens the contents of PolicyAdmissionRule from a JSON
// response object.
func flattenPolicyAdmissionRuleSlice(c *Client, i interface{}) []PolicyAdmissionRule {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyAdmissionRule{}
	}

	if len(a) == 0 {
		return []PolicyAdmissionRule{}
	}

	items := make([]PolicyAdmissionRule, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyAdmissionRule(c, item.(map[string]interface{})))
	}

	return items
}

// expandPolicyAdmissionRule expands an instance of PolicyAdmissionRule into a JSON
// request object.
func expandPolicyAdmissionRule(c *Client, f *PolicyAdmissionRule) (map[string]interface{}, error) {
	if dcl.IsEmptyValueIndirect(f) {
		return nil, nil
	}

	m := make(map[string]interface{})
	if v := f.EvaluationMode; !dcl.IsEmptyValueIndirect(v) {
		m["evaluationMode"] = v
	}
	if v := f.RequireAttestationsBy; v != nil {
		m["requireAttestationsBy"] = v
	}
	if v := f.EnforcementMode; !dcl.IsEmptyValueIndirect(v) {
		m["enforcementMode"] = v
	}

	return m, nil
}

// flattenPolicyAdmissionRule flattens an instance of PolicyAdmissionRule from a JSON
// response object.
func flattenPolicyAdmissionRule(c *Client, i interface{}) *PolicyAdmissionRule {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}

	r := &PolicyAdmissionRule{}

	if dcl.IsEmptyValueIndirect(i) {
		return EmptyPolicyAdmissionRule
	}
	r.EvaluationMode = flattenPolicyAdmissionRuleEvaluationModeEnum(m["evaluationMode"])
	r.RequireAttestationsBy = dcl.FlattenStringSlice(m["requireAttestationsBy"])
	r.EnforcementMode = flattenPolicyAdmissionRuleEnforcementModeEnum(m["enforcementMode"])

	return r
}

// flattenPolicyAdmissionRuleEvaluationModeEnumMap flattens the contents of PolicyAdmissionRuleEvaluationModeEnum from a JSON
// response object.
func flattenPolicyAdmissionRuleEvaluationModeEnumMap(c *Client, i interface{}) map[string]PolicyAdmissionRuleEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyAdmissionRuleEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyAdmissionRuleEvaluationModeEnum{}
	}

	items := make(map[string]PolicyAdmissionRuleEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyAdmissionRuleEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyAdmissionRuleEvaluationModeEnumSlice flattens the contents of PolicyAdmissionRuleEvaluationModeEnum from a JSON
// response object.
func flattenPolicyAdmissionRuleEvaluationModeEnumSlice(c *Client, i interface{}) []PolicyAdmissionRuleEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyAdmissionRuleEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyAdmissionRuleEvaluationModeEnum{}
	}

	items := make([]PolicyAdmissionRuleEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyAdmissionRuleEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyAdmissionRuleEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyAdmissionRuleEvaluationModeEnum with the same value as that string.
func flattenPolicyAdmissionRuleEvaluationModeEnum(i interface{}) *PolicyAdmissionRuleEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyAdmissionRuleEvaluationModeEnumRef("")
	}

	return PolicyAdmissionRuleEvaluationModeEnumRef(s)
}

// flattenPolicyAdmissionRuleEnforcementModeEnumMap flattens the contents of PolicyAdmissionRuleEnforcementModeEnum from a JSON
// response object.
func flattenPolicyAdmissionRuleEnforcementModeEnumMap(c *Client, i interface{}) map[string]PolicyAdmissionRuleEnforcementModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyAdmissionRuleEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyAdmissionRuleEnforcementModeEnum{}
	}

	items := make(map[string]PolicyAdmissionRuleEnforcementModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyAdmissionRuleEnforcementModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyAdmissionRuleEnforcementModeEnumSlice flattens the contents of PolicyAdmissionRuleEnforcementModeEnum from a JSON
// response object.
func flattenPolicyAdmissionRuleEnforcementModeEnumSlice(c *Client, i interface{}) []PolicyAdmissionRuleEnforcementModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyAdmissionRuleEnforcementModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyAdmissionRuleEnforcementModeEnum{}
	}

	items := make([]PolicyAdmissionRuleEnforcementModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyAdmissionRuleEnforcementModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyAdmissionRuleEnforcementModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyAdmissionRuleEnforcementModeEnum with the same value as that string.
func flattenPolicyAdmissionRuleEnforcementModeEnum(i interface{}) *PolicyAdmissionRuleEnforcementModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyAdmissionRuleEnforcementModeEnumRef("")
	}

	return PolicyAdmissionRuleEnforcementModeEnumRef(s)
}

// flattenPolicyGlobalPolicyEvaluationModeEnumMap flattens the contents of PolicyGlobalPolicyEvaluationModeEnum from a JSON
// response object.
func flattenPolicyGlobalPolicyEvaluationModeEnumMap(c *Client, i interface{}) map[string]PolicyGlobalPolicyEvaluationModeEnum {
	a, ok := i.(map[string]interface{})
	if !ok {
		return map[string]PolicyGlobalPolicyEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return map[string]PolicyGlobalPolicyEvaluationModeEnum{}
	}

	items := make(map[string]PolicyGlobalPolicyEvaluationModeEnum)
	for k, item := range a {
		items[k] = *flattenPolicyGlobalPolicyEvaluationModeEnum(item.(interface{}))
	}

	return items
}

// flattenPolicyGlobalPolicyEvaluationModeEnumSlice flattens the contents of PolicyGlobalPolicyEvaluationModeEnum from a JSON
// response object.
func flattenPolicyGlobalPolicyEvaluationModeEnumSlice(c *Client, i interface{}) []PolicyGlobalPolicyEvaluationModeEnum {
	a, ok := i.([]interface{})
	if !ok {
		return []PolicyGlobalPolicyEvaluationModeEnum{}
	}

	if len(a) == 0 {
		return []PolicyGlobalPolicyEvaluationModeEnum{}
	}

	items := make([]PolicyGlobalPolicyEvaluationModeEnum, 0, len(a))
	for _, item := range a {
		items = append(items, *flattenPolicyGlobalPolicyEvaluationModeEnum(item.(interface{})))
	}

	return items
}

// flattenPolicyGlobalPolicyEvaluationModeEnum asserts that an interface is a string, and returns a
// pointer to a *PolicyGlobalPolicyEvaluationModeEnum with the same value as that string.
func flattenPolicyGlobalPolicyEvaluationModeEnum(i interface{}) *PolicyGlobalPolicyEvaluationModeEnum {
	s, ok := i.(string)
	if !ok {
		return PolicyGlobalPolicyEvaluationModeEnumRef("")
	}

	return PolicyGlobalPolicyEvaluationModeEnumRef(s)
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Policy) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalPolicy(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.urlNormalized()
		ncr := cr.urlNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		return true
	}
}

type policyDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         policyApiOperation
}

func convertFieldDiffsToPolicyDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]policyDiff, error) {
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
	var diffs []policyDiff
	// For each operation name, create a policyDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := policyDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToPolicyApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToPolicyApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (policyApiOperation, error) {
	switch opName {

	case "updatePolicyUpdatePolicyOperation":
		return &updatePolicyUpdatePolicyOperation{FieldDiffs: fieldDiffs}, nil

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}

func extractPolicyFields(r *Policy) error {
	// *PolicyAdmissionRule is a reused type - that's not compatible with function extractors.

	return nil
}
func extractPolicyAdmissionWhitelistPatternsFields(r *Policy, o *PolicyAdmissionWhitelistPatterns) error {
	return nil
}
func extractPolicyAdmissionRuleFields(r *Policy, o *PolicyAdmissionRule) error {
	return nil
}

func postReadExtractPolicyFields(r *Policy) error {
	// *PolicyAdmissionRule is a reused type - that's not compatible with function extractors.

	return nil
}
func postReadExtractPolicyAdmissionWhitelistPatternsFields(r *Policy, o *PolicyAdmissionWhitelistPatterns) error {
	return nil
}
func postReadExtractPolicyAdmissionRuleFields(r *Policy, o *PolicyAdmissionRule) error {
	return nil
}
