package defaults

import (
	"crypto/md5" // nolint: gosec
	"encoding/hex"
	"math/rand/v2"
	"reflect"
	"sync"
	"time"
)

// Factory ...
func Factory(variable any) {
	getFactoryFiller().Fill(variable)
}

var (
	factoryFillerOnce sync.Once
	factoryFiller     *Filler
)

func getFactoryFiller() *Filler {
	factoryFillerOnce.Do(func() {
		factoryFiller = newFactoryFiller()
	})

	return factoryFiller
}

func newFactoryFiller() *Filler {
	// nolint: gosec
	r := rand.New(rand.NewPCG(1, 2))

	funcs := make(map[reflect.Kind]FillerFunc, 0)

	funcs[reflect.Bool] = func(field *FieldData) {
		if r.IntN(2) == 1 {
			field.Value.SetBool(true)
		} else {
			field.Value.SetBool(false)
		}
	}

	funcs[reflect.Int] = func(field *FieldData) {
		field.Value.SetInt(int64(r.Int()))
	}

	funcs[reflect.Int8] = funcs[reflect.Int]
	funcs[reflect.Int16] = funcs[reflect.Int]
	funcs[reflect.Int32] = funcs[reflect.Int]
	funcs[reflect.Int64] = funcs[reflect.Int]

	funcs[reflect.Float32] = func(field *FieldData) {
		field.Value.SetFloat(r.Float64())
	}

	funcs[reflect.Float64] = funcs[reflect.Float32]

	funcs[reflect.Uint] = func(field *FieldData) {
		field.Value.SetUint(uint64(r.Uint32()))
	}

	funcs[reflect.Uint8] = funcs[reflect.Uint]
	funcs[reflect.Uint16] = funcs[reflect.Uint]
	funcs[reflect.Uint32] = funcs[reflect.Uint]
	funcs[reflect.Uint64] = funcs[reflect.Uint]

	funcs[reflect.String] = func(field *FieldData) {
		field.Value.SetString(randomString())
	}

	funcs[reflect.Slice] = func(field *FieldData) {
		if field.Value.Type().Elem().Kind() == reflect.Uint8 {
			if field.Value.Bytes() != nil {
				return
			}

			field.Value.SetBytes([]byte(randomString()))
		}
	}

	funcs[reflect.Struct] = func(field *FieldData) {
		fields := getFactoryFiller().GetFieldsFromValue(field.Value, nil)
		getFactoryFiller().SetDefaultValues(fields)
	}

	return &Filler{FuncByKind: funcs, Tag: "factory"}
}

func randomString() string {
	// nolint: gosec
	hash := md5.Sum([]byte(time.Now().UTC().String()))
	return hex.EncodeToString(hash[:])
}
