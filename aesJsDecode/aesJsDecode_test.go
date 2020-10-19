/**
 * @Author xiaoxiao
 * @Description CREATE FILE aesJsDecode
 * @Date 2020/10/19 9:12 上午
 **/
package aesJsDecode

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestDecrypt(t *testing.T) {
	type args struct {
		ciphertext []byte
		key        []byte
	}
	bs, err := hex.DecodeString("b4a3031d15d44a33e678cfdc1fb6858b")
	if err != nil {
		t.Errorf("DecryptJs() error = %v", err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: string("test"), args: args{ciphertext: bs, key: []byte("ABCDEF1234123412")}, want: []byte(`"123"`), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.ciphertext, tt.args.key)
			t.Log(string(got),string(tt.want))
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
