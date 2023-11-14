package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type NacosClient struct {
	Client config_client.IConfigClient
}

func CreateNacosClient(namespaceId string, logDir, cacheDir string) (*NacosClient, error) {

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
			IpAddr:      "122.51.59.109",
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

	return &NacosClient{
		Client: client,
	}, err
}

func (c *NacosClient) GetConfig(dataId, group string) (string, error) {
	return c.Client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}
