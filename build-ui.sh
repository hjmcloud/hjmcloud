set -e 
cd ui
npm run build
rm -rf ../public/ui
cp -r ./dist ../public/ui
cd -