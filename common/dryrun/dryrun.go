package dryrun

import "net/http"

// NewDryRun todo
func NewDryRun() *DryRun {
	return &DryRun{
		enable: false,
	}
}

// DryRun todo
type DryRun struct {
	enable bool
}

// EnabeDryRun todo
func (d *DryRun) EnabeDryRun() {
	d.enable = true
}

// IsDryRun todo
func (d *DryRun) IsDryRun() bool {
	return d.enable
}

// GetDryRunParamFromHTTP todo
func (d *DryRun) GetDryRunParamFromHTTP(r *http.Request) {
	qs := r.URL.Query()
	if qs.Get("dry_run") == "true" {
		d.EnabeDryRun()
	}
}
