//+build faker

package model

import "github.com/jsdidierlaurent/monitoror/models/tiles"

type (
	PortParams struct {
		Hostname string           `json:"hostname" query:"hostname"`
		Port     int              `json:"port" query:"port"`
		Status   tiles.TileStatus `json:"status" query:"status"`
	}
)

func (p *PortParams) Validate() bool {
	return p.Hostname != "" && p.Port != 0
}