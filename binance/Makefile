
%.api.go : %.proto.go tmpl_api.go genapi.py
	python3 genapi.py $< > $@

objs = \
	   accountinfo.api.go \
	   postorder.api.go \
	   getorder.api.go \
	   deleteorder.api.go \
	   getopenorders.api.go \
	   allorders.api.go \
	   mytrades.api.go \
	   exchange_info.api.go \
	   depth.api.go \
	   userdatastream.api.go \
	   getfunding_asset.api.go \
	   api_restrictions.api.go \
	   tickerprice.api.go \
	   ticker24hr.api.go \
	   kline.api.go \

api: $(objs)
	@echo Done

clean:
	rm -rf $(objs)
