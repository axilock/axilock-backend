npm run predeploy
rm -rf ../prod-frontend/static/*; cp -r ./build/* ../prod-frontend/
cd ../prod-frontend/
git add .
git commit -m "Deploy frontend"
git push
cd ../axi-frontend/