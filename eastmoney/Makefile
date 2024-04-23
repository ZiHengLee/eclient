
%.api.go : %.proto.go tmpl_api.go genapi.py
	python3 genapi.py $< > $@

objs = \
	   stock_info.api.go \

api: $(objs)
	@echo Done

clean:
	rm -rf $(objs)
