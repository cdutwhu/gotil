package rflx

import "reflect"

// TryInvoke : func Name must be Exportable
func TryInvoke(st interface{}, name string, args ...interface{}) (rets []interface{}, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			failPnOnErr(5, "%v", fEf("%v", r))
		}
	}()

	stVal := vof(st)
	failP1OnErrWhen(stVal.Kind() != typPTR || stVal.Elem().Kind() != typSTRUCT, "%v", fEf("PARAM_INVALID_STRUCT_PTR"))

	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = vof(args[i])
	}
	if _, ok := stVal.Type().MethodByName(name); ok {
		for _, ret := range stVal.MethodByName(name).Call(inputs) {
			rets = append(rets, ret.Interface())
		}
		return rets, true
	}
	return rets, false
}

// MustInvokeWithMW :
func MustInvokeWithMW(st interface{}, name string, args ...interface{}) []interface{} {
	rets, ok := TryInvokeWithMW(st, name, args...)
	failP1OnErrWhen(!ok, "%v: No [%s]", fEf("INTERNAL"), name)
	return rets
}

// TryInvokeWithMW : func Name must be Exportable
func TryInvokeWithMW(st interface{}, name string, args ...interface{}) (rets []interface{}, ok bool) {
	for k, v := range Struct2Map(st) {
		// fPln(k, v)
		if k == "MW" || k == "MiddleWare" || k == "MIDDLEWARE" {
			if mMW, ok := v.(map[string]map[string][]interface{}); ok {
			NEXTFN:
				for fn, mCallerParams := range mMW {
					for _, caller := range []string{name, "*"} {
						if params, ok := mCallerParams[caller]; ok {
							// "$1" -> args[0] etc... ; "$@" -> args string
							for i, param := range params {
								if paramStr, ok := param.(string); ok {
									if repParam.MatchString(paramStr) {
										num, err := scParseUint(paramStr[1:], 10, 64)
										failOnErr("%v", err)
										failOnErrWhen(int(num) > len(args) || int(num) < 0, "MiddleWare: %v", fEf("PARAM_INVALID_INDEX"))
										if num == 0 {
											params[i] = name
										} else {
											params[i] = args[num-1]
										}
									} else if paramStr == "$@" {
										argStrs := make([]string, len(args))
										for i, arg := range args {
											argStrs[i] = fSf("%v", arg)
										}
										params[i] = sJoin(argStrs, " ")
									}
								}
							}
							TryInvoke(st, fn, params...)
							continue NEXTFN
						}
					}
				}
			}
		}
	}
	return TryInvoke(st, name, args...)
}

// InvokeRst :
func InvokeRst(rets interface{}, idx int) interface{} {
	slc, ok := rets.([]interface{})
	failP1OnErrWhen(!ok, "%v", fEf("PARAM_INVALID"))
	failP1OnErrWhen(idx >= len(slc), "%v", fEf("PARAM_INVALID_INDEX"))
	return slc[idx]
}
