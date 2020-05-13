#!/bin/bash
# publish-plugin.sh builds and publishes a package containing the resource provider to
# s3://rel.pulumi.com/releases/plugins.
set -o nounset -o errexit -o pipefail

# Update this with the name of the provider
PROVIDER_NAME="vault"

ROOT=$(dirname $0)/..
WORK_PATH=$(mktemp -d)
VERSION=$(jq -r '.version' < "${ROOT}/sdk/nodejs/bin/package.json")
PLUGIN_PACKAGE_NAME="pulumi-resource-${PROVIDER_NAME}-${VERSION}-$(go env GOOS)-$(go env GOARCH).tar.gz"
PLUGIN_PACKAGE_DIR="$(mktemp -d)"
PLUGIN_PACKAGE_PATH="${PLUGIN_PACKAGE_DIR}/${PLUGIN_PACKAGE_NAME}"

# When crossbuilding, we want to ensure we have .exe for the windows binaries.
BIN_SUFFIX=
if [ "$(go env GOOS)" = "windows" ]; then
    BIN_SUFFIX=".exe"
fi

(cd "${ROOT}/provider" && go build \
   -ldflags "-X github.com/pulumi/pulumi-${PROVIDER_NAME}/provider/v2/pkg/version.Version=${VERSION}" \
   -o "${WORK_PATH}/pulumi-resource-${PROVIDER_NAME}${BIN_SUFFIX}" \
   "${ROOT}/cmd/pulumi-resource-${PROVIDER_NAME}")

# Tar up the plugin
tar -czf ${PLUGIN_PACKAGE_PATH} -C ${WORK_PATH} .

# Push the current AWS credentials, since we'll need them to assume a second role.
export INITIAL_AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
export INITIAL_AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"

# rel.pulumi.com is in our production account, so assume that role first
CREDS_JSON=$(aws sts assume-role \
                 --role-arn "arn:aws:iam::058607598222:role/UploadPulumiReleases" \
                 --role-session-name "upload-plugin-pulumi-resource-vault" \
                 --external-id "upload-pulumi-release")

# Use the credentials we just assumed.
export AWS_ACCESS_KEY_ID=$(echo ${CREDS_JSON}     | jq ".Credentials.AccessKeyId" --raw-output)
export AWS_SECRET_ACCESS_KEY=$(echo ${CREDS_JSON} | jq ".Credentials.SecretAccessKey" --raw-output)
export AWS_SECURITY_TOKEN=$(echo ${CREDS_JSON}    | jq ".Credentials.SessionToken" --raw-output)

echo "Uploading ${PLUGIN_PACKAGE_NAME} to s3://rel.pulumi.com..."

aws s3 cp --only-show-errors "${PLUGIN_PACKAGE_PATH}" "s3://rel.pulumi.com/releases/plugins/${PLUGIN_PACKAGE_NAME}"

# Assume the role to publish plugins to s3://get.pulumi.com. We upload the plugins to two buckets while
# we transition to only publishing/serving them from get.pulumi.com.
echo "Uploading ${PLUGIN_PACKAGE_NAME} to s3://get.pulumi.com..."

# Restore the initial AWS credentials we had, since the assumed role doesn't have the
# ability to assume this other role to publish into a different bucket.
unset {AWS_ACCESS_KEY_ID,AWS_SECRET_ACCESS_KEY,AWS_SECURITY_TOKEN}
export AWS_ACCESS_KEY_ID="${INITIAL_AWS_ACCESS_KEY_ID}"
export AWS_SECRET_ACCESS_KEY="${INITIAL_AWS_SECRET_ACCESS_KEY}"

CREDS_JSON=$(aws sts assume-role \
                 --role-arn "arn:aws:iam::058607598222:role/PulumiUploadRelease" \
                 --role-session-name "upload-plugin-pulumi-resource-vault" \
                 --external-id "upload-pulumi-release")
export AWS_ACCESS_KEY_ID=$(echo ${CREDS_JSON}     | jq ".Credentials.AccessKeyId" --raw-output)
export AWS_SECRET_ACCESS_KEY=$(echo ${CREDS_JSON} | jq ".Credentials.SecretAccessKey" --raw-output)
export AWS_SECURITY_TOKEN=$(echo ${CREDS_JSON}    | jq ".Credentials.SessionToken" --raw-output)

aws s3 cp \
    --only-show-errors --acl public-read \
    "${PLUGIN_PACKAGE_PATH}" "s3://get.pulumi.com/releases/plugins/${PLUGIN_PACKAGE_NAME}"

rm -rf "${PLUGIN_PACKAGE_DIR}"
rm -rf "${WORK_PATH}"
