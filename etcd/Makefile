.PHONY:

# proto

gen-proto-v1:
	protoc -I proto --goV1_out=pkg/model proto/*.proto

gen-proto-v2:
	protoc -I proto --goV2_out=pkg/model proto/*.proto

# etcd

etcd-cluster-run:
	docker-compose -f build/docker-compose.yml up -d 

etcd-cluster-stop:
	docker-compose -f build/docker-compose.yml stop

etcd-cluster-rm:
	docker-compose -f build/docker-compose.yml rm