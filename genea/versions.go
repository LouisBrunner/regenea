package genea

type Version uint

const (
	VersionV1 Version = 1
	VersionV2 Version = 2
)

type Header struct {
	Version *Version `json:"version" validate:"required,dive,gt=0"`
}
