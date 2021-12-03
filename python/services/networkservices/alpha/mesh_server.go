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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/alpha/networkservices_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
)

// Server implements the gRPC interface for Mesh.
type MeshServer struct{}

// ProtoToMeshTypeEnum converts a MeshTypeEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaMeshTypeEnum(e alphapb.NetworkservicesAlphaMeshTypeEnum) *alpha.MeshTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaMeshTypeEnum_name[int32(e)]; ok {
		e := alpha.MeshTypeEnum(n[len("NetworkservicesAlphaMeshTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMesh converts a Mesh resource from its proto representation.
func ProtoToMesh(p *alphapb.NetworkservicesAlphaMesh) *alpha.Mesh {
	obj := &alpha.Mesh{
		Name:             dcl.StringOrNil(p.GetName()),
		CreateTime:       dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:       dcl.StringOrNil(p.GetUpdateTime()),
		Description:      dcl.StringOrNil(p.GetDescription()),
		Type:             ProtoToNetworkservicesAlphaMeshTypeEnum(p.GetType()),
		Scope:            dcl.StringOrNil(p.GetScope()),
		InterceptionPort: dcl.Int64OrNil(p.GetInterceptionPort()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Location:         dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// MeshTypeEnumToProto converts a MeshTypeEnum enum to its proto representation.
func NetworkservicesAlphaMeshTypeEnumToProto(e *alpha.MeshTypeEnum) alphapb.NetworkservicesAlphaMeshTypeEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaMeshTypeEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaMeshTypeEnum_value["MeshTypeEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaMeshTypeEnum(v)
	}
	return alphapb.NetworkservicesAlphaMeshTypeEnum(0)
}

// MeshToProto converts a Mesh resource to its proto representation.
func MeshToProto(resource *alpha.Mesh) *alphapb.NetworkservicesAlphaMesh {
	p := &alphapb.NetworkservicesAlphaMesh{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetType(NetworkservicesAlphaMeshTypeEnumToProto(resource.Type))
	p.SetScope(dcl.ValueOrEmptyString(resource.Scope))
	p.SetInterceptionPort(dcl.ValueOrEmptyInt64(resource.InterceptionPort))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyMesh handles the gRPC request by passing it to the underlying Mesh Apply() method.
func (s *MeshServer) applyMesh(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaMeshRequest) (*alphapb.NetworkservicesAlphaMesh, error) {
	p := ProtoToMesh(request.GetResource())
	res, err := c.ApplyMesh(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MeshToProto(res)
	return r, nil
}

// applyNetworkservicesAlphaMesh handles the gRPC request by passing it to the underlying Mesh Apply() method.
func (s *MeshServer) ApplyNetworkservicesAlphaMesh(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaMeshRequest) (*alphapb.NetworkservicesAlphaMesh, error) {
	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMesh(ctx, cl, request)
}

// DeleteMesh handles the gRPC request by passing it to the underlying Mesh Delete() method.
func (s *MeshServer) DeleteNetworkservicesAlphaMesh(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaMeshRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMesh(ctx, ProtoToMesh(request.GetResource()))

}

// ListNetworkservicesAlphaMesh handles the gRPC request by passing it to the underlying MeshList() method.
func (s *MeshServer) ListNetworkservicesAlphaMesh(ctx context.Context, request *alphapb.ListNetworkservicesAlphaMeshRequest) (*alphapb.ListNetworkservicesAlphaMeshResponse, error) {
	cl, err := createConfigMesh(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMesh(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaMesh
	for _, r := range resources.Items {
		rp := MeshToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkservicesAlphaMeshResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMesh(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}