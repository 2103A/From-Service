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

//func ServiceRegister(namespaceId string, ip, serviceName, groupName string, port uint64) {
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         namespaceId,
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "tmp/nacos/log",
//		CacheDir:            "tmp/nacos/cache",
//		LogLevel:            "debug",
//	}
//
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr:      nacosName,
//			ContextPath: "/nacos",
//			Port:        nacosPort,
//			Scheme:      "http",
//		},
//	}
//
//	// 创建服务发现客户端的另一种方式 (推荐)
//	namingClient, err := clients.NewNamingClient(
//		vo.NacosClientParam{
//			ClientConfig:  &clientConfig,
//			ServerConfigs: serverConfigs,
//		},
//	)
//	if err != nil {
//		panic("创建服务发现客户端失败" + err.Error())
//	}
//
//	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
//		Ip:          ip,
//		Port:        port,
//		ServiceName: serviceName,
//		Weight:      10,
//		Enable:      true,
//		Healthy:     true,
//		Ephemeral:   true,
//		Metadata:    map[string]string{"idc": "2103A"},
//		GroupName:   groupName, // 默认值DEFAULT_GROUP
//	})
//	if !success {
//		panic("服务注册失败")
//	}
//}
//
//func ServiceDiscovery() {
//	// SelectInstances 只返回满足这些条件的实例列表：healthy=${HealthyOnly},enable=true 和weight>0
//	instances, err := namingClient.SelectInstances(vo.SelectInstancesParam{
//		ServiceName: "demo.go",
//		GroupName:   "group-a",             // 默认值DEFAULT_GROUP
//		Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
//		HealthyOnly: true,
//	})
//
//}
