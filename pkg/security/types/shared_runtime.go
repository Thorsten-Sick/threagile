/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package types

type SharedRuntime struct {
	Id                     string   `json:"id,omitempty" yaml:"id,omitempty"`
	Title                  string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description            string   `json:"description,omitempty" yaml:"description,omitempty"`
	Tags                   []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	TechnicalAssetsRunning []string `json:"technical_assets_running,omitempty" yaml:"technical_assets_running,omitempty"`
}

func (what SharedRuntime) IsTaggedWithAny(tags ...string) bool {
	return containsCaseInsensitiveAny(what.Tags, tags...)
}

func (what SharedRuntime) HighestConfidentiality(model *Model) Confidentiality {
	highest := Public
	for _, id := range what.TechnicalAssetsRunning {
		techAsset := model.TechnicalAssets[id]
		if techAsset.HighestProcessedConfidentiality(model) > highest {
			highest = techAsset.HighestProcessedConfidentiality(model)
		}
	}
	return highest
}

func (what SharedRuntime) HighestIntegrity(model *Model) Criticality {
	highest := Archive
	for _, id := range what.TechnicalAssetsRunning {
		techAsset := model.TechnicalAssets[id]
		if techAsset.HighestProcessedIntegrity(model) > highest {
			highest = techAsset.HighestProcessedIntegrity(model)
		}
	}
	return highest
}

func (what SharedRuntime) HighestAvailability(model *Model) Criticality {
	highest := Archive
	for _, id := range what.TechnicalAssetsRunning {
		techAsset := model.TechnicalAssets[id]
		if techAsset.HighestProcessedAvailability(model) > highest {
			highest = techAsset.HighestProcessedAvailability(model)
		}
	}
	return highest
}

type BySharedRuntimeTitleSort []*SharedRuntime

func (what BySharedRuntimeTitleSort) Len() int      { return len(what) }
func (what BySharedRuntimeTitleSort) Swap(i, j int) { what[i], what[j] = what[j], what[i] }
func (what BySharedRuntimeTitleSort) Less(i, j int) bool {
	return what[i].Title < what[j].Title
}
