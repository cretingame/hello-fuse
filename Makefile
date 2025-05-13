EXEC = hello-fuse
MOUNT_POINT = /run/user/$(shell id -u)/hello


help:
	@echo make mount or make umount


$(MOUNT_POINT):
	install --group=$(shell id -g) --owner=$(shell id -g) --directory $(MOUNT_POINT)

$(CURDIR)/$(EXEC):
	go build

run: $(MOUNT_POINT) $(CURDIR)/$(EXEC)
	$(CURDIR)/$(EXEC) $(MOUNT_POINT)

clean:
	rm -f $(CURDIR)/$(EXEC)
	rmdir $(MOUNT_POINT)


.PHONY: clean
