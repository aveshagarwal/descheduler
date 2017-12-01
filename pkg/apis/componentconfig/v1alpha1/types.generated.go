/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package v1alpha1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg1_v1.TypeMeta
		var v1 time.Duration
		_, _ = v0, v1
	}
}

func (x *DeschedulerConfiguration) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [8]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = x.DeschedulingInterval != 0
			yyq2[4] = x.PolicyConfigFile != ""
			yyq2[5] = x.DryRun != false
			yyq2[6] = x.NodeSelector != ""
			yyq2[7] = x.EnableProfiling != false
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(8)
			} else {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yym10 := z.EncBinary()
					_ = yym10
					if false {
					} else if z.HasExtensions() && z.EncExt(x.DeschedulingInterval) {
					} else {
						r.EncodeInt(int64(x.DeschedulingInterval))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("deschedulingInterval"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym11 := z.EncBinary()
					_ = yym11
					if false {
					} else if z.HasExtensions() && z.EncExt(x.DeschedulingInterval) {
					} else {
						r.EncodeInt(int64(x.DeschedulingInterval))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.KubeconfigFile))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("kubeconfigFile"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.KubeconfigFile))
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					yym16 := z.EncBinary()
					_ = yym16
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.PolicyConfigFile))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("policyConfigFile"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym17 := z.EncBinary()
					_ = yym17
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.PolicyConfigFile))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[5] {
					yym19 := z.EncBinary()
					_ = yym19
					if false {
					} else {
						r.EncodeBool(bool(x.DryRun))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq2[5] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("dryRun"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym20 := z.EncBinary()
					_ = yym20
					if false {
					} else {
						r.EncodeBool(bool(x.DryRun))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[6] {
					yym22 := z.EncBinary()
					_ = yym22
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.NodeSelector))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[6] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("nodeSelector"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym23 := z.EncBinary()
					_ = yym23
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.NodeSelector))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[7] {
					yym25 := z.EncBinary()
					_ = yym25
					if false {
					} else {
						r.EncodeBool(bool(x.EnableProfiling))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq2[7] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("enableProfiling"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym26 := z.EncBinary()
					_ = yym26
					if false {
					} else {
						r.EncodeBool(bool(x.EnableProfiling))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *DeschedulerConfiguration) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *DeschedulerConfiguration) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "deschedulingInterval":
			if r.TryDecodeAsNil() {
				x.DeschedulingInterval = 0
			} else {
				yyv8 := &x.DeschedulingInterval
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv8) {
				} else {
					*((*int64)(yyv8)) = int64(r.DecodeInt(64))
				}
			}
		case "kubeconfigFile":
			if r.TryDecodeAsNil() {
				x.KubeconfigFile = ""
			} else {
				yyv10 := &x.KubeconfigFile
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*string)(yyv10)) = r.DecodeString()
				}
			}
		case "policyConfigFile":
			if r.TryDecodeAsNil() {
				x.PolicyConfigFile = ""
			} else {
				yyv12 := &x.PolicyConfigFile
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					*((*string)(yyv12)) = r.DecodeString()
				}
			}
		case "dryRun":
			if r.TryDecodeAsNil() {
				x.DryRun = false
			} else {
				yyv14 := &x.DryRun
				yym15 := z.DecBinary()
				_ = yym15
				if false {
				} else {
					*((*bool)(yyv14)) = r.DecodeBool()
				}
			}
		case "nodeSelector":
			if r.TryDecodeAsNil() {
				x.NodeSelector = ""
			} else {
				yyv16 := &x.NodeSelector
				yym17 := z.DecBinary()
				_ = yym17
				if false {
				} else {
					*((*string)(yyv16)) = r.DecodeString()
				}
			}
		case "enableProfiling":
			if r.TryDecodeAsNil() {
				x.EnableProfiling = false
			} else {
				yyv18 := &x.EnableProfiling
				yym19 := z.DecBinary()
				_ = yym19
				if false {
				} else {
					*((*bool)(yyv18)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *DeschedulerConfiguration) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj20 int
	var yyb20 bool
	var yyhl20 bool = l >= 0
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv21 := &x.Kind
		yym22 := z.DecBinary()
		_ = yym22
		if false {
		} else {
			*((*string)(yyv21)) = r.DecodeString()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv23 := &x.APIVersion
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			*((*string)(yyv23)) = r.DecodeString()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.DeschedulingInterval = 0
	} else {
		yyv25 := &x.DeschedulingInterval
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv25) {
		} else {
			*((*int64)(yyv25)) = int64(r.DecodeInt(64))
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.KubeconfigFile = ""
	} else {
		yyv27 := &x.KubeconfigFile
		yym28 := z.DecBinary()
		_ = yym28
		if false {
		} else {
			*((*string)(yyv27)) = r.DecodeString()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.PolicyConfigFile = ""
	} else {
		yyv29 := &x.PolicyConfigFile
		yym30 := z.DecBinary()
		_ = yym30
		if false {
		} else {
			*((*string)(yyv29)) = r.DecodeString()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.DryRun = false
	} else {
		yyv31 := &x.DryRun
		yym32 := z.DecBinary()
		_ = yym32
		if false {
		} else {
			*((*bool)(yyv31)) = r.DecodeBool()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.NodeSelector = ""
	} else {
		yyv33 := &x.NodeSelector
		yym34 := z.DecBinary()
		_ = yym34
		if false {
		} else {
			*((*string)(yyv33)) = r.DecodeString()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.EnableProfiling = false
	} else {
		yyv35 := &x.EnableProfiling
		yym36 := z.DecBinary()
		_ = yym36
		if false {
		} else {
			*((*bool)(yyv35)) = r.DecodeBool()
		}
	}
	for {
		yyj20++
		if yyhl20 {
			yyb20 = yyj20 > l
		} else {
			yyb20 = r.CheckBreak()
		}
		if yyb20 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj20-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}
