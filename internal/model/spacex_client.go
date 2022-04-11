package model

import (
	"context"
)

type SpaceXClient struct { }

func NewSpaceXClient() *SpaceXClient {
	client := SpaceXClient{}

	return &client
}

func (r *SpaceXClient) Open() bool {

	return true
}

func (r *SpaceXClient) Close() bool {

	return true
}

func (r SpaceXClient) GetLaunches(ctx context.Context) (*[]SpaceXLaunch, error) {

	return nil, nil
}
