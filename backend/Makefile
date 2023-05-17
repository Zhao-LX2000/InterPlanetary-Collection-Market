hz new api:
#    hz new -module github.com/IPAM --router_dir=router --handler_dir=handler --model_dir=model -idl ../../idl/api.thrift
    hz new --router_dir=router --handler_dir=handler --model_dir=model -idl ../../idl/api.thrift

hz update api:
#    hz update -module github.com/IPAM --model_dir=model -idl ../../idl/api.thrift

     hz update -module github.com/IPAM/server/api --model_dir=model --handler_dir=handler -idl ../../idl/api.thrift

kitex new server:

    kitex --thrift-plugin validator -module github.com/IPAM idl/artwork.thrift


    kitex --thrift-plugin validator -module github.com/IPAM -service test_user ../../idl/test_user.thrift


kitex new server with exist_kitex_gen：
    kitex --thrift-plugin validator -module github.com/IPAM -service artwork -use ../../kitex_gen ../../idl/artwork.thrift

kitex new server and change kitex_gen：
    kitex --thrift-plugin validator -module github.com/IPAM -service artwork -gen-path ../../kitex_gen ../../idl/artwork.thrift
    kitex --thrift-plugin validator -module github.com/IPAM -service artwork -gen-path ../../kitex_gen ../../idl/test_user.thrift
    kitex --thrift-plugin validator -module github.com/IPAM -service file -gen-path ./kitex_gen ./idl/file.thrift

