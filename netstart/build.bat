
set GOOS=linux

rm -rf target/*
mkdir target

go build -ldflags "-X jsj.golangtc/netstart/netstar.Env=true" -o target/netstar_test
go build -ldflags "-X jsj.golangtc/netstart/netstar.Env=false" -o target/netstar