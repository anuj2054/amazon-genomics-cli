package context

import (
	"testing"

	"github.com/aws/amazon-genomics-cli/internal/pkg/aws/cdk"
	"github.com/aws/amazon-genomics-cli/internal/pkg/cli/spec"
	awsmocks "github.com/aws/amazon-genomics-cli/internal/pkg/mocks/aws"
	storagemocks "github.com/aws/amazon-genomics-cli/internal/pkg/mocks/storage"
	"github.com/golang/mock/gomock"
)

const (
	testContextName1       = "testContextName1"
	testContextName2       = "testContextName2"
	testContextName3       = "testContextName3"
	testUnknownContextName = "unknown-context-name"
	testS3Location1        = "s3://test-s3-location-1"
	testS3Location2        = "s3://test-s3-location-2"
	testProjectName        = "testProjectName"
	testWesUrl             = "test-wes-url"
	testLogGroupName       = "test-log-group-name"
	testOutputBucket       = "test-output-bucket"
	testArtifactBucket     = "test-artifact-bucket"
	testHomeDir            = "test-home-dir"
	testUserEmail          = "bender@amazon.com"
	testUserId             = "bender123"
)

var (
	testValidProjectSpec = spec.Project{
		Name: testProjectName,
		Data: []spec.Data{{Location: testS3Location1}, {Location: testS3Location2, ReadOnly: true}},
		Contexts: map[string]spec.Context{
			testContextName1: {
				Engines: []spec.Engine{
					{Type: "wdl", Engine: "cromwell"},
				},
			},
			testContextName2: {
				Engines: []spec.Engine{
					{Type: "wdl", Engine: "cromwell"},
				},
			},
			testContextName3: {
				Engines: []spec.Engine{
					{Type: "wdl", Engine: "cromwell"},
				},
			},
		},
	}
)

type mockClients struct {
	ctrl            *gomock.Controller
	cdkMock         *awsmocks.MockCdkClient
	projMock        *storagemocks.MockProjectClient
	cfnMock         *awsmocks.MockCfnClient
	ssmMock         *awsmocks.MockSsmClient
	configMock      *storagemocks.MockConfigClient
	progressStream1 cdk.ProgressStream
	progressStream2 cdk.ProgressStream
}

func createMocks(t *testing.T) mockClients {
	ctrl := gomock.NewController(t)

	return mockClients{
		ctrl:            ctrl,
		cdkMock:         awsmocks.NewMockCdkClient(ctrl),
		projMock:        storagemocks.NewMockProjectClient(ctrl),
		cfnMock:         awsmocks.NewMockCfnClient(ctrl),
		ssmMock:         awsmocks.NewMockSsmClient(ctrl),
		configMock:      storagemocks.NewMockConfigClient(ctrl),
		progressStream1: make(cdk.ProgressStream),
		progressStream2: make(cdk.ProgressStream),
	}
}
