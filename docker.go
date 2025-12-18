package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func getRealContainers() ([]Container, error) {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        return nil, err
    }
    defer cli.Close()
    
    ctx := context.Background()
    containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
    if err != nil {
        return nil, err
    }
    
    var result []Container
    for _, c := range containers {
		name := c.Names[0]
		if len(name) > 0 && name[0] == '/' {
            name = name[1:]
        }
		if (len(c.ID) > 12) {
			c.ID = c.ID[:12]
		}
        result = append(result, Container{
			Id:      c.ID[:12],
			Name:    name,
			Status:  c.State,
			Image:   c.Image,
			Created: c.Created,
        })
    }
    
    return result, nil
}

func formatUptime(created int64) string {
	seconds := time.Now().Unix() - created
	hours := seconds / 3600
	days := seconds / 86400

	if days > 0 {
		return fmt.Sprintf("%dd %dh", days, (seconds%86400)/3600)
	} else if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, (seconds%3600)/60)
	} else {
		return fmt.Sprintf("%dm", seconds/60)
	}
}