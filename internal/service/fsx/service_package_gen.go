// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package fsx

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	fsx_sdkv1 "github.com/aws/aws-sdk-go/service/fsx"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceONTAPFileSystem,
			TypeName: "aws_fsx_ontap_file_system",
			Name:     "Ontap File System",
		},
		{
			Factory:  DataSourceOntapStorageVirtualMachine,
			TypeName: "aws_fsx_ontap_storage_virtual_machine",
			Name:     "ONTAP Storage Virtual Machine",
		},
		{
			Factory:  DataSourceOpenzfsSnapshot,
			TypeName: "aws_fsx_openzfs_snapshot",
		},
		{
			Factory:  DataSourceWindowsFileSystem,
			TypeName: "aws_fsx_windows_file_system",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBackup,
			TypeName: "aws_fsx_backup",
			Name:     "Backup",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDataRepositoryAssociation,
			TypeName: "aws_fsx_data_repository_association",
			Name:     "Data Repository Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceFileCache,
			TypeName: "aws_fsx_file_cache",
			Name:     "File Cache",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceLustreFileSystem,
			TypeName: "aws_fsx_lustre_file_system",
			Name:     "Lustre File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceONTAPFileSystem,
			TypeName: "aws_fsx_ontap_file_system",
			Name:     "ONTAP File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceONTAPStorageVirtualMachine,
			TypeName: "aws_fsx_ontap_storage_virtual_machine",
			Name:     "ONTAP Storage Virtual Machine",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceOntapVolume,
			TypeName: "aws_fsx_ontap_volume",
			Name:     "ONTAP Volume",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceOpenZFSFileSystem,
			TypeName: "aws_fsx_openzfs_file_system",
			Name:     "OpenZFS File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceOpenzfsSnapshot,
			TypeName: "aws_fsx_openzfs_snapshot",
			Name:     "OpenZFS Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceOpenzfsVolume,
			TypeName: "aws_fsx_openzfs_volume",
			Name:     "OpenZFS Volume",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceWindowsFileSystem,
			TypeName: "aws_fsx_windows_file_system",
			Name:     "Windows File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.FSx
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*fsx_sdkv1.FSx, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return fsx_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
