package configs

type ApifoxOverwriteBehavior string

const (
	// 覆盖现有模式
	APIFOX_OVERWRITE_BEHAVIOR_OVERWRITE ApifoxOverwriteBehavior = "OVERWRITE_EXISTING"
	// 自动合并更改到现有的
	APIFOX_OVERWRITE_BEHAVIOR_AUTO_MERGE ApifoxOverwriteBehavior = "AUTO_MERGE"
	// 跳过更改并保留现有的
	APIFOX_OVERWRITE_BEHAVIOR_SKIP ApifoxOverwriteBehavior = "KEEP_EXISTING"
	// 保留现有，创建新的
	APIFOX_OVERWRITE_BEHAVIOR_CREATE ApifoxOverwriteBehavior = "CREATE_NEW"

	DefaultFilePath = "./docs/swagger.json"
)

/*
接口文档配置

locale: 接口文档语言, 默认值: zh-CN, 枚举值: zh-CN, en-US

project-id: 项目ID

api-version: 接口文档版本, 枚举值: 2024-03-28

token: 接口文档token

target-endpoint-folder-id: 接口文档文件夹ID

target-schema-folder-id: 接口文档文件夹ID

endpoint-overwrite-behavior: 接口文档覆盖行为, 默认值: OVERWRITE_EXISTING, 枚举值: OVERWRITE_EXISTING, AUTO_MERGE, KEEP_EXISTING, CREATE_NEW

schema-overwrite-behavior: 接口文档覆盖行为, 默认值: OVERWRITE_EXISTING, 枚举值: OVERWRITE_EXISTING, AUTO_MERGE, KEEP_EXISTING, CREATE_NEW

update-folder-of-changed-endpoint: 是否更新更改的接口的文件夹, 默认设置为 false

prepend-base-path: 是否将基础路径添加到接口的路径中，默认设置为 false。我们建议将其设置为 false，这样基础路径可以保留在“环境面板”中，而不是每个接口内部。如果希望在接口路径中添加路径前缀，则应将其设置为 true。
*/
type ApifoxConfig struct {
	Locale                        string                  `mapstructure:"locale" json:"locale" yaml:"locale" toml:"locale"`
	ProjectId                     string                  `mapstructure:"project-id" json:"project-id" yaml:"project-id" toml:"project-id"`
	Token                         string                  `mapstructure:"token" json:"token" yaml:"token" toml:"token"`
	ApiVersion                    string                  `mapstructure:"api-version" json:"api-version" yaml:"api-version" toml:"api-version"`
	TargetEndpointFolderId        string                  `mapstructure:"target-endpoint-folder-id" json:"target-endpoint-folder-id" yaml:"target-endpoint-folder-id" toml:"target-endpoint-folder-id"`
	TargetSchemaFolderId          string                  `mapstructure:"target-schema-folder-id" json:"target-schema-folder-id" yaml:"target-schema-folder-id" toml:"target-schema-folder-id"`
	EndpointOverwriteBehavior     ApifoxOverwriteBehavior `mapstructure:"endpoint-overwrite-behavior" json:"endpoint-overwrite-behavior" yaml:"endpoint-overwrite-behavior" toml:"endpoint-overwrite-behavior"`
	SchemaOverwriteBehavior       ApifoxOverwriteBehavior `mapstructure:"schema-overwrite-behavior" json:"schema-overwrite-behavior" yaml:"schema-overwrite-behavior" toml:"schema-overwrite-behavior"`
	UpdateFolderOfChangedEndpoint bool                    `mapstructure:"update-folder-of-changed-endpoint" json:"update-folder-of-changed-endpoint" yaml:"update-folder-of-changed-endpoint" toml:"update-folder-of-changed-endpoint"`
	PrependBasePath               bool                    `mapstructure:"prepend-base-path" json:"prepend-base-path" yaml:"prepend-base-path" toml:"prepend-base-path"`
	FilePath                      string                  `mapstructure:"file-path" json:"file-path" yaml:"file-path" toml:"file-path"`
}
