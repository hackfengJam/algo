IGNORE = ./

PACKAGES := $(filter-out $(IGNORE), $(sort $(dir $(wildcard ./*/))))

FILTER =
LINTS := $(filter-out $(FILTER), $(PACKAGES))

lint: $(LINTS:=.lint)

$(LINTS:=.lint):
	golint -set_exit_status $(subst .lint,,$@)...

unittest:
	go test ./...

.PHONE: lint