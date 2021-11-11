module github.com/p76081158/5g-nsmf

go 1.14

require (
	github.com/Arafatk/glot v0.0.0-20180312013246-79d5219000f0
	github.com/p76081158/free5gc-nssmf v0.0.0-20210616172500-b20fa150920c
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/p76081158/5g-nsmf/api/f5gnssmf => ./api/f5gnssmf
	github.com/p76081158/5g-nsmf/modules/executor => ./modules/executor
	github.com/p76081158/5g-nsmf/modules/logger => ./modules/logger
	github.com/p76081158/5g-nsmf/modules/nsrhandler => ./modules/nsrhandler
	github.com/p76081158/5g-nsmf/modules/optimizer/scheduler => ./modules/optimizer/scheduler
	github.com/p76081158/5g-nsmf/modules/optimizer/slicebinpack => ./modules/optimizer/slicebinpack
	github.com/p76081158/5g-nsmf/modules/optimizer/tenantbinpack => ./modules/optimizer/tenantbinpack
	github.com/p76081158/5g-nsmf/modules/ueransim/gnb => ./modules/ueransim/gnb
	github.com/p76081158/5g-nsmf/modules/ueransim/ue/generator => ./modules/ueransim/ue/generator
)
