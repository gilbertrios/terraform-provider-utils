package provider

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Base64 Encode Function
var _ function.Function = &Base64EncodeFunction{}

type Base64EncodeFunction struct{}

func NewBase64EncodeFunction() function.Function {
	return &Base64EncodeFunction{}
}

func (f *Base64EncodeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base64_encode"
}

func (f *Base64EncodeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Encodes a string to base64",
		Description: "Takes a string and returns its base64 encoded representation.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to encode",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *Base64EncodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	result := base64.StdEncoding.EncodeToString([]byte(input))
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// Base64 Decode Function
var _ function.Function = &Base64DecodeFunction{}

type Base64DecodeFunction struct{}

func NewBase64DecodeFunction() function.Function {
	return &Base64DecodeFunction{}
}

func (f *Base64DecodeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base64_decode"
}

func (f *Base64DecodeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Decodes a base64 string",
		Description: "Takes a base64 encoded string and returns its decoded representation.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The base64 string to decode",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *Base64DecodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(function.NewArgumentFuncError(0, fmt.Sprintf("Invalid base64 string: %s", err)))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, string(decoded)))
}

// SHA256 Function
var _ function.Function = &SHA256Function{}

type SHA256Function struct{}

func NewSHA256Function() function.Function {
	return &SHA256Function{}
}

func (f *SHA256Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "sha256"
}

func (f *SHA256Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Computes SHA256 hash",
		Description: "Takes a string and returns its SHA256 hash as a hexadecimal string.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to hash",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *SHA256Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	hash := sha256.Sum256([]byte(input))
	result := fmt.Sprintf("%x", hash)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// MD5 Function
var _ function.Function = &MD5Function{}

type MD5Function struct{}

func NewMD5Function() function.Function {
	return &MD5Function{}
}

func (f *MD5Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "md5"
}

func (f *MD5Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Computes MD5 hash",
		Description: "Takes a string and returns its MD5 hash as a hexadecimal string.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to hash",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *MD5Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	hash := md5.Sum([]byte(input))
	result := fmt.Sprintf("%x", hash)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// UUIDv4 Function
var _ function.Function = &UUIDv4Function{}

type UUIDv4Function struct{}

func NewUUIDv4Function() function.Function {
	return &UUIDv4Function{}
}

func (f *UUIDv4Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "uuidv4"
}

func (f *UUIDv4Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Generates a deterministic UUID v4",
		Description: "Takes a string and generates a deterministic UUID v4 based on that string using MD5 hashing.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to use as seed for UUID generation",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *UUIDv4Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	hash := md5.Sum([]byte(input))
	// Set version (4) and variant bits according to RFC 4122
	hash[6] = (hash[6] & 0x0f) | 0x40
	hash[8] = (hash[8] & 0x3f) | 0x80

	result := fmt.Sprintf("%x-%x-%x-%x-%x",
		hash[0:4], hash[4:6], hash[6:8], hash[8:10], hash[10:16])
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// Slugify Function
var _ function.Function = &SlugifyFunction{}

type SlugifyFunction struct{}

func NewSlugifyFunction() function.Function {
	return &SlugifyFunction{}
}

func (f *SlugifyFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "slugify"
}

func (f *SlugifyFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Converts a string to a URL-friendly slug",
		Description: "Takes a string and converts it to lowercase, replacing spaces with hyphens and removing special characters.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to slugify",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *SlugifyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	// Convert to lowercase
	result := strings.ToLower(input)
	// Replace spaces with hyphens
	result = strings.ReplaceAll(result, " ", "-")
	// Remove non-alphanumeric characters except hyphens
	reg := regexp.MustCompile("[^a-z0-9-]+")
	result = reg.ReplaceAllString(result, "")
	// Remove duplicate hyphens
	reg = regexp.MustCompile("-+")
	result = reg.ReplaceAllString(result, "-")
	// Trim hyphens from start and end
	result = strings.Trim(result, "-")

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// Truncate Function
var _ function.Function = &TruncateFunction{}

type TruncateFunction struct{}

func NewTruncateFunction() function.Function {
	return &TruncateFunction{}
}

func (f *TruncateFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "truncate"
}

func (f *TruncateFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Truncates a string to a maximum length",
		Description: "Takes a string and a maximum length, returning the truncated string with an optional suffix.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to truncate",
			},
			function.Int64Parameter{
				Name:        "max_length",
				Description: "The maximum length of the result",
			},
			function.StringParameter{
				Name:        "suffix",
				Description: "Optional suffix to add when truncated (e.g., '...')",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *TruncateFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var maxLength int64
	var suffix string

	resp.Error = function.ConcatFuncErrors(
		req.Arguments.Get(ctx, &input, &maxLength, &suffix),
	)
	if resp.Error != nil {
		return
	}

	if maxLength < 0 {
		resp.Error = function.ConcatFuncErrors(function.NewArgumentFuncError(1, "max_length must be non-negative"))
		return
	}

	runes := []rune(input)
	if int64(len(runes)) <= maxLength {
		resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, input))
		return
	}

	suffixLen := int64(len([]rune(suffix)))
	truncateAt := maxLength - suffixLen
	if truncateAt < 0 {
		truncateAt = 0
	}

	result := string(runes[:truncateAt]) + suffix
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// Reverse Function
var _ function.Function = &ReverseFunction{}

type ReverseFunction struct{}

func NewReverseFunction() function.Function {
	return &ReverseFunction{}
}

func (f *ReverseFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "reverse"
}

func (f *ReverseFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Reverses a string",
		Description: "Takes a string and returns it reversed.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to reverse",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ReverseFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	result := string(runes)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// ToUpper Function
var _ function.Function = &ToUpperFunction{}

type ToUpperFunction struct{}

func NewToUpperFunction() function.Function {
	return &ToUpperFunction{}
}

func (f *ToUpperFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "to_upper"
}

func (f *ToUpperFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Converts string to uppercase",
		Description: "Takes a string and returns it in uppercase.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ToUpperFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	result := strings.ToUpper(input)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// ToLower Function
var _ function.Function = &ToLowerFunction{}

type ToLowerFunction struct{}

func NewToLowerFunction() function.Function {
	return &ToLowerFunction{}
}

func (f *ToLowerFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "to_lower"
}

func (f *ToLowerFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Converts string to lowercase",
		Description: "Takes a string and returns it in lowercase.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to convert",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ToLowerFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	result := strings.ToLower(input)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// Trim Function
var _ function.Function = &TrimFunction{}

type TrimFunction struct{}

func NewTrimFunction() function.Function {
	return &TrimFunction{}
}

func (f *TrimFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "trim"
}

func (f *TrimFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Trims whitespace from string",
		Description: "Takes a string and returns it with leading and trailing whitespace removed.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to trim",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *TrimFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))
	if resp.Error != nil {
		return
	}

	result := strings.TrimSpace(input)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// Join Function
var _ function.Function = &JoinFunction{}

type JoinFunction struct{}

func NewJoinFunction() function.Function {
	return &JoinFunction{}
}

func (f *JoinFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "join"
}

func (f *JoinFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Joins a list of strings",
		Description: "Takes a list of strings and a separator, returning them joined together.",
		Parameters: []function.Parameter{
			function.ListParameter{
				Name:        "list",
				Description: "The list of strings to join",
				ElementType: types.StringType,
			},
			function.StringParameter{
				Name:        "separator",
				Description: "The separator to use between elements",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *JoinFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var list []string
	var separator string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &list, &separator))
	if resp.Error != nil {
		return
	}

	result := strings.Join(list, separator)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}

// Split Function
var _ function.Function = &SplitFunction{}

type SplitFunction struct{}

func NewSplitFunction() function.Function {
	return &SplitFunction{}
}

func (f *SplitFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "split"
}

func (f *SplitFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Splits a string into a list",
		Description: "Takes a string and a separator, returning a list of substrings.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "The string to split",
			},
			function.StringParameter{
				Name:        "separator",
				Description: "The separator to split on",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (f *SplitFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var separator string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input, &separator))
	if resp.Error != nil {
		return
	}

	result := strings.Split(input, separator)
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
