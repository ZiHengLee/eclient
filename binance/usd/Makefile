
%.api.go : %.proto.go usd_tmpl_api.go genapi.py
	python3 genapi.py $< > $@

objs = \
	   usd_account.api.go \
	   usd_allorders.api.go \
	   usd_kline.api.go \
	   usd_openorder.api.go \
	   usd_order.api.go \

api: $(objs)
	@echo Done

clean:
	rm -rf $(objs)
