package handler

import (
	"fmt"
	"net"
)

type HostIPHandler interface {
	GetHostIP() (string, error)
}
type IPv6 struct {
	interfaceName string
}

func NewIPv6(interfaceName string) *IPv6 {
	return &IPv6{interfaceName}
}

func (i *IPv6) GetHostIP() (string, error) {
	var ipv6Addresses string

	// 获取指定的网络接口
	iface, err := net.InterfaceByName(i.interfaceName)
	if err != nil {
		return "", fmt.Errorf("could not find interface %s: %v", i.interfaceName, err)
	}

	// 获取网络接口的所有地址
	addrs, err := iface.Addrs()
	if err != nil {
		return "", fmt.Errorf("could not get addresses for interface %s: %v", i.interfaceName, err)
	}

	for _, addr := range addrs {

		ipNet, ok := addr.(*net.IPNet)
		if ok && ipNet.IP.To4() == nil && ipNet.IP.To16() != nil && ipNet.IP.IsGlobalUnicast() {
			// 如果是 IPv6 地址且为全局单播地址，添加到结果中
			ipv6Addresses = ipNet.IP.String()
			break
		}
	}

	return ipv6Addresses, nil
}
