package scm

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

type etcdSCM struct {
	v3 *clientv3.Client
}

func newEtcdClient(addr []string) (*etcdSCM, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: timeout,
	})
	if err != nil {
		return nil, err
	}
	etcd := &etcdSCM{
		v3: cli,
	}
	return etcd, nil
}

// Close a etcd client
func (c *etcdSCM) Close() {
	c.v3.Close()
}

// Put a value by key to etcd
func (c *etcdSCM) Put(key string, val string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err := c.v3.Put(ctx, key, val)
	defer cancel()
	return err
}

// Get a value by key from etcd
func (c *etcdSCM) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	resp, err := c.v3.Get(ctx, key)
	defer cancel()
	if err != nil {
		return nil, err
	}
	var buf []byte
	if resp.Count > 0 {
		buf = resp.Kvs[0].Value
	}
	return buf, nil
}
