package eventlistener

import (
	"github.com/go-chassis/go-archaius/event"
	"github.com/go-chassis/go-chassis/control/archaius"
	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-mesh/openlogging"
)

// constants for loadbalancer strategy name, and timeout
const (
	//LoadBalanceKey is variable of type string that matches load balancing events
	LoadBalanceKey          = "^cse\\.loadbalance\\."
	regex4normalloadbalance = "^cse\\.loadbalance\\.(strategy|SessionStickinessRule|retryEnabled|retryOnNext|retryOnSame|backoff)"
)

//LoadbalancingEventListener is a struct
type LoadbalancingEventListener struct {
	Key string
}

//Event is a method used to handle a load balancing event
func (e *LoadbalancingEventListener) Event(evt *event.Event) {
	openlogging.GetLogger().Debugf("LB event, key: %s, type: %s", evt.Key, evt.EventType)
	if err := config.ReadLBFromArchaius(); err != nil {
		openlogging.GetLogger().Error("can not unmarshal new lb config: " + err.Error())
	}
	archaius.SaveToLBCache(config.GetLoadBalancing())
}
