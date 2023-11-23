install:
	kubectl create namespace $(STAGE)
	helm install --values values.yaml --values "values.$(STAGE).yaml" --namespace $(STAGE) simple-web-server ./
upgrade:
	helm upgrade --values values.yaml --values "values.$(STAGE).yaml" --namespace $(STAGE) simple-web-server ./
uninstall:
	helm uninstall --namespace $(STAGE) simple-web-server
	kubectl delete namespace $(STAGE)

.PHONY: install upgrade uninstall
