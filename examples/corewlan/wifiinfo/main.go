// Command wifiinfo reports the state of each Wi-Fi interface on the system:
// power, mode, SSID, BSSID, RSSI, noise, channel, and transmit rate.
//
// Reading the SSID and BSSID requires location authorization; without it those
// fields are reported as empty.
//
// Usage: wifiinfo
package main

import (
	"fmt"
	"os"

	"github.com/tmc/apple/corewlan"
)

func main() {
	client := corewlan.GetCWWiFiClientClass().SharedWiFiClient()
	if client.GetID() == 0 {
		fmt.Fprintf(os.Stderr, "wifiinfo: no Wi-Fi client available\n")
		os.Exit(1)
	}

	interfaces := client.Interfaces()
	if len(interfaces) == 0 {
		fmt.Fprintf(os.Stderr, "wifiinfo: no Wi-Fi interfaces found\n")
		os.Exit(1)
	}

	for _, iface := range interfaces {
		fmt.Printf("%s\n", iface.InterfaceName())
		fmt.Printf("\thardware address: %s\n", or(iface.HardwareAddress(), "unknown"))
		fmt.Printf("\tpower on:         %t\n", iface.PowerOn())
		fmt.Printf("\tservice active:   %t\n", iface.ServiceActive())
		fmt.Printf("\tmode:             %s\n", iface.InterfaceMode())
		fmt.Printf("\tcountry code:     %s\n", or(iface.CountryCode(), "unknown"))
		fmt.Printf("\tSSID:             %s\n", or(iface.Ssid(), "unavailable"))
		fmt.Printf("\tBSSID:            %s\n", or(iface.Bssid(), "unavailable"))
		fmt.Printf("\tsecurity:         %s\n", iface.Security())
		fmt.Printf("\tPHY mode:         %s\n", iface.ActivePHYMode())
		fmt.Printf("\tRSSI:             %d dBm\n", iface.RssiValue())
		fmt.Printf("\tnoise:            %d dBm\n", iface.NoiseMeasurement())
		fmt.Printf("\ttransmit power:   %d mW\n", iface.TransmitPower())
		fmt.Printf("\ttransmit rate:    %g Mbps\n", iface.TransmitRate())

		channel := iface.WlanChannel()
		if channel.GetID() == 0 {
			fmt.Printf("\tchannel:          none\n")
			continue
		}
		fmt.Printf("\tchannel:          %d (%s, %s)\n",
			channel.ChannelNumber(), channel.ChannelBand(), channel.ChannelWidth())
	}
}

func or(s, alt string) string {
	if s == "" {
		return alt
	}
	return s
}
