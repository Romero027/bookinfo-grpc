bash ./kubernetes/cleanup.sh
sudo bash build-images.sh
kubectl apply -f ./kubernetes/bookinfo-grpc.yaml
kubectl apply -f ./kubernetes/jaeger.yaml