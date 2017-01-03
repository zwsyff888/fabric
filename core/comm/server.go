/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package comm

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//A SecureServerConfig structure is used to configure security (e.g. TLS) for a
//GRPCServer instance
type SecureServerConfig struct {
	//Whether or not to use TLS for communication
	UseTLS bool
	//PEM-encoded X509 public key to be used by the server for TLS communication
	ServerCertificate []byte
	//PEM-encoded private key to be used by the server for TLS communication
	ServerKey []byte
	//Set of PEM-encoded X509 certificate authorities to optionally send
	//as part of the server handshake
	ServerRootCAs [][]byte
	//Whether or not TLS client must present certificates for authentication
	RequireClientCert bool
	//Set of PEM-encoded X509 certificate authorities to use when verifying
	//client certificates
	ClientRootCAs [][]byte
}

//GRPCServer defines an interface representing a GRPC-based server
type GRPCServer interface {
	//Address returns the listen address for the GRPCServer
	Address() string
	//Start starts the underlying grpc.Server
	Start() error
	//Stop stops the underlying grpc.Server
	Stop()
	//Server returns the grpc.Server instance for the GRPCServer
	Server() *grpc.Server
	//Listener returns the net.Listener instance for the GRPCServer
	Listener() net.Listener
	//ServerCertificate returns the tls.Certificate used by the grpc.Server
	ServerCertificate() tls.Certificate
	//TLSEnabled is a flag indicating whether or not TLS is enabled for this
	//GRPCServer instance
	TLSEnabled() bool
}

type grpcServerImpl struct {
	//Listen address for the server specified as hostname:port
	address string
	//Listener for handling network requests
	listener net.Listener
	//GRPC server
	server *grpc.Server
	//Certificate presented by the server for TLS communication
	serverCertificate tls.Certificate
	//Key used by the server for TLS communication
	serverKeyPEM []byte
	//List of certificate authorities to optionally pass to the client during
	//the TLS handshake
	serverRootCAs []tls.Certificate
	//List of certificate authorities to be used to authenticate clients if
	//client authentication is required
	clientRootCAs *x509.CertPool
	//Is TLS enabled?
	tlsEnabled bool
}

//NewGRPCServer creates a new implementation of a GRPCServer given a
//listen address.
func NewGRPCServer(address string, secureConfig SecureServerConfig) (GRPCServer, error) {

	if address == "" {
		return nil, errors.New("Missing address parameter")
	}
	//create our listener
	lis, err := net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	return NewGRPCServerFromListener(lis, secureConfig)

}

//NewGRPCServerFromListener creates a new implementation of a GRPCServer given
//an existing net.Listener instance.
func NewGRPCServerFromListener(listener net.Listener, secureConfig SecureServerConfig) (GRPCServer, error) {

	grpcServer := &grpcServerImpl{
		address:  listener.Addr().String(),
		listener: listener,
	}

	//set up our server options
	var serverOpts []grpc.ServerOption
	//check secureConfig
	if secureConfig.UseTLS {
		//both key and cert are required
		if secureConfig.ServerKey != nil && secureConfig.ServerCertificate != nil {
			grpcServer.tlsEnabled = true
			//load server public and private keys
			cert, err := tls.X509KeyPair(secureConfig.ServerCertificate, secureConfig.ServerKey)
			if err != nil {
				return nil, err
			}
			grpcServer.serverCertificate = cert

			//set up our TLS config

			//base server certificate
			certificates := []tls.Certificate{grpcServer.serverCertificate}
			tlsConfig := &tls.Config{
				Certificates: certificates,
			}
			//checkif client authentication is required
			if secureConfig.RequireClientCert {
				//require TLS client auth
				tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
				//if we have client root CAs, create a certPool
				if len(secureConfig.ClientRootCAs) > 0 {
					grpcServer.clientRootCAs = x509.NewCertPool()
					for _, clientRootCA := range secureConfig.ClientRootCAs {
						if !grpcServer.clientRootCAs.AppendCertsFromPEM(clientRootCA) {
							return nil, errors.New("Failed to load client root certificates")
						}
					}
					tlsConfig.ClientCAs = grpcServer.clientRootCAs
				}
			}

			//create credentials
			creds := credentials.NewTLS(tlsConfig)

			//add to server options
			serverOpts = append(serverOpts, grpc.Creds(creds))

		} else {
			return nil, errors.New("secureConfig must contain both ServerKey and " +
				"ServerCertificate when UseTLS is true")
		}
	}
	grpcServer.server = grpc.NewServer(serverOpts...)

	return grpcServer, nil
}

//Address returns the listen address for this GRPCServer instance
func (gServer *grpcServerImpl) Address() string {
	return gServer.address
}

//Listener returns the net.Listener for the GRPCServer instance
func (gServer *grpcServerImpl) Listener() net.Listener {
	return gServer.listener
}

//Server returns the grpc.Server for the GRPCServer instance
func (gServer *grpcServerImpl) Server() *grpc.Server {
	return gServer.server
}

//ServerCertificate returns the tls.Certificate used by the grpc.Server
func (gServer *grpcServerImpl) ServerCertificate() tls.Certificate {
	return gServer.serverCertificate
}

//TLSEnabled is a flag indicating whether or not TLS is enabled for the
//GRPCServer instance
func (gServer *grpcServerImpl) TLSEnabled() bool {
	return gServer.tlsEnabled
}

//Start starts the underlying grpc.Server
func (gServer *grpcServerImpl) Start() error {
	return gServer.server.Serve(gServer.listener)
}

//Stop stops the underlying grpc.Server
func (gServer *grpcServerImpl) Stop() {
	gServer.server.Stop()
}
