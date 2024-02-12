package hunch

var globalCfg = &globalConfig{
	forgetAll:   false,
	ignoreErr:   false,
	onlySuccess: false,
	earlyDone:   false,
	takeVal:     true,
}
var defglobalCfg = &globalConfig{
	forgetAll:   false,
	ignoreErr:   false,
	onlySuccess: false,
	earlyDone:   false,
	takeVal:     true,
}

func resetGlobalCfg() {
	*globalCfg = *defglobalCfg
}

func copyGlobalCfg() (cfg globalConfig) {
	cfg = *globalCfg
	resetGlobalCfg()
	return
}

func SetIgnoreError(ignoreErr bool) {
	globalCfg.ignoreErr = ignoreErr
}

func SetOnlySuccess(onlySuccess bool) {
	globalCfg.onlySuccess = onlySuccess
}

func SetEarlyDone(earlyDone bool) {
	globalCfg.earlyDone = earlyDone
}

func SetTakeValue(takeVal bool) {
	globalCfg.takeVal = takeVal
}

func SetForgetAll(forgetAll bool) {
	globalCfg.forgetAll = forgetAll
}
