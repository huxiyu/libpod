package entities

import (
	"context"

	"github.com/containers/common/pkg/config"
)

type ImageEngine interface {
	Build(ctx context.Context, containerFiles []string, opts BuildOptions) (*BuildReport, error)
	Config(ctx context.Context) (*config.Config, error)
	Diff(ctx context.Context, nameOrId string, options DiffOptions) (*DiffReport, error)
	Exists(ctx context.Context, nameOrId string) (*BoolReport, error)
	History(ctx context.Context, nameOrId string, opts ImageHistoryOptions) (*ImageHistoryReport, error)
	Import(ctx context.Context, opts ImageImportOptions) (*ImageImportReport, error)
	Inspect(ctx context.Context, names []string, opts InspectOptions) (*ImageInspectReport, error)
	List(ctx context.Context, opts ImageListOptions) ([]*ImageSummary, error)
	Load(ctx context.Context, opts ImageLoadOptions) (*ImageLoadReport, error)
	Prune(ctx context.Context, opts ImagePruneOptions) (*ImagePruneReport, error)
	Pull(ctx context.Context, rawImage string, opts ImagePullOptions) (*ImagePullReport, error)
	Push(ctx context.Context, source string, destination string, opts ImagePushOptions) error
	Remove(ctx context.Context, images []string, opts ImageRemoveOptions) (*ImageRemoveReport, error)
	Save(ctx context.Context, nameOrId string, tags []string, options ImageSaveOptions) error
	Search(ctx context.Context, term string, opts ImageSearchOptions) ([]ImageSearchReport, error)
	Shutdown(ctx context.Context)
	Tag(ctx context.Context, nameOrId string, tags []string, options ImageTagOptions) error
	Tree(ctx context.Context, nameOrId string, options ImageTreeOptions) (*ImageTreeReport, error)
	Untag(ctx context.Context, nameOrId string, tags []string, options ImageUntagOptions) error
	ManifestCreate(ctx context.Context, names, images []string, opts ManifestCreateOptions) (string, error)
	ManifestInspect(ctx context.Context, name string) ([]byte, error)
	ManifestAdd(ctx context.Context, opts ManifestAddOptions) (string, error)
}
