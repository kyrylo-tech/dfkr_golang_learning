package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ipToUint32(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip.To4())
}

func uint32ToIP(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}

func maskStringToPrefix(mask string) (int, error) {
	ip := net.ParseIP(mask)
	if ip == nil {
		return 0, fmt.Errorf("invalid mask")
	}
	m := net.IPMask(ip.To4())
	ones, _ := m.Size()
	return ones, nil
}

func main() {
	networkFlag := flag.String("network", "", "CIDR network (e.g. 10.0.0.0/22)")
	subnetMaskFlag := flag.String("subnetMask", "", "new subnet mask (/num or 255.255.255.0)")
	flag.Parse()

	var networkCIDR string
	var newMaskStr string

	args := flag.Args()
	if *networkFlag != "" {
		networkCIDR = *networkFlag
	} else if len(args) > 0 {
		networkCIDR = args[0]
	}

	if *subnetMaskFlag != "" {
		newMaskStr = *subnetMaskFlag
	} else if len(args) > 1 {
		newMaskStr = args[1]
	}

	if networkCIDR == "" || newMaskStr == "" {
		fmt.Println("Usage:")
		fmt.Println("calc 192.168.1.0/28 30")
		fmt.Println("calc 10.0.0.0/22 255.255.255.0")
		fmt.Println("calc --network=10.0.0.0/22 --subnetMask=26")
		return
	}

	ip, ipNet, err := net.ParseCIDR(networkCIDR)
	if err != nil {
		fmt.Println("Invalid CIDR")
		return
	}

	basePrefix, _ := ipNet.Mask.Size()

	// визначення нової маски
	var newPrefix int
	if strings.Contains(newMaskStr, ".") {
		newPrefix, err = maskStringToPrefix(newMaskStr)
		if err != nil {
			fmt.Println("Invalid subnet mask")
			return
		}
	} else {
		newPrefix, err = strconv.Atoi(newMaskStr)
		if err != nil {
			fmt.Println("Invalid mask number")
			return
		}
	}

	if newPrefix <= basePrefix || newPrefix > 32 {
		fmt.Println("New mask must be larger than base mask and <= 32")
		return
	}

	subnetCount := 1 << (newPrefix - basePrefix)
	hostCount := (1 << (32 - newPrefix)) - 2
	if newPrefix == 31 {
		hostCount = 2
	}
	if newPrefix == 32 {
		hostCount = 1
	}

	fmt.Printf("Вхідна мережа: %s\n", networkCIDR)
	fmt.Printf("Нова маска підмережі: /%d\n", newPrefix)
	fmt.Println("-----------------------------------")
	fmt.Printf("Кількість підмереж: %d\n", subnetCount)
	fmt.Printf("Кількість хостів в одній підмережі: %d\n", hostCount)
	fmt.Println("-----------------------------------")
	fmt.Println("Підмережа\tПерший Хост\tОстанній Хост\tBroadcast")
	fmt.Println("---------------------------------------------------------------------------")

	step := uint32(1 << (32 - newPrefix))
	start := ipToUint32(ip)

	for i := 0; i < subnetCount; i++ {
		netAddr := start + uint32(i)*step
		broadcast := netAddr + step - 1

		firstHost := netAddr + 1
		lastHost := broadcast - 1

		if newPrefix >= 31 {
			firstHost = netAddr
			lastHost = broadcast
		}

		fmt.Printf("%s/%d\t%s\t%s\t%s\n",
			uint32ToIP(netAddr),
			newPrefix,
			uint32ToIP(firstHost),
			uint32ToIP(lastHost),
			uint32ToIP(broadcast),
		)
	}
}
