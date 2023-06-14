# Simple HTTP Server

## Testing the application locally

* Setup the environment variables
  ```bash
  export APP_USERNAME={username}
  export APP_PASSWORD={password}
    ```

* Start the go application
  ```bash
  cd ./src
  go run .
  ```
  Now the application should be running in your localhost at port 8080

* Test the application
  ```bash
  curl -i -u{username}:{password} http://localhost:8080/token     #should return 200
  curl -i http://localhost:8080/token                             #should return 401(unauthorzied)
  ```


## Deploying the application to an existing kubernetes cluster(with ingress controller installed)

* Create a secret to hold the application password 
  ```bash
  kubectl create secret generic simple-server --from-literal=APP_PASSWORD={password} -n {app_namespace}
  ```

* Deploy the kubernetes manifests
  ```bash
  kubectl apply -f manifestWithIngress.yaml -n {app_namespace}
  ```
  **Note:** The docker image repo/tag is hardcoded in the manifest file. It is defaulted to a public image which was pushed after building the latest version of the app.

* Test the application
  ```bash
  ingressIp=$(kubectl get ingress -n {app_namespace} simple-server -o yaml | yq .status.loadBalancer.ingress[0].ip)
  curl --resolve "simple-server.go:8080:$ingressIp" -uadmin:{password} -i http://simple-server.go:8080/token     #should return 200
  curl --resolve "simple-server.go:8080:$ingressIp" -i http://simple-server.go:8080/token                             #should return 401(unauthorzied)
  ```
 
  
## Deploying the application to an existing kubernetes cluster(using loadbalancer)

* Create a secret to hold the application password 
  ```bash
  kubectl create secret generic simple-server --from-literal=APP_PASSWORD={password} -n {app_namespace}
  ```

* Deploy the kubernetes manifests
  ```bash
  kubectl apply -f manifestWithLoadBalancer.yaml -n {app_namespace}
  ```
  **Note:** The docker image repo/tag is hardcoded in the manifest file. It is defaulted to a public image which was pushed after building the latest version of the app.

* Test the application
  ```bash
  lbIp=$(kubectl get svc -n {app_namespace} simple-server -o yaml | yq .status.loadBalancer.ingress[0].ip)
  curl -uadmin:{password} -i http://$lbIp:8080/token     #should return 200
  curl -i http://$lbIp:8080/token                             #should return 401(unauthorzied)
  ```
 
  