package router_test

import (
	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/core/config/model"
	"github.com/go-chassis/go-chassis/core/router"
	_ "github.com/go-chassis/go-chassis/initiator"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	config.Init()
	t.Run("init router", func(t *testing.T) {
		config.RouterDefinition = &model.RouterConfig{}
		err := router.Init()
		assert.NoError(t, err)
	})

	t.Run("build a wrong router,return err", func(t *testing.T) {
		err := router.BuildRouter("wrong")
		assert.Error(t, err)
		config.RouterDefinition.Router.Infra = "wrong"
		err = router.Init()
		assert.Error(t, err)
	})
	t.Run("install and build a wrong router,return err", func(t *testing.T) {
		router.InstallRouterService("wrong", func() (router.Router, error) {
			return nil, errors.New("1")
		})
		err := router.BuildRouter("wrong")
		assert.Error(t, err)
	})
	t.Run("validate rule, exact 100, should success", func(t *testing.T) {
		ok := router.ValidateRule(map[string][]*model.RouteRule{
			"service1": {
				{Routes: []*model.RouteTag{{Weight: 50}, {Weight: 50}}},
			},
		})
		assert.True(t, ok)
	})
	t.Run("validate rule, greater than 100, should fail", func(t *testing.T) {
		ok := router.ValidateRule(map[string][]*model.RouteRule{
			"service2": {
				{Routes: []*model.RouteTag{{Weight: 900}, {Weight: 50}}},
			},
		})
		assert.False(t, ok)
	})

}
