package scm

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

type etcdSCM struct {
	cli *clientv3.Client
	kv  clientv3.KV
}

func newEtcdClient(addr ...string) (*etcdSCM, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: timeout,
	})
	if err != nil {
		return nil, err
	}
	etcd := &etcdSCM{
		cli: client,
		kv:  clientv3.NewKV(client),
	}
	return etcd, nil
}

// Close a etcd client
func (c *etcdSCM) Close() {
	c.cli.Close()
}

// Put key-value to etcd
func (c *etcdSCM) Put(key string, val string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err := c.kv.Put(ctx, key, val)
	defer cancel()
	return err
}

// Get a value by key from etcd
func (c *etcdSCM) Get(key string) (map[string][]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	resp, err := c.kv.Get(ctx, key, clientv3.WithPrefix())
	defer cancel()
	if err != nil {
		return nil, err
	}
	buf := make(map[string][]byte)
	for _, v := range resp.Kvs {
		buf[string(v.Key)] = v.Value
	}
	return buf, nil
}

// Delete key-value from etcd
func (c *etcdSCM) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err := c.kv.Delete(ctx, key)
	defer cancel()
	return err
}
