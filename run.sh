bash ./kubernetes/cleanup.sh
sudo bash build-images.sh
kubectl apply -Rf ./kubernetes/apply
