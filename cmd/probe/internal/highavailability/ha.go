package highavailability

import (
	"context"
	"time"

	"github.com/dapr/kit/logger"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/apecloud/kubeblocks/cmd/probe/internal/component"
	"github.com/apecloud/kubeblocks/cmd/probe/internal/component/mongodb"
	"github.com/apecloud/kubeblocks/cmd/probe/internal/dcs"
)

type Ha struct {
	ctx       context.Context
	dbManager component.DBManager
	dcs       dcs.DCS
	logger    logger.Logger
}

func NewHa(logger logger.Logger) *Ha {

	dcs, _ := dcs.NewKubernetesStore(logger)
	ha := &Ha{
		ctx:       context.Background(),
		dcs:       dcs,
		logger:    logger,
		dbManager: mongodb.Mgr,
	}
	return ha
}

func (ha *Ha) RunCycle() {
	cluster, err := ha.dcs.GetCluster()
	if err != nil {
		ha.logger.Infof("Get Cluster err: %v.", err)
		return
	}

	switch {
	//case !dbManger.IsRunning():
	//	logger.Infof("DB Service is not running,  wait for sqlctl to start it")
	//	if dcs.HasLock() {
	//		dcs.ReleaseLock()
	//	}

	//case !dbManger.IsHealthy():
	//	logger.Infof("DB Service is not healthy,  do some recover")
	//	if dcs.HasLock() {
	//		dcs.ReleaseLock()
	//	}
	//	dbManager.Recover()

	//case !cluster.HasThisMember():
	//	logger.Infof("Current node is not in cluster, add it to cluster")
	//	dbManager.AddToCluser(cluster)
	//	cluster.AddThisMember()

	//case cluster.IsUnlocked():
	//	logger.Infof("Cluster has no leader, attemp to take the leader")
	//	candidate := ""
	//	if cluster.SwitchOver != nil && cluster.SwitchOver.Candinate != "" {
	//		candiate = cluster.SwitchOver.Candinate
	//	}
	//	healthiestMember := ha.dbManager.GetHealthiesMember(cluster, candinate)
	//	if healthiestMember != nil && healthiestMember.Name == dbManager.MemberName {
	//		if dcs.attempAcquireLeader() {
	//			dbManager.Premote()
	//		}
	//	}

	case cluster.Leader.Name == ha.dbManager.GetCurrentMemberName():
		ha.logger.Infof("This member is Cluster's leader")
		if ok, _ := ha.dbManager.IsLeader(context.TODO()); ok {
			ha.logger.Infof("Refresh leader ttl")
			ha.dcs.UpdateLock()
		} else if ha.dbManager.HasOtherHealthyLeader() != nil {
			ha.logger.Infof("Release leader")
			ha.dcs.ReleaseLock()
		}

		//case cluster.SwitchOver.Leader != dbManager.Name && dbManager.IsLeader():
		//	logger.Infof("Cluster has no leader, attemp to take the leader")
		//	dbManager.Demote()

		//case cluster.SwitchOver.Leader == dbManager.Name && !dbManager.IsLeader():
		//	logger.Infof("Cluster has no leader, attemp to take the leader")
		//	dbManager.Premote()
	}
}

func (ha *Ha) Start() {
	ha.logger.Info("HA starting")
	cluster, err := ha.dcs.GetCluster()
	if errors.IsNotFound(err) {
		ha.logger.Infof("Cluster %s is not found, so HA exists.", ha.dcs.GetClusterName())
		return
	}

	ha.logger.Debugf("cluster: %v", cluster)
	isInitialized, _ := ha.dbManager.IsClusterInitialized()
	for !isInitialized {
		ha.logger.Infof("Waiting for the database cluster to be initialized.")
		// TODO: implement dbmanager initialize to replace pod's entrypoint scripts
		// if I am the node of index 0, then do initialization
		// ha.dbManager.Initialize()
		time.Sleep(1 * time.Second)
		isInitialized, _ = ha.dbManager.IsClusterInitialized()
	}
	ha.logger.Infof("The database cluster is initialized.")

	isExist, _ := ha.dcs.IsLockExist()
	for !isExist {
		if ok, _ := ha.dbManager.IsLeader(context.Background()); ok {
			ha.dcs.Initialize()
			break
		}
		ha.logger.Infof("Waiting for the database Leader to be ready.")
		time.Sleep(1 * time.Second)
		isExist, _ = ha.dcs.IsLockExist()
	}

	for true {
		ha.RunCycle()
		time.Sleep(1 * time.Second)
		return
	}
}

func (ha *Ha) ShutdownWithWait() {
}
