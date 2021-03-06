package api

import (
	"fmt"
	"net"
)

func (client *Client) ipamOp(ID string, op string) (*net.IPNet, error) {
	ip, err := client.httpVerb(op, fmt.Sprintf("/ip/%s", ID), nil)
	if err != nil {
		return nil, err
	}
	return parseIP(ip)
}

// returns an IP for the ID given, allocating a fresh one if necessary
func (client *Client) AllocateIP(ID string) (*net.IPNet, error) {
	return client.ipamOp(ID, "POST")
}

func (client *Client) AllocateIPInSubnet(ID string, subnet *net.IPNet) (*net.IPNet, error) {
	ip, err := client.httpVerb("POST", fmt.Sprintf("/ip/%s/%s", ID, subnet), nil)
	if err != nil {
		return nil, err
	}
	return parseIP(ip)
}

// returns an IP for the ID given, or nil if one has not been
// allocated
func (client *Client) LookupIP(ID string) (*net.IPNet, error) {
	return client.ipamOp(ID, "GET")
}

// Claim a specific IP on behalf of the ID
func (client *Client) ClaimIP(ID string, addr net.IP) error {
	_, err := client.httpVerb("PUT", fmt.Sprintf("/ip/%s/%s", ID, addr), nil)
	return err
}

// release an IP which is no longer needed
func (client *Client) ReleaseIP(ID string) error {
	_, err := client.ipamOp(ID, "DELETE")
	return err
}

func (client *Client) DefaultSubnet() (*net.IPNet, error) {
	cidr, err := client.httpVerb("GET", fmt.Sprintf("/ipinfo/defaultsubnet"), nil)
	if err != nil {
		return nil, err
	}
	_, ipnet, err := net.ParseCIDR(cidr)
	return ipnet, err
}

func parseIP(body string) (*net.IPNet, error) {
	ip, ipnet, err := net.ParseCIDR(string(body))
	if err != nil {
		return nil, err
	}
	ipnet.IP = ip
	return ipnet, nil
}
