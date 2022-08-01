all: clean data/processed/mpx-data.json

.PHONY:
clean:
	rm -f data/raw/* data/processed/mpx-data.json

data/processed/mpx-data.json: data/raw/mpx-data.json
	go run process_data.go $< $(basename $@) && \
	jq -s 'flatten' data/processed/*.json > $@

data/raw/mpx-data.json:
	wget -O $@ "https://www.cdc.gov/poxvirus/monkeypox/response/modules/MX-response-case-count-US.json"