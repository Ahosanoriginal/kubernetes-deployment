# Kubernetes-deployment
Simple Go microservices deployed on minikube



'''
$ kubectl create -f gateway-deployment.yaml 

(same as kubectl run gateway-deployment --image=gateway:latest --port=8080 --label "app=gateway-app)

'''



To know port and name for a deployement or a service add it after your command: 

'''
$  -o=custom-columns=NAME:.metadata.name,IP:.spec.clusterIP


'''

To know the url of a service outside the cluster :

 '''
 $ minikube service gateway-service --url
 '''