package main

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func startContainer(id string) error {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err
    }
    defer cli.Close()
    
    ctx := context.Background()
    return cli.ContainerStart(ctx, id, container.StartOptions{})
}

func stopContainer(id string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()
	return cli.ContainerStop(ctx, id, container.StopOptions{})
}

func removeContainer(id string) error {
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        return err
    }
    defer cli.Close()
    
    ctx := context.Background()
    return cli.ContainerRemove(ctx, id, container.RemoveOptions{Force: true})
}