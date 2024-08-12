rm -rf ./app_dist
rm -rf ./console_dist
cd app && yarn run build
cd ../console && yarn run build
cd ../