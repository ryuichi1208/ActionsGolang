func writeFormat(w io.Writer, s stackCounter, f Format, hz int) error {
	switch f {
	case FormatFolded:
		return writeFolded(w, s)
	case FormatPprof:
		return toPprof(s, hz).Write(w)
	default:
		return fmt.Errorf("unknown format: %q", f)
	}
}
