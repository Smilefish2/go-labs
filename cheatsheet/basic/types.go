package basic

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"learn-go-with-cli/cmd"
	"math"
	"os"
	"runtime"
	"unsafe"
)

var pointerCmd = &cobra.Command{
	Use:   "cheatsheet:basic:types",
	Short: "Go基本数据类型",
	Run: func(cmd *cobra.Command, args []string) {

		var _bool bool
		_boolType := fmt.Sprintf("%T", _bool)
		_boolSize := fmt.Sprintf("%d", unsafe.Sizeof(_bool))
		_boolDefault := fmt.Sprintf("%v", _bool)

		var _byte byte
		_byteType := "byte"
		_byteSize := fmt.Sprintf("%d", unsafe.Sizeof(_byte))
		_byteDefault := fmt.Sprintf("%v", _byte)

		var _int int
		_intType := fmt.Sprintf("%T", _int)
		_intSize := fmt.Sprintf("%d", unsafe.Sizeof(_int))
		_intDefault := fmt.Sprintf("%v", _int)
		_intRange := fmt.Sprintf("%d ~ %d", math.MinInt, math.MaxInt)

		var _uint uint
		_uintType := fmt.Sprintf("%T", _uint)
		_uintSize := fmt.Sprintf("%d", unsafe.Sizeof(_uint))
		_uintDefault := fmt.Sprintf("%v", _uint)
		var _maxUint uint64 = math.MaxUint64
		_uintRange := fmt.Sprintf("%d ~ %d", 0, _maxUint)

		var _int8 int8
		_int8Type := fmt.Sprintf("%T", _int8)
		_int8Size := fmt.Sprintf("%d", unsafe.Sizeof(_int8))
		_int8Default := fmt.Sprintf("%v", _int8)
		_int8Range := fmt.Sprintf("%d ~ %d", math.MinInt8, math.MaxInt8)

		var _uint8 uint8
		_uint8Type := fmt.Sprintf("%T", _uint8)
		_uint8Size := fmt.Sprintf("%d", unsafe.Sizeof(_uint8))
		_uint8Default := fmt.Sprintf("%v", _uint8)
		_uint8Range := fmt.Sprintf("%d ~ %d", 0, math.MaxUint8)

		var _int16 int16
		_int16Type := fmt.Sprintf("%T", _int16)
		_int16Size := fmt.Sprintf("%d", unsafe.Sizeof(_int16))
		_int16Default := fmt.Sprintf("%v", _int16)
		_int16Range := fmt.Sprintf("%d ~ %d", math.MinInt16, math.MaxInt16)

		var _uint16 uint16
		_uint16Type := fmt.Sprintf("%T", _uint16)
		_uint16Size := fmt.Sprintf("%d", unsafe.Sizeof(_uint16))
		_uint16Default := fmt.Sprintf("%v", _uint16)
		_uint16Range := fmt.Sprintf("%d ~ %d", 0, math.MaxUint16)

		var _int32 int32
		_int32Type := fmt.Sprintf("%T", _int32)
		_int32Size := fmt.Sprintf("%d", unsafe.Sizeof(_int32))
		_int32Default := fmt.Sprintf("%v", _int32)
		_int32Range := fmt.Sprintf("%d ~ %d", math.MinInt32, math.MaxInt32)

		var _uint32 uint32
		_uint32Type := fmt.Sprintf("%T", _uint32)
		_uint32Size := fmt.Sprintf("%d", unsafe.Sizeof(_uint32))
		_uint32Default := fmt.Sprintf("%v", _uint32)
		_uint32Range := fmt.Sprintf("%d ~ %d", 0, math.MaxUint32)

		var _int64 int64
		_int64Type := fmt.Sprintf("%T", _int64)
		_int64Size := fmt.Sprintf("%d", unsafe.Sizeof(_int64))
		_int64Default := fmt.Sprintf("%v", _int64)
		_int64Range := fmt.Sprintf("%d ~ %d", math.MinInt64, math.MaxInt64)

		var _uint64 uint64
		_uint64Type := fmt.Sprintf("%T", _uint64)
		_uint64Size := fmt.Sprintf("%d", unsafe.Sizeof(_uint64))
		_uint64Default := fmt.Sprintf("%v", _uint64)
		var _maxUint64 uint64 = math.MaxUint64
		_uint64Range := fmt.Sprintf("%d ~ %d", 0, _maxUint64)

		var _float32 float32
		_float32Type := fmt.Sprintf("%T", _float32)
		_float32Size := fmt.Sprintf("%d", unsafe.Sizeof(_float32))
		_float32Default := fmt.Sprintf("%v", _float32)
		_float32Range := fmt.Sprintf("%.9E ~ %.9E", math.SmallestNonzeroFloat32, math.MaxFloat32)

		var _float64 float64
		_float64Type := fmt.Sprintf("%T", _float64)
		_float64Size := fmt.Sprintf("%d", unsafe.Sizeof(_float64))
		_float64Default := fmt.Sprintf("%v", _float64)
		_float64Range := fmt.Sprintf("%.9E ~ %.9E", math.SmallestNonzeroFloat64, math.MaxFloat64)

		var _complex64 complex64
		_complex64Type := fmt.Sprintf("%T", _complex64)
		_complex64Size := fmt.Sprintf("%d", unsafe.Sizeof(_complex64))
		_complex64Default := fmt.Sprintf("%v", _complex64)

		var _complex128 complex128
		_complex128Type := fmt.Sprintf("%T", _complex128)
		_complex128Size := fmt.Sprintf("%d", unsafe.Sizeof(_complex128))
		_complex128Default := fmt.Sprintf("%v", _complex128)

		var _rune rune
		_runeType := "rune"
		_runeSize := fmt.Sprintf("%d", unsafe.Sizeof(_rune))
		_runeDefault := fmt.Sprintf("%v", _rune)

		var _uintptr uintptr
		_uintptrType := fmt.Sprintf("%T", _uintptr)
		_uintptrSize := fmt.Sprintf("%d", unsafe.Sizeof(_uintptr))
		_uintptrDefault := fmt.Sprintf("%v", _uintptr)

		var _string string
		_stringType := fmt.Sprintf("%T", _string)
		_stringSize := fmt.Sprintf("%d", unsafe.Sizeof(_string))
		_stringDefault := fmt.Sprintf("\"%v\"", _string)

		var _array [1]int
		_arrayType := "array"
		_arraySize := fmt.Sprintf("%d", unsafe.Sizeof(_array))
		_arrayDefault := fmt.Sprintf("%v", _array)

		type _struct_ struct {
			test  int
			test2 string
			test3 float64
		}
		var _struct _struct_
		_structType := "struct" // fmt.Sprintf("%T", _struct)
		_structSize := fmt.Sprintf("dynamic（%d）", unsafe.Sizeof(_struct))
		_structDefault := fmt.Sprintf("%v", _struct)

		var _func func()
		_funcType := "function" //fmt.Sprintf("%T", _func)
		_funcSize := fmt.Sprintf("%d", unsafe.Sizeof(_func))
		_funcDefault := fmt.Sprintf("%v", _func)

		var _interface interface{}
		_interfaceType := "interface" //fmt.Sprintf("%T", _interface)
		_interfaceSize := fmt.Sprintf("%d", unsafe.Sizeof(_interface))
		_interfaceDefault := fmt.Sprintf("%v", _interface)

		var _map map[string]int
		_mapType := "map" //fmt.Sprintf("%T", _map)
		_mapSize := fmt.Sprintf("%d", unsafe.Sizeof(_map))
		_mapDefault := fmt.Sprintf("%v", _map)

		var _slice []string
		_sliceType := "slice" //fmt.Sprintf("%T", _slice)
		_sliceSize := fmt.Sprintf("%d", unsafe.Sizeof(_slice))
		_sliceDefault := fmt.Sprintf("%v", _slice)

		var _chan chan string
		_chanType := "channel" //fmt.Sprintf("%T", _chan)
		_chanSize := fmt.Sprintf("%d", unsafe.Sizeof(_chan))
		_chanDefault := fmt.Sprintf("%v", _chan)

		data := [][]string{
			{_boolType, _boolSize, _boolDefault, "", ""},
			{_byteType, _byteSize, _byteDefault, "", "byte = uint8"},
			{_intType, "4 or 8 (" + _intSize + ")", _intDefault, _intRange, "x86 = int32, x64 = int64"},
			{_uintType, "4 or 8 (" + _uintSize + ")", _uintDefault, _uintRange, "x86 = uint32, x64 = uint64"},
			{_int8Type, _int8Size, _int8Default, _int8Range, ""},
			{_uint8Type, _uint8Size, _uint8Default, _uint8Range, ""},
			{_int16Type, _int16Size, _int16Default, _int16Range, ""},
			{_uint16Type, _uint16Size, _uint16Default, _uint16Range, ""},
			{_int32Type, _int32Size, _int32Default, _int32Range, ""},
			{_uint32Type, _uint32Size, _uint32Default, _uint32Range, ""},
			{_int64Type, _int64Size, _int64Default, _int64Range, ""},
			{_uint64Type, _uint64Size, _uint64Default, _uint64Range, ""},
			{_float32Type, _float32Size, _float32Default, _float32Range, ""},
			{_float64Type, _float64Size, _float64Default, _float64Range, ""},
			{_complex64Type, _complex64Size, _complex64Default, "", ""},
			{_complex128Type, _complex128Size, _complex128Default, "", ""},
			{_runeType, _runeSize, _runeDefault, "", "rune = int32, unicode code point"},
			{_uintptrType, "4 or 8 (" + _uintptrSize + ")", _uintptrDefault, "", "uintptr = uint"},
			{_stringType, _stringSize, _stringDefault, "", ""},
			{_arrayType, _arraySize, _arrayDefault, "", ""},
			{_structType, _structSize, _structDefault, "", "随意测试：总大小为内部基本数据类型合计大小"},
			{_funcType, _funcSize, _funcDefault, "", ""},
			{_interfaceType, _interfaceSize, _interfaceDefault, "", ""},
			{_mapType, _mapSize, _mapDefault, "", ""},
			{_sliceType, _sliceSize, _sliceDefault, "", ""},
			{_chanType, _chanSize, _chanDefault, "", ""},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Type", "Size", "Default", "Range", "Note"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output

		color.Blue("系统架构:%s,%s,%s", runtime.Compiler, runtime.GOARCH, runtime.GOOS)

	},
}

func init() {
	cmd.RootCmd.AddCommand(pointerCmd)
}
