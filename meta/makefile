base = fausto
repository = fausto
# tag = $(word 2, $(MAKECMDGOALS))
tag = lib-go
function = $(base)-$(tag)
role = lambda-ex
region = sa-east-1
image = $(base):$(tag)

logs:
	@aws logs tail /aws/lambda/$(function) --follow

deploy:
	@docker build --platform linux/amd64 -t $(image) .
	@docker login --username AWS --password $$(aws ecr get-login-password --region $(region)) $$(aws sts get-caller-identity | jq -r '.Account').dkr.ecr.$(region).amazonaws.com
	@aws ecr create-repository --repository-name $(repository) --region $(region) --image-scanning-configuration scanOnPush=true --image-tag-mutability MUTABLE
	@docker tag $(image) $$(aws ecr describe-repositories | jq -r '.repositories[] | select(.repositoryName == "$(repository)").repositoryUri'):$(tag)
	@docker push $$(aws ecr describe-repositories | jq -r '.repositories[] | select(.repositoryName == "$(repository)").repositoryUri'):$(tag)
	# @aws iam create-role --role-name $(role) \--assume-role-policy-document '{"Version": "2012-10-17","Statement": [{ "Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'
	# @aws iam attach-role-policy --role-name $(role) --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
	@aws lambda create-function --function-name $(function) --package-type Image --code ImageUri=$$(aws ecr describe-repositories | jq -r '.repositories[] | select(.repositoryName == "$(repository)").repositoryUri'):$(tag) --role $$(aws iam get-role --role-name $(role) --output json | jq -r '.Role.Arn')
	@aws lambda update-function-code --function-name $(function) --image-uri $$(aws ecr describe-repositories | jq -r '.repositories[] | select(.repositoryName == "$(repository)").repositoryUri'):$(tag)
