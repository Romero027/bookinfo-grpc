set -ex
go install -ldflags="-s -w" ./cmd/...
bash ./kubernetes/cleanup.sh
sudo rm -rf /data/volumes
sudo bash build-images.sh
kubectl apply -Rf ./kubernetes/apply
echo "wait for apply"
sleep 100
./wrk/wrk -t1 -c1 -d 1s http://10.96.88.88:8080 -L -s ./scripts/lua/bookinfo.lua
set +ex