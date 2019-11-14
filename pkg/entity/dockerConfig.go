/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the genesis.

	Genesis is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Genesis is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package entity

//DockerConfig represents the configuration needed to communicate with docker daemons
type DockerConfig struct {
	//CACertPath is the filepath to the CA Certificate
	CACertPath string `json:"caCertPath"`
	//CertPath is the filepath to the Certificate for TLS
	CertPath string `json:"certPath"`
	//KeyPath is the filepath to the private key for TLS
	KeyPath string `json:"keyPath"`
	//LocalMode causes the TLS parameters to be ignored and Genesis
	//to assume that the docker daemon is on the local machine
	LocalMode bool `json:"localMode"`
}