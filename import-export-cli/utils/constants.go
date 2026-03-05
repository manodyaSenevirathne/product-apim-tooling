/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package utils

import (
	"os"
	"os/user"
	"path/filepath"
)

const ProjectName = "apictl"

var MICmd = "apictl"

func GetMICmdName() string {
	if MICmd == "mi" {
		return ""
	}
	envProjName := os.Getenv("MICmd")
	if envProjName == "mi" {
		MICmd = envProjName
		return ""
	}
	return MICmd
}

// File Names and Paths
var CurrentDir, _ = os.Getwd()

const ConfigDirName = ".wso2apictl"

const MIConfigDirName = ".wso2mi"

var HomeDirectory = getConfigHomeDir()

func getConfigHomeDir() string {
	value := os.Getenv("APICTL_CONFIG_DIR")
	if len(value) == 0 {
		value, err := os.UserHomeDir()
		if len(value) == 0 || err != nil {
			current, err := user.Current()
			if err != nil || current == nil {
				HandleErrorAndExit("User's HOME folder location couldn't be identified", nil)
				return ""
			}
			return current.HomeDir
		}
		return value
	}
	return value
}

func GetConfigDirPath() string {
	if MICmd == "mi" {
		return filepath.Join(HomeDirectory, MIConfigDirName)
	}
	return filepath.Join(HomeDirectory, ConfigDirName)
}

func getLocalCredentialsDirectoryName() string {
	if MICmd == "mi" {
		return filepath.Join(HomeDirectory, MILocalCredentialsDirectoryName)
	}
	return filepath.Join(HomeDirectory, LocalCredentialsDirectoryName)
}

var ConfigDirPath = filepath.Join(HomeDirectory, ConfigDirName)

const LocalCredentialsDirectoryName = ".wso2apictl.local"
const MILocalCredentialsDirectoryName = ".wso2mi.local"
const EnvKeysAllFileName = "env_keys_all.yaml"
const MainConfigFileName = "main_config.yaml"
const SampleMainConfigFileName = "main_config.yaml.sample"
const DefaultAPISpecFileName = "default_api.yaml"

var LocalCredentialsDirectoryPath = getLocalCredentialsDirectoryName()
var EnvKeysAllFilePath = filepath.Join(LocalCredentialsDirectoryPath, EnvKeysAllFileName)
var MainConfigFilePath = filepath.Join(GetConfigDirPath(), MainConfigFileName)
var SampleMainConfigFilePath = filepath.Join(ConfigDirPath, SampleMainConfigFileName)
var DefaultAPISpecFilePath = filepath.Join(ConfigDirPath, DefaultAPISpecFileName)

const DefaultExportDirName = "exported"
const ExportedApisDirName = "apis"
const ExportedMCPServersDirName = "mcp-servers"
const ExportedPoliciesDirName = "policies"
const ExportedThrottlePoliciesDirName = "rate-limiting"
const ExportedAPIPoliciesDirName = "api"
const ExportedApiProductsDirName = "api-products"
const ExportedAppsDirName = "apps"
const ExportedMigrationArtifactsDirName = "migration"
const CertificatesDirName = "certs"

const (
	InitProjectDefinitions              = "Definitions"
	InitProjectDefinitionsSwagger       = InitProjectDefinitions + string(os.PathSeparator) + "swagger.yaml"
	InitProjectDefinitionsGraphQLSchema = InitProjectDefinitions + string(os.PathSeparator) + "schema.graphql"
	InitProjectDefinitionsAsyncAPI      = InitProjectDefinitions + string(os.PathSeparator) + "asyncapi.yaml"
	InitProjectImage                    = "Image"
	InitProjectDocs                     = "Docs"
	InitProjectSequences                = "Policies"
	InitProjectClientCertificates       = "Client-certificates"
	InitProjectEndpointCertificates     = "Endpoint-certificates"
	InitProjectInterceptors             = "Interceptors"
	InitProjectLibs                     = "libs"
	InitProjectWSDL                     = "WSDL"
)

const DeploymentDirPrefix = "DeploymentArtifacts_"
const DeploymentCertificatesDirectory = "certificates"

var DefaultExportDirPath = filepath.Join(GetConfigDirPath(), DefaultExportDirName)
var DefaultCertDirPath = filepath.Join(ConfigDirPath, CertificatesDirName)

const defaultApiApplicationImportExportSuffix = "api/am/admin/v4"
const defaultPublisherApiImportExportSuffix = "api/am/publisher/v4"
const defaultApiListEndpointSuffix = "api/am/publisher/v4/apis"
const defaultMcpServerListEndpointSuffix = "api/am/publisher/v4/mcp-servers"
const defaultAPIPolicyListEndpointSuffix = "api/am/publisher/v4/operation-policies"
const defaultApiProductListEndpointSuffix = "api/am/publisher/v4/api-products"
const defaultUnifiedSearchEndpointSuffix = "api/am/publisher/v4/search"
const defaultAdminApplicationListEndpointSuffix = "api/am/admin/v4/applications"
const defaultDevPortalApplicationListEndpointSuffix = "api/am/devportal/v3/applications"
const defaultDevPortalThrottlingPoliciesEndpointSuffix = "api/am/devportal/v3/throttling-policies"
const defaultClientRegistrationEndpointSuffix = "client-registration/v0.17/register"
const defaultTokenEndPoint = "oauth2/token"
const defaultRevokeEndpointSuffix = "oauth2/revoke"
const defaultAPILoggingBaseEndpoint = "api/am/devops/v0/tenant-logs"
const defaultAPILoggingApisEndpoint = "apis"
const defaultCorrelationLoggingEndpoint = "api/am/devops/v0/config/correlation"
const defaultAIServiceEndpoint = "https://dev-tools.wso2.com/apim-ai-service/v2"
const defaultAITokenServiceEndpoint = "https://api.asgardeo.io/t/wso2devtools/oauth2/token"

const DefaultEnvironmentName = "default"
const DefaultTenantDomain = "carbon.super"

// API Product related constants
const DefaultApiProductVersion = "1.0.0"
const DefaultApiProductType = "APIProduct"

// MCP Server related constants
const DefaultMcpServerType = "MCP"

// Application keys related constants
const ProductionKeyType = "PRODUCTION"
const SandboxKeyType = "SANDBOX"

var GrantTypesToBeSupported = []string{"refresh_token", "password", "client_credentials"}

// WSO2PublicCertificate : wso2 public certificate in PEM format
var WSO2PublicCertificate = []byte{45,45,45,45,45,66,69,71,73,78,32,67,69,82,84,73,70,73,67,65,84,69,45,45,45,45,45,10,77,73,73,68,117,84,67,67,65,113,71,103,65,119,73,66,65,103,73,85,74,81,73,90,101,119,112,89,114,105,105,53,113,57,112,89,121,55,47,90,88,75,120,113,47,97,99,119,68,81,89,74,75,111,90,73,104,118,99,78,65,81,69,76,13,10,66,81,65,119,90,68,69,76,77,65,107,71,65,49,85,69,66,104,77,67,86,86,77,120,67,122,65,74,66,103,78,86,66,65,103,77,65,107,78,66,77,82,89,119,70,65,89,68,86,81,81,72,68,65,49,78,98,51,86,117,100,71,70,112,13,10,98,105,66,87,97,87,86,51,77,81,48,119,67,119,89,68,86,81,81,75,68,65,82,88,85,48,56,121,77,81,48,119,67,119,89,68,86,81,81,76,68,65,82,88,85,48,56,121,77,82,73,119,69,65,89,68,86,81,81,68,68,65,108,115,13,10,98,50,78,104,98,71,104,118,99,51,81,119,72,104,99,78,77,106,89,119,77,122,65,48,77,68,103,119,77,68,73,48,87,104,99,78,77,106,103,119,78,106,65,50,77,68,103,119,77,68,73,48,87,106,66,107,77,81,115,119,67,81,89,68,13,10,86,81,81,71,69,119,74,86,85,122,69,76,77,65,107,71,65,49,85,69,67,65,119,67,81,48,69,120,70,106,65,85,66,103,78,86,66,65,99,77,68,85,49,118,100,87,53,48,89,87,108,117,73,70,90,112,90,88,99,120,68,84,65,76,13,10,66,103,78,86,66,65,111,77,66,70,100,84,84,122,73,120,68,84,65,76,66,103,78,86,66,65,115,77,66,70,100,84,84,122,73,120,69,106,65,81,66,103,78,86,66,65,77,77,67,87,120,118,89,50,70,115,97,71,57,122,100,68,67,67,13,10,65,83,73,119,68,81,89,74,75,111,90,73,104,118,99,78,65,81,69,66,66,81,65,68,103,103,69,80,65,68,67,67,65,81,111,67,103,103,69,66,65,76,74,122,47,65,70,102,69,73,82,53,53,107,112,86,68,90,56,113,50,48,75,83,13,10,98,98,83,102,113,112,119,121,101,43,82,69,115,57,102,89,47,109,105,90,49,109,99,76,104,119,74,114,87,90,47,89,53,81,43,87,86,108,73,109,101,66,66,102,80,107,100,79,51,119,66,106,56,111,88,101,101,76,84,68,90,105,66,53,13,10,66,55,100,52,114,78,43,85,110,114,78,117,71,103,84,72,67,89,118,101,67,79,66,121,53,71,82,47,97,49,86,105,99,101,99,102,54,81,53,101,49,106,73,86,87,118,87,108,111,89,119,107,99,89,74,102,49,70,106,122,114,106,89,116,13,10,121,85,114,55,76,70,87,78,73,50,119,66,122,100,117,108,105,83,111,82,105,82,43,121,114,65,77,111,84,43,87,113,81,84,111,52,47,117,102,49,67,43,110,105,122,121,101,80,122,113,43,102,113,77,121,75,48,109,83,69,83,103,66,110,13,10,102,69,80,106,76,90,54,72,87,66,82,102,120,103,70,79,57,54,122,79,73,87,70,66,101,66,55,74,88,97,74,118,90,109,106,65,112,73,88,97,101,43,80,89,56,72,86,121,80,56,104,109,85,56,66,51,97,47,113,120,108,117,90,53,13,10,70,120,78,84,55,76,85,73,69,74,118,115,106,120,112,71,97,80,75,84,68,120,76,80,100,108,104,105,109,68,101,111,75,104,114,116,79,48,114,110,69,67,47,117,82,108,76,56,88,55,106,98,116,50,106,66,70,43,105,101,89,104,107,67,13,10,65,119,69,65,65,97,78,106,77,71,69,119,70,65,89,68,86,82,48,82,66,65,48,119,67,52,73,74,98,71,57,106,89,87,120,111,98,51,78,48,77,65,115,71,65,49,85,100,68,119,81,69,65,119,73,69,56,68,65,100,66,103,78,86,13,10,72,83,85,69,70,106,65,85,66,103,103,114,66,103,69,70,66,81,99,68,65,81,89,73,75,119,89,66,66,81,85,72,65,119,73,119,72,81,89,68,86,82,48,79,66,66,89,69,70,77,97,74,109,69,117,100,52,116,87,53,47,50,76,50,13,10,98,55,107,85,80,97,97,105,86,113,78,104,77,65,48,71,67,83,113,71,83,73,98,51,68,81,69,66,67,119,85,65,65,52,73,66,65,81,66,121,104,108,110,89,90,113,98,99,122,85,73,99,47,43,114,76,76,43,72,101,120,107,115,53,13,10,118,49,88,68,67,47,85,97,48,88,114,83,56,70,113,100,78,80,104,97,121,84,114,51,84,74,67,76,116,75,66,120,104,48,108,122,65,114,118,54,84,116,68,47,65,106,84,109,73,76,86,101,118,107,49,119,101,120,111,48,100,48,103,71,13,10,105,84,121,88,104,80,113,87,79,65,82,106,71,106,111,118,99,67,121,51,122,109,109,117,114,57,81,90,43,53,107,73,66,77,71,83,105,117,98,76,116,119,65,118,82,108,72,98,122,72,57,104,47,73,108,90,120,107,85,57,76,121,57,47,13,10,118,69,50,50,47,104,88,52,74,106,87,121,65,112,82,76,51,90,47,55,67,54,76,82,84,97,110,78,109,52,79,66,71,73,100,69,70,98,113,98,109,114,79,105,68,52,57,102,76,115,121,80,70,89,78,88,78,82,54,98,116,51,112,52,13,10,88,72,109,100,75,114,77,117,77,101,108,80,81,87,105,74,55,50,122,75,83,120,107,54,86,98,90,120,81,69,119,105,83,101,81,86,113,71,52,70,112,100,84,72,99,71,81,84,54,74,99,54,49,82,77,85,86,80,74,100,73,119,88,86,13,10,85,117,122,97,78,68,122,43,102,48,78,99,102,56,67,86,101,67,68,84,87,90,118,83,90,86,81,87,77,50,90,104,87,119,102,119,57,43,66,89,114,43,79,88,50,78,122,101,52,108,75,105,102,108,54,103,67,51,51,122,10,45,45,45,45,45,69,78,68,32,67,69,82,84,73,70,73,67,65,84,69,45,45,45,45,45,10}

// Headers and Header Values
const HeaderAuthorization = "Authorization"
const HeaderContentType = "Content-Type"
const HeaderConnection = "Connection"
const HeaderAccept = "Accept"
const HeaderProduces = "Produces"
const HeaderConsumes = "Consumes"
const HeaderContentEncoding = "Content-Encoding"
const HeaderTransferEncoding = "transfer-encoding"
const HeaderValueChunked = "chunked"
const HeaderValueGZIP = "gzip"
const HeaderValueKeepAlive = "keep-alive"
const HeaderValueApplicationZip = "application/zip"
const HeaderValueApplicationJSON = "application/json"
const HeaderValueXWWWFormUrlEncoded = "application/x-www-form-urlencoded"
const HeaderValueAuthBearerPrefix = "Bearer"
const HeaderValueAuthBasicPrefix = "Basic"
const HeaderValueMultiPartFormData = "multipart/form-data"
const HeaderToken = "token="
const TokenTypeForRevocation = "&token_type_hint=access_token"

// Logging Prefixes
const LogPrefixInfo = "[INFO]: "
const LogPrefixWarning = "[WARN]: "
const LogPrefixError = "[ERROR]: "

// String Constants
const SearchAndTag = "&"

// Other
const DefaultTokenValidityPeriod = 3600
const DefaultHttpRequestTimeout = 10000

// AI
const DefaultAIThreadCount = 3
const DefaultAIEndpoint = "https://e95488c8-8511-4882-967f-ec3ae2a0f86f-prod.e1-us-east-azure.choreoapis.dev/lgpt/interceptor-service/interceptor-service-be2/v1.0"

// TLSRenegotiationNever : never negotiate
const TLSRenegotiationNever = "never"

// TLSRenegotiationOnce : negotiate once
const TLSRenegotiationOnce = "once"

// TLSRenegotiationFreely : negotiate freely
const TLSRenegotiationFreely = "freely"

// Migration export
const MaxAPIsToExportOnce = 20
const MaxAppsToExportOnce = 20
const MaxMCPServersToExportOnce = 20
const MigrationAPIsExportMetadataFileName = "migration-apis-export-metadata.yaml"
const MigrationAppsExportMetadataFileName = "migration-apps-export-metadata.yaml"
const MigrationMCPServersExportMetadataFileName = "migration-mcp-servers-export-metadata.yaml"
const LastSucceededApiFileName = "last-succeeded-api.log"
const LastSucceededAppFileName = "last-succeeded-app.log"
const LastSucceededMCPServerFileName = "last_succeeded_mcp_server.log"
const LastSuceededContentDelimiter = " " // space
const DefaultResourceTenantDomain = "tenant-default"
const ApplicationId = "applicationId"
const ApiId = "apiId"
const APIProductId = "apiProductId"
const DefaultCliApp = "default-apictl-app"
const DefaultTokenType = "JWT"

const LifeCycleAction = "action"

var ValidInitialStates = []string{"CREATED", "PUBLISHED"}

// The list of repos and directories that can be used when replcing env variables
var EnvReplaceFilePaths = []string{
	"Policies",
}

// The list of file extensions when replcing env variables related to Policies
var EnvReplacePoliciesFileExtensions = []string{
	"j2",
	"gotmpl",
}

// project types
const (
	ProjectTypeNone        = "None"
	ProjectTypeApi         = "API"
	ProjectTypeMcpServer   = "MCP Server"
	ProjectTypeApiProduct  = "API Product"
	ProjectTypeApplication = "Application"
	ProjectTypeRevision    = "Revision"
	ProjectTypePolicy      = "Policy"
	ProjectTypeAPIPolicy   = "API Policy"
)

// project param files
const ParamFile = "params.yaml"
const ParamsIntermediateFile = "intermediate_params.yaml"

const (
	APIDefinitionFileYaml         = "api.yaml"
	APIDefinitionFileJson         = "api.json"
	MCPServerDefinitionFileYaml   = "mcp_server.yaml"
	MCPServerDefinitionFileJson   = "mcp_server.json"
	APIProductDefinitionFileYaml  = "api_product.yaml"
	APIProductDefinitionFileJson  = "api_product.json"
	ApplicationDefinitionFileYaml = "application.yaml"
	ApplicationDefinitionFileJson = "application.json"
)

// project meta files
const (
	MetaFileAPI         = "api_meta.yaml"
	MetaFileMCPServer   = "mcp_server_meta.yaml"
	MetaFileAPIProduct  = "api_product_meta.yaml"
	MetaFileApplication = "application_meta.yaml"
)

// Constants related to meta file structs
const DeployImportRotateRevision = "deploy.import.rotateRevision"
const DeployImportSkipSubscriptions = "deploy.import.skipSubscriptions"

const DeploymentEnvFile = "deployment_environments.yaml"
const PrivateJetModeConst = "privateJet"
const SidecarModeConst = "sidecar"

// Default values for Help commands
const DefaultApisDisplayLimit = 25
const DefaultApiProductsDisplayLimit = 25
const DefaultAppsDisplayLimit = 25
const DefaultExportFormat = "YAML"
const DefaultPoliciesDisplayLimit = 25

const InitDirName = string(os.PathSeparator) + "init" + string(os.PathSeparator)

// AWS API security document constants
const DefaultAWSDocFileName = "document.yaml"

const ResourcePolicyDocName = "resource_policy_doc"
const ResourcePolicyDocDisplayName = "Resource Policy"
const ResourcePolicyDocSummary = "This document contains details related to AWS resource policies"

const CognitoUserPoolDocName = "cognito_userpool_doc"
const CognitoDocDisplayName = "Cognito Userpool"
const CognitoDocSummary = "This document contains details related to AWS cognito user pools"

const AWSAPIKeyDocName = "aws_apikey_doc"
const ApiKeysDocDisplayName = "AWS APIKeys"
const ApiKeysDocSummary = "This document contains details related to AWS API keys"

const AWSSigV4DocName = "aws_sigv4_doc"
const AWSSigV4DocDisplayName = "AWS Signature Version4"
const AWSSigV4DocSummary = "This document contains details related to AWS signature version 4"

// MiCmdLiteral denote the alias for micro integrator related commands
const MiCmdLiteral = "mi"

// MiManagementAPIContext
const MiManagementAPIContext = "management"

// Mi Management Resource paths
const MiManagementCarbonAppResource = "applications"
const MiManagementServiceResource = "services"
const MiManagementAPIResource = "apis"
const MiManagementProxyServiceResource = "proxy-services"
const MiManagementInboundEndpointResource = "inbound-endpoints"
const MiManagementEndpointResource = "endpoints"
const MiManagementMessageProcessorResource = "message-processors"
const MiManagementTemplateResource = "templates"
const MiManagementConnectorResource = "connectors"
const MiManagementMessageStoreResource = "message-stores"
const MiManagementLocalEntrieResource = "local-entries"
const MiManagementSequenceResource = "sequences"
const MiManagementTaskResource = "tasks"
const MiManagementLogResource = "logs"
const MiManagementLoggingResource = "logging"
const MiManagementServerResource = "server"
const MiManagementDataServiceResource = "data-services"
const MiManagementMiLoginResource = "login"
const MiManagementMiLogoutResource = "logout"
const MiManagementUserResource = "users"
const MiManagementTransactionResource = "transactions"
const MiManagementTransactionCountResource = "count"
const MiManagementTransactionReportResource = "report"
const MiManagementExternalVaultsResource = "external-vaults"
const MiManagementExternalVaultHashiCorpResource = "hashicorp"
const MiManagementRoleResource = "roles"

const ZipFileSuffix = ".zip"

// Output format types
const JsonArrayFormatType = "jsonArray"

const ThrottlingPolicyTypeSub = "subscription"
const ThrottlingPolicyTypeApp = "application"
const ThrottlingPolicyTypeAdv = "advanced"
const ThrottlingPolicyTypeCus = "custom"
