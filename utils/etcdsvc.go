package utils

import (
	"context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// 服务注册

type Service struct {
	client *clientv3.Client
}

func NewService(etcdSvc []string) (*Service, error) {
	// 配置 etcd
	config := clientv3.Config{
		Endpoints:   etcdSvc,
		DialTimeout: 10 * time.Second,
	}
	cli, err := clientv3.New(config)
	if err != nil {
		log.Println("etcd 配置出错")
		return nil, err
	}
	return &Service{client: cli}, nil
}

// 注册服务
func (this *Service) RegService(serviceId string, serviceName string, serviceAddress string) error {
	ctx := context.Background()
	prefix := "/services/"
	key := prefix + serviceId + "/" + serviceName
	kv := clientv3.NewKV(this.client)
	lease := clientv3.NewLease(this.client)

	// 设置60秒的租约 etcdctl lease grand 60
	lresp, err := lease.Grant(ctx, 60)
	if err != nil {
		return err
	}

	// etcdctl put key value --lease=xxxxx(leaseid)
	_, err = kv.Put(ctx, key, serviceAddress, clientv3.WithLease(lresp.ID))
	if err != nil {
		return err
	}

	// 60秒后租约就过期了，设置定时续租就不会过期
	// etcdctl lease keep-alive xxxxx(leaseid)
	leaseChan, err := lease.KeepAlive(ctx, lresp.ID)
	go KeepLive(leaseChan)
	if err != nil {
		return err
	}

	return nil
}

// 反注册
func (this *Service) UnRegService(sid string) error {
	prefix := "/services/"
	key := prefix + sid
	kv := clientv3.NewKV(this.client)
	_, err := kv.Delete(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	return nil
}

// 定时续租（更新 TTL）
func KeepLive(leaseKeepAlive <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case ret := <-leaseKeepAlive:
			if ret != nil {
				log.Println("续租成功", time.Now())
			}
		}
	}
}
