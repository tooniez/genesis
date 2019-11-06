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

package usecase

import(
	"context"
	"github.com/docker/docker/client"
	"github.com/whiteblock/genesis/pkg/entity"
	"github.com/whiteblock/genesis/pkg/service"
)

type DockerUseCase interface{
	Execute(ctx context.Context,cmd command.Command) entity.Result
}

type dockerUseCase struct {
	conf entity.DockerConfig
	service service.DockerService
}

func NewDockerUseCase(conf entity.DockerConfig, service service.DockerService) (DockerUseCase, error) {
	return dockerUseCase{conf:conf, service: service},nil
}

func (duck dockerUseCase) Execute(ctx context.Context, cmd command.Command) entity.Result {
	cli, err := NewClientWithOpts(
		client.WithAPIVersionNegotiation(),
		client.WithHost(cmd.Target.IP),
		client.WithTLSClientConfig(duck.conf.CACertPath,duck.conf.CertPath,duck.conf.KeyPath)
	)
	if err != nil {
		return entity.Result{Error:err}
	}
	//TODO: route it to the right function call in service
}

