package handlers

import (
	"context"
	app "github.com/obada-protocol/demo-service/http"
	nodes "github.com/obada-protocol/demo-service/system/nodes"
	"net/http"
)

type nodesGroup struct{}

func (ng nodesGroup) all(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	nodes, err := nodes.GetNodes()

	if err != nil {
		return err
	}

	return app.RespondJson(ctx, w, nodes, http.StatusOK)
}
