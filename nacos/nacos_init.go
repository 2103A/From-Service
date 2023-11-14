package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosConfig struct {
	NacosClient config_client.IConfigClient
}

func CreateNacosClient(namespaceId string, logDir, cacheDir string) *NacosConfig {

	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              logDir,
		CacheDir:            cacheDir,
		LogLevel:            "debug",
	}

	serviceConfigs := []constant.ServerConfig{
		{
			IpAddr:      "118.89.85.75",
			ContextPath: "/nacos",
			Port:        30655,
			Scheme:      "http",
		},
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serviceConfigs,
		},
	)
	if err != nil {
		panic("get nacos client false" + err.Error())
		return nil
	}

	return &NacosConfig{
		NacosClient: client,
	}

}

func (n *NacosConfig) GetConfigInfo(dataId, group string) (string, error) {
	content, err := n.NacosClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return "", err
	}
	return content, nil
}
