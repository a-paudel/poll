rm -r .embedded
mkdir .embedded
cd ../client
pnpm build
cd ../server
cp -r ../client/build .embedded/client
go build