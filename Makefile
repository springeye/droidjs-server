PROTO_CC :=  protoc
INCLUDE  :=  .
PROTO_FLAGS := --go_out=$(INCLUDE)

PROTO__SOURCES := $(wildcard proto/*.proto)
obj-y          := $(PROTO__SOURCES:%.proto=$(INCLUDE)/%.pb.go)

OBJECT      := #unused
##############################
#  change OBJECT to set execute file name
##############################

all : $(obj-y)

.PHONY: all 
$(INCLUDE)/%.pb.go: %.proto
	$(PROTO_CC) $(PROTO_FLAGS) $^
	@echo 'generate file' $@

clean:
	rm -f $(obj-y)