minikube delete
minikube start --cpus=36 --cni=bridge
source .env
minikube addons enable registry
minikube addons enable ingress
helm repo add cockroachdb https://charts.cockroachdb.com/
helm repo update
PROJECT_NAME=medicalchain-test
kubectl create ns ${PROJECT_NAME} --dry-run=client -o yaml | kubectl apply -f -
kubectl delete ns ${PROJECT_NAME}
helm delete cockroach-main --namespace=${PROJECT_NAME}
export $(grep -v '^#' .env | xargs) && envsubst < deploy/production/cockroach.values.tmp.yaml > cockroach.values.yaml
helm install cockroach-main -f cockroach.values.yaml --namespace=${PROJECT_NAME} cockroachdb/cockroachdb --version 6.0.9
kubectl -n=${PROJECT_NAME} get pvc,pv
kubectl -n=${PROJECT_NAME} get pv datadir-cockroach-main-cockroachdb-0 datadir-cockroach-main-cockroachdb-1 datadir-cockroach-main-cockroachdb-2
kubectl -n=${PROJECT_NAME} delete pvc datadir-cockroach-main-cockroachdb-0 datadir-cockroach-main-cockroachdb-1 datadir-cockroach-main-cockroachdb-2
kubectl -n=${PROJECT_NAME} delete pv persistentvolume/pvc-2ed877f7-93b0-4b55-81ea-dc2f6136d07a persistentvolume/pvc-9b7da370-12f4-43c9-b363-22a132b09460
kubectl -n=${PROJECT_NAME} describe pod cockroach
kubectl -n=${PROJECT_NAME} patch persistentvolume/pvc-0578da0a-1fb4-457a-8f38-364220d8cb89 -p '{"metadata":{"finalizers":null}}'
kubectl -n=${PROJECT_NAME} get pods,statefulsets,services,ingresses
kubectl -n=${PROJECT_NAME} logs -l app.kubernetes.io/component=cockroachdb --all-containers=true -f

echo '{"identifier":"0333888648","type":"PHONE"}' | grpc-client-cli -service medical_chain.AuthService -method GetCredential localhost:10020

export $(grep -v '^#' .env | xargs) && envsubst < deploy/production/chart/values.tmp.yaml > service.values.yaml
helm delete main-service --namespace=${PROJECT_NAME}
helm install main-service -f service.values.yaml --namespace=${PROJECT_NAME} deploy/production/chart/

helm template main-service -f service.values.yaml --namespace=${PROJECT_NAME} deploy/production/chart/

kubectl -n=${PROJECT_NAME} describe pod main-service
kubectl  -n=${PROJECT_NAME} get pods,statefulsets,services,ingresses,pv,pvc
kubectl -n=${PROJECT_NAME} logs -l app=main-service-selector --all-containers=true -f
kubectl -n=${PROJECT_NAME} logs -l app=main-service-selector -c init-auth-service  -f

kubectl  -n=${PROJECT_NAME} get pods,statefulsets,services,ingresses
#kubectl -n=${PROJECT_NAME} logs -l app=nginx-selector --all-containers=true -f
#kubectl -n=${PROJECT_NAME} logs -l app=backend-selector --all-containers=true -f
#kubectl -n=${PROJECT_NAME} logs -l app=blockchain-gateway-selector --all-containers=true -f
#kubectl -n=${PROJECT_NAME} logs -l app=nginx-selector --all-containers=true -f


helm repo add elastic https://helm.elastic.co

helm repo update

helm install es --namespace=linksrus-data   \
       --values chart-settings/es-settings.yaml \
       elastic/elasticsearch

kubectl port-forward pod/auth-service-deployment-649c4b9668-pgcnr 9999:80 --address 0.0.0.0  --namespace=${PROJECT_NAME}
kubectl port-forward cdb-cockroachdb-0 9999:8080 --namespace=linksrus-data

kubectl delete pvc
kubectl get pv

kubectl run -it --rm cockroach-client \
      --image=cockroachdb/cockroach \
      --restart=Never \
      --command -- \
      ./cockroach sql --insecure --host=cdb-cockroachdb-public.linksrus-data


export $(grep -v '^#' .env | xargs) && envsubst < deploy/production/sish-client-chart/values.tmp.yaml > main-tunnel.values.yaml
helm delete main-tunnel --namespace=${PROJECT_NAME}
helm install main-tunnel -f main-tunnel.values.yaml --namespace=${PROJECT_NAME} deploy/production/sish-client-chart/
kubectl -n=${PROJECT_NAME} describe pod sish-main
kubectl  -n=${PROJECT_NAME} get pods,statefulsets,services,ingresses
kubectl -n=${PROJECT_NAME} logs -l app=sish-main-selector --all-containers=true -f
kubectl --namespace=${PROJECT_NAME} exec --stdin --tty main-service-deployment-649c4b9668-pgcnr -- /bin/sh