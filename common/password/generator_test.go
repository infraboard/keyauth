package password

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Generator
		wantErr error
	}{
		{
			name: "default config",
			args: args{nil},
			want: func() *Generator {
				cfg := &DefaultConfig
				cfg.CharacterSet = buildCharacterSet(cfg)
				cfg.Length = LengthStrong
				return &Generator{cfg}
			}(),
		},
		{
			name: "set config",
			args: args{&Config{
				IncludeLowercaseLetters: true,
			}},
			want: func() *Generator {
				cfg := Config{IncludeLowercaseLetters: true}
				cfg.CharacterSet = "abcdefghijklmnopqrstuvwxyz"
				cfg.Length = LengthStrong
				return &Generator{&cfg}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithDefault(t *testing.T) {
	tests := []struct {
		name    string
		want    *Generator
		wantErr error
	}{
		{
			name: "default config",
			want: func() *Generator {
				cfg := &DefaultConfig
				cfg.CharacterSet = buildCharacterSet(cfg)
				return &Generator{cfg}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewWithDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildCharacterSet(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "exclude similar characters",
			args: args{
				config: &Config{
					IncludeLowercaseLetters:    true,
					IncludeSymbols:             true,
					IncludeNumbers:             true,
					ExcludeSimilarCharacters:   true,
					ExcludeAmbiguousCharacters: true,
				},
			},
			want: "abcdefghkmnpqrstuvwxyz23456789!$%^&*_+@#?.-=?",
		},
		{
			name: "exclude numbers",
			args: args{
				config: &Config{
					IncludeLowercaseLetters:    true,
					IncludeSymbols:             true,
					IncludeNumbers:             false,
					ExcludeSimilarCharacters:   true,
					ExcludeAmbiguousCharacters: true,
				},
			},
			want: "abcdefghkmnpqrstuvwxyz!$%^&*_+@#?.-=?",
		},
		{
			name: "full list",
			args: args{
				config: &Config{
					IncludeLowercaseLetters:    true,
					IncludeSymbols:             true,
					IncludeNumbers:             true,
					ExcludeSimilarCharacters:   false,
					ExcludeAmbiguousCharacters: false,
				},
			},
			want: "abcdefghijklmnopqrstuvwxyz0123456789!$%^&*()_+{}:@[];'#<>?,./|\\-=?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildCharacterSet(tt.args.config); got != tt.want {
				t.Errorf("buildCharacterSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_Generate(t *testing.T) {
	type fields struct {
		Config *Config
	}
	tests := []struct {
		name    string
		fields  fields
		test    func(*string, string)
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{&DefaultConfig},
			test: func(pwd *string, characterSet string) {
				assert.Len(t, *pwd, DefaultConfig.Length)
				err := stringMatchesCharacters(*pwd, characterSet)
				if err != nil {
					t.Errorf("Generate() error = %v", err)
					return
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.fields.Config)
			got, err := g.Generate()
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.test(got, g.CharacterSet)
		})
	}
}

func TestGenerator_GenerateMany(t *testing.T) {
	type fields struct {
		Config *Config
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		test    func([]string, string)
		wantErr error
	}{
		{
			name:   "valid",
			args:   args{amount: 5},
			fields: fields{Config: &DefaultConfig},
			test: func(pwds []string, characterSet string) {
				assert.Len(t, pwds, 5)

				for _, pwd := range pwds {
					err := stringMatchesCharacters(pwd, characterSet)
					if err != nil {
						t.Errorf("Generate() error = %v", err)
						return
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.fields.Config)
			got, err := g.GenerateMany(tt.args.amount)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GenerateMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.test(got, g.CharacterSet)
		})
	}
}

func stringMatchesCharacters(str, characters string) error {
	set := strings.Split(characters, "")
	strSet := strings.Split(str, "")

	for _, strChr := range strSet {
		found := false
		for _, setChr := range set {
			if strChr == setChr {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("%v should not be in the str", strChr)
		}
	}

	return nil
}

func TestGenerator_GenerateWithLength(t *testing.T) {
	type fields struct {
		Config *Config
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		test    func(*string, string)
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{&DefaultConfig},
			args:   args{length: 5},
			test: func(pwd *string, characterSet string) {
				assert.Len(t, *pwd, 5)
				err := stringMatchesCharacters(*pwd, characterSet)
				if err != nil {
					t.Errorf("Generate() error = %v", err)
					return
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.fields.Config)
			got, err := g.GenerateWithLength(tt.args.length)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GenerateWithLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.test(got, g.CharacterSet)
		})
	}
}

func TestGenerator_GenerateManyWithLength(t *testing.T) {
	type fields struct {
		Config *Config
	}
	type args struct {
		amount int
		length int
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		test    func([]string, string)
		wantErr error
	}{
		{
			name:   "valid",
			args:   args{amount: 5, length: 5},
			fields: fields{Config: &DefaultConfig},
			test: func(pwds []string, characterSet string) {
				assert.Len(t, pwds, 5)

				for _, pwd := range pwds {
					assert.Len(t, pwd, 5)
					err := stringMatchesCharacters(pwd, characterSet)
					if err != nil {
						t.Errorf("Generate() error = %v", err)
						return
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.fields.Config)

			got, err := g.GenerateManyWithLength(tt.args.amount, tt.args.length)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GenerateManyWithLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.test(got, g.CharacterSet)
		})
	}
}
