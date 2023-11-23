package service

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	HOST = "118.89.85.75"
	PORT = 30655
)

type NacosServiceClient struct {
	NaMingClient naming_client.INamingClient
}

func NewNacosServiceClient(nameSpaceId string) (*NacosServiceClient, error) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         nameSpaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      HOST,
			ContextPath: "/nacos",
			Port:        PORT,
			Scheme:      "http",
		},
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return nil, err
	}
	return &NacosServiceClient{
		NaMingClient: namingClient,
	}, nil
}

func (n *NacosServiceClient) ServiceRegister(ip, serviceName, groupName string, port uint64) error {
	success, err := n.NaMingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "2103A"},
		GroupName:   groupName, // 默认值DEFAULT_GROUP
	})
	if !success && err != nil {
		return err
	}
	return nil
}

func (n *NacosServiceClient) ServiceDiscovery(serviceName, groupName string) ([]model.Instance, error) {
	// SelectInstances 只返回满足这些条件的实例列表：healthy=${HealthyOnly},enable=true 和weight>0
	instances, err := n.NaMingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serviceName,
		GroupName:   groupName,             // 默认值DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		HealthyOnly: true,
	})

	if err != nil {
		return nil, err
	}

	return instances, nil

}
