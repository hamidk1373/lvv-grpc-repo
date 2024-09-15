s=nil
c=nil

gen:
	protoc --go_out=. --go-grpc_out=. allprotos/$(s).proto

pub: 
	git add .
	git commit -am "$(c)"
	git push
	./publish	