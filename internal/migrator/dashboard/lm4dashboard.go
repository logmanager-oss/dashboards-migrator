package dashboard

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/logmanager-oss/dashboards-migrator/internal/migrator/visualization"
	"github.com/logmanager-oss/dashboards-migrator/internal/objects"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm3"
	"github.com/logmanager-oss/dashboards-migrator/internal/types/lm4"
)

type LM4Dashboard struct {
	savedObject         *lm4.SavedObject
	search              *lm4.SearchSourceJSON
	Panels              []lm4.PanelJSON
	id                  string
	currentGridPosition struct {
		X int
		Y int
	}
	references       []lm4.Reference
	currentRefNumber int
}

func NewLM4Dashboard() *LM4Dashboard {
	return &LM4Dashboard{
		savedObject: objects.GetDefaultDashboardSavedObject(),
		search:      objects.GetDefaultSearchObject(false),
		id:          uuid.New().String(),
	}
}

func (dashboard *LM4Dashboard) MigrateDashboard(title string, filters []lm3.GlobalFilter, indexPattern string) (*lm4.SavedObject, error) {
	gridRaw, err := json.Marshal(dashboard.Panels)
	if err != nil {
		return nil, err
	}

	dashboard.search.Filter = dashboard.migrateFilters(filters, indexPattern)
	searchRaw, err := json.Marshal(dashboard.search)
	if err != nil {
		return nil, err
	}

	dashboard.savedObject.Attributes.Title = title
	dashboard.savedObject.Attributes.PanelsJSON = string(gridRaw)
	dashboard.savedObject.Attributes.KibanaSavedObjectMeta.SearchSourceJSON = string(searchRaw)
	dashboard.savedObject.References = dashboard.references

	return dashboard.savedObject, nil
}

func (dashboard *LM4Dashboard) LinkVisualizationToDashboardObject(params *visualization.MigrationParams, visualizationType string) {
	grid := dashboard.calculateGridPosition(params.Span)

	panel := dashboard.buildPanelObject(
		grid,
		params.ID,
		params.Title,
	)

	dashboard.Panels = append(dashboard.Panels, *panel)

	ref := dashboard.buildVisualizationReferenceObject(
		params.ID,
		visualizationType,
	)

	dashboard.references = append(dashboard.references, *ref)

	dashboard.currentRefNumber++
}

func (dashboard *LM4Dashboard) calculateGridPosition(span int) *lm4.GridData {
	width := span * 4

	if dashboard.currentGridPosition.X == objects.MaxRowWidth || dashboard.currentGridPosition.X+width > objects.MaxRowWidth {
		dashboard.currentGridPosition.X = 0
		dashboard.currentGridPosition.Y += objects.DefaultVisHeight
	}

	gridData := &lm4.GridData{
		X: dashboard.currentGridPosition.X,
		Y: dashboard.currentGridPosition.Y,
	}

	gridData.W = width
	gridData.H = objects.DefaultVisHeight
	dashboard.currentGridPosition.X += gridData.W

	return gridData
}

func (dashboard *LM4Dashboard) buildPanelObject(grid *lm4.GridData, id string, title string) *lm4.PanelJSON {
	panel := objects.GetDefaultPanelObject()
	panel.GridData = *grid
	panel.GridData.I = id
	panel.PanelIndex = id
	panel.Title = title
	panel.PanelRefName = fmt.Sprintf("panel_%d", dashboard.currentRefNumber)

	return panel
}

func (dashboard *LM4Dashboard) buildVisualizationReferenceObject(id string, refType string) *lm4.Reference {
	ref := objects.GetDefaultReferenceObject()
	ref.Name = fmt.Sprintf("panel_%d", dashboard.currentRefNumber)
	ref.ID = id
	ref.Type = refType

	return ref
}

func (dashboard *LM4Dashboard) migrateFilters(lm3filters []lm3.GlobalFilter, indexPattern string) []lm4.GlobalFilter {
	var lm4filters []lm4.GlobalFilter

	for _, lm3filter := range lm3filters {
		indexRefName := fmt.Sprintf("kibanaSavedObjectMeta.searchSourceJSON.filter[%d].meta.index", dashboard.currentRefNumber)
		lm4filter := objects.GetDefaultFilterObject(indexRefName)

		if lm3filter.Alias != "" {
			lm4filter.Meta.Alias = lm3filter.Alias
		}
		if lm3filter.Mandate == "mustNot" {
			lm4filter.Meta.Negate = true
		}
		if !lm3filter.Active {
			lm4filter.Meta.Disabled = true
		}

		if lm3filter.Type == "querystring" {
			lm4filter.Query.QueryString = map[string]string{"query": lm3filter.Query}
			lm4filter.Meta.Type = "query_string"
			lm4filter.Meta.Key = "query"
		} else {
			lm4filter.Meta.Params = struct {
				Query string "json:\"query\""
			}{
				Query: "",
			}

			if lm3filter.Query != "" {
				lm4filter.Meta.Params.Query = lm3filter.Query
			} else {
				lm4filter.Meta.Params.Query = lm3filter.Value
			}

			lm4filter.Query.MatchPhrase = map[string]string{lm3filter.Field: lm4filter.Meta.Params.Query}

			lm4filter.Meta.Type = "phrase"
			lm4filter.Meta.Key = lm3filter.Field
		}

		lm4filters = append(lm4filters, *lm4filter)

		ref := dashboard.buildFilterReferenceObject(indexPattern, indexRefName)
		dashboard.references = append(dashboard.references, *ref)
		dashboard.currentRefNumber++
	}

	return lm4filters
}

func (dashboard *LM4Dashboard) buildFilterReferenceObject(id string, indexRefName string) *lm4.Reference {
	ref := objects.GetDefaultReferenceObject()
	ref.Name = indexRefName
	ref.ID = id
	ref.Type = "index-pattern"

	return ref
}
