package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	nacosName = "118.89.85.75"
	nacosPort = 30655
)

func GetNacosConfig(namespaceId string, dataId, group string) (string, error) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serviceConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacosName,
			ContextPath: "/nacos",
			Port:        nacosPort,
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
		return "", err
	}

	return client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}
