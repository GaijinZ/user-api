eval $(minikube docker-env)
minikube kubectl -- delete -f userapi.yaml
minikube kubectl -- create -f userapi.yaml
