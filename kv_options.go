package herald

// KVOption is a functional option for KVGroupWithOpts.
type KVOption func(*kvConfig)

type kvConfig struct {
	separator *string // override theme KVSeparator (nil = use theme default)
	rawKeys   bool    // skip applying KVKey style (keys are pre-styled)
	rawValues bool    // skip applying KVValue style (values are pre-styled)
	indent    int     // left indent in spaces (0 = no indent)
}

// WithKVGroupSeparator overrides the separator between key and value for this
// call only. Pass an empty string to suppress the separator entirely.
func WithKVGroupSeparator(s string) KVOption {
	return func(cfg *kvConfig) { cfg.separator = &s }
}

// WithKVRawKeys skips applying the theme's KVKey style to keys. Use this
// when keys are already styled (e.g. via Var, Bold, or other herald methods).
func WithKVRawKeys(raw bool) KVOption {
	return func(cfg *kvConfig) { cfg.rawKeys = raw }
}

// WithKVRawValues skips applying the theme's KVValue style to values. Use
// this when values are already styled or contain mixed styled content.
func WithKVRawValues(raw bool) KVOption {
	return func(cfg *kvConfig) { cfg.rawValues = raw }
}

// WithKVIndent prepends each line with n spaces of indentation.
func WithKVIndent(n int) KVOption {
	return func(cfg *kvConfig) { cfg.indent = n }
}
