package config

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/JSON"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/YAML"
)

type CcpClientConnectionPeerEndorser struct {
	Endorser string `json:"endorser" yaml:"endorser"`
}

type CcpClientConnectionPeer struct {
	Peer *CcpClientConnectionPeerEndorser `json:"peer" yaml:"peer"`
}

type CcpClientConnection struct {
	Timeout *CcpClientConnectionPeer `json:"timeout" yaml:"timeout"`
}

type CcpClient struct {
	Organization string               `json:"organization" yaml:"organization"`
	Connection   *CcpClientConnection `json:"connection" yaml:"connection"`
}

type CcpOrganization struct {
	MspId                  string   `json:"mspid" yaml:"mspid"`
	Peers                  []string `json:"peers" yaml:"peers"`
	CertificateAuthorities []string `json:"certificateAuthorities" yaml:"certificateAuthorities"`
}

type CcpPeerTlsCACerts struct {
	Pem string `json:"pem" yaml:"pem"`
}

type CcpPeerGrpcOptions struct {
	SslTargetNameOverride string `json:"ssl-target-name-override" yaml:"ssl-target-name-override"`
	HostnameOverride      string `json:"hostnameOverride" yaml:"hostname-override"`
}

type CcpPeer struct {
	Url         string              `json:"url" yaml:"url"` // grpc
	TlsCACerts  *CcpPeerTlsCACerts  `json:"tlsCaCerts" yaml:"tlsCACerts"`
	GrpcOptions *CcpPeerGrpcOptions `json:"grpcOptions" yaml:"grpcOptions"`
}

type CcpCaTlsCaCert struct {
	Pem []string `json:"pem" yaml:"pem"`
}

type CcpCaHttpOptions struct {
	Verify bool `json:"verify" yaml:"verify"`
}

type CcpCertificateAuthorities struct {
	Url         string            `json:"url" yaml:"url"`
	CaName      string            `json:"caName" yaml:"caName"`
	TlsCaCerts  *CcpCaTlsCaCert   `json:"tlsCaCerts" yaml:"tlsCaCerts"`
	HttpOptions *CcpCaHttpOptions `json:"httpOptions" yaml:"httpOptions"`
}

type CCP struct {
	Name                   string                                `json:"name" yaml:"name"`
	Version                string                                `json:"version" yaml:"version"`
	Client                 *CcpClient                            `json:"client" yaml:"client"`
	Organizations          map[string]*CcpOrganization           `json:"organizations" yaml:"organizations"`
	Peers                  map[string]*CcpPeer                   `json:"peers" yaml:"peers"`
	CertificateAuthorities map[string]*CcpCertificateAuthorities `json:"certificateAuthorities" yaml:"certificateAuthorities"`
}

func GenerateEmptyCcp(name string, version string, orgName string) *CCP {
	return &CCP{
		Name:    name,
		Version: version,
		Client: &CcpClient{
			Organization: orgName,
			Connection:   &CcpClientConnection{Timeout: &CcpClientConnectionPeer{Peer: &CcpClientConnectionPeerEndorser{Endorser: "300"}}},
		},
		Organizations: map[string]*CcpOrganization{
			orgName: {
				MspId:                  fmt.Sprintf("%sMSP", orgName),
				Peers:                  []string{},
				CertificateAuthorities: []string{},
			},
		},
		Peers:                  map[string]*CcpPeer{},
		CertificateAuthorities: map[string]*CcpCertificateAuthorities{},
	}
}

func (that *CCP) AddPeer(peerName string, orgName string, domainRoot string, gRpcPort uint, peerPemPath string) {
	domain := fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot)
	that.Organizations[orgName].Peers = append(that.Organizations[orgName].Peers, domain)
	that.Peers[domain] = &CcpPeer{
		Url:        fmt.Sprintf("grpcs://%s:%d", domain, gRpcPort),
		TlsCACerts: &CcpPeerTlsCACerts{Pem: peerPemPath},
	}
}

func (that *CCP) AddCa(peerName string, orgName string, domainRoot string, gRpcPort uint, caPem string, verify bool) {
	domain := fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot)
	that.Organizations[orgName].CertificateAuthorities = append(that.Organizations[orgName].CertificateAuthorities, domain)
	that.CertificateAuthorities[domain] = &CcpCertificateAuthorities{
		Url:    fmt.Sprintf("https://%s:%d", domain, gRpcPort),
		CaName: fmt.Sprintf("ca%s%s", peerName, orgName),
		TlsCaCerts: &CcpCaTlsCaCert{Pem: []string{
			caPem,
		}},
		HttpOptions: &CcpCaHttpOptions{Verify: verify},
	}
}

func (that *CCP) ExportYaml(path string) error {
	return YAML.ExportToFileYaml(that, path)
}

func (that *CCP) ExportJson(path string) error {
	return JSON.ExportToFileJson(that, path)
}
