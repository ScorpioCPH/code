GO = go
GOBUILD = $(GO) build
DEST = ./bin

.PHONY : release
release :
	$(GOBUILD) -o $(DEST)/docker-client ./src

.PHONY : docker-client
docker-client:
	$(GOBUILD) -race -o $(DEST)/$@ ./src

# .PHONY : e2e
# e2e :
# 	cd ./test && bash ./e2e.sh

# .PHONY : test
# test : | docker-client e2e clean

.PHONY : clean
clean :
	rm -rf $(DEST)

