.DEFAULT_GOAL := deploy

kubectl := kubectl --context="$(DOMAIN_NAME)" --namespace="$(NAMESPACE)"

SECRET_CREDENTIALS := `$(kubectl) --namespace="$(PULL_SECRET_NAMESPACE)" get secret $(PULL_SECRET_NAME) -o=json | jq -r '.data.".dockerconfigjson" | @base64d | fromjson | .auths | to_entries | .[0].value | "-a component.jenkins-credentials.username=\(.username),component.jenkins-credentials.password=\(.password)"'`

deploy:
	@env | sort
	sleep 10
ifneq ($(PULL_SECRET_NAME),)
	hub render templates/$(CREDENTIALS_KIND).yaml $(SECRET_CREDENTIALS)
endif
	- $(kubectl) apply -f templates/$(CREDENTIALS_KIND).yaml
.PHONY: deploy

undeploy:
	- $(kubectl) delete -f templates/$(CREDENTIALS_KIND).yaml
.PHONY: undeploy
