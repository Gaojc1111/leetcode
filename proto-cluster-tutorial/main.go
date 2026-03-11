package main

import (
	"fmt"
	"proto-cluster-tutorial/smartbulb"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/asynkron/protoactor-go/cluster/clusterproviders/consul"
	"github.com/asynkron/protoactor-go/cluster/identitylookup/disthash"
	"github.com/asynkron/protoactor-go/remote"
)

func main() {

	// 创建标准Proto.Actor
	system := actor.NewActorSystem()

	provider, _ := consul.New()
	lookup := disthash.New()

	remoteConfig := remote.Configure("localhost", 0)

	bulbKind := smartbulb.NewSmartBulbKind(func() smartbulb.SmartBulb { return &smartbulb.Bulb{} }, 0)
	clusterConfig := cluster.Configure("ProtoClusterTutorial", provider, lookup, remoteConfig, cluster.WithKinds(bulbKind))
	// clusterConfig := cluster.Configure("ProtoClusterTutorial", provider, lookup, remoteConfig)

	// 创建集群对象 来和集群进行交互
	c := cluster.New(system, clusterConfig)

	c.StartMember()
	fmt.Println("Cluster member started. Press Enter to exit...")

	bulbClient := smartbulb.GetSmartBulbGrainClient(c, "kitchen")
	// if _, err := bulbClient.TurnOn(nil); err != nil {
	// 	panic(err)
	// }

	var count int
	for {
		if count > 30 {
			break
		}
		state, _ := bulbClient.GetState(nil)
		if !state.IsOn {
			bulbClient.TurnOn(nil)
			count++
			fmt.Println("Turning on")
		} else {
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Scanln()
	c.Shutdown(true)
}
