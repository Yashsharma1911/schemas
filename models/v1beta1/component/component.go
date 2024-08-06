// Package component provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package component

import (
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/meshery/schemas/models/v1alpha1/capability"
	"github.com/meshery/schemas/models/v1beta1/model"
)

// Defines values for ComponentDefinitionCapabilitiesEntityState.
const (
	ComponentDefinitionCapabilitiesEntityStateDeclaration ComponentDefinitionCapabilitiesEntityState = "declaration"
	ComponentDefinitionCapabilitiesEntityStateInstance    ComponentDefinitionCapabilitiesEntityState = "instance"
)

// Defines values for ComponentDefinitionCapabilitiesStatus.
const (
	ComponentDefinitionCapabilitiesStatusDisabled ComponentDefinitionCapabilitiesStatus = "disabled"
	ComponentDefinitionCapabilitiesStatusEnabled  ComponentDefinitionCapabilitiesStatus = "enabled"
)

// Defines values for ComponentDefinitionFormat.
const (
	CUE  ComponentDefinitionFormat = "CUE"
	JSON ComponentDefinitionFormat = "JSON"
)

// Defines values for ComponentDefinitionModelMetadataCapabilitiesEntityState.
const (
	ComponentDefinitionModelMetadataCapabilitiesEntityStateDeclaration ComponentDefinitionModelMetadataCapabilitiesEntityState = "declaration"
	ComponentDefinitionModelMetadataCapabilitiesEntityStateInstance    ComponentDefinitionModelMetadataCapabilitiesEntityState = "instance"
)

// Defines values for ComponentDefinitionModelMetadataCapabilitiesStatus.
const (
	ComponentDefinitionModelMetadataCapabilitiesStatusDisabled ComponentDefinitionModelMetadataCapabilitiesStatus = "disabled"
	ComponentDefinitionModelMetadataCapabilitiesStatusEnabled  ComponentDefinitionModelMetadataCapabilitiesStatus = "enabled"
)

// Defines values for ComponentDefinitionStatus.
const (
	Duplicate ComponentDefinitionStatus = "duplicate"
	Enabled   ComponentDefinitionStatus = "enabled"
	Ignored   ComponentDefinitionStatus = "ignored"
)

// Defines values for ComponentDefinitionStylesBorderStyle.
const (
	Dashed ComponentDefinitionStylesBorderStyle = "dashed"
	Dotted ComponentDefinitionStylesBorderStyle = "dotted"
	Double ComponentDefinitionStylesBorderStyle = "double"
	Solid  ComponentDefinitionStylesBorderStyle = "solid"
)

// Defines values for ComponentDefinitionStylesGhost.
const (
	No  ComponentDefinitionStylesGhost = "no"
	Yes ComponentDefinitionStylesGhost = "yes"
)

// Defines values for ComponentDefinitionStylesShape.
const (
	Barrel               ComponentDefinitionStylesShape = "barrel"
	BottomRoundRectangle ComponentDefinitionStylesShape = "bottom-round-rectangle"
	ConcaveHexagon       ComponentDefinitionStylesShape = "concave-hexagon"
	CutRectangle         ComponentDefinitionStylesShape = "cut-rectangle"
	Diamond              ComponentDefinitionStylesShape = "diamond"
	Ellipse              ComponentDefinitionStylesShape = "ellipse"
	Heptagon             ComponentDefinitionStylesShape = "heptagon"
	Hexagon              ComponentDefinitionStylesShape = "hexagon"
	Octagon              ComponentDefinitionStylesShape = "octagon"
	Pentagon             ComponentDefinitionStylesShape = "pentagon"
	Polygon              ComponentDefinitionStylesShape = "polygon"
	Rectangle            ComponentDefinitionStylesShape = "rectangle"
	Rhomboid             ComponentDefinitionStylesShape = "rhomboid"
	RoundDiamond         ComponentDefinitionStylesShape = "round-diamond"
	RoundHeptagon        ComponentDefinitionStylesShape = "round-heptagon"
	RoundHexagon         ComponentDefinitionStylesShape = "round-hexagon"
	RoundOctagon         ComponentDefinitionStylesShape = "round-octagon"
	RoundPentagon        ComponentDefinitionStylesShape = "round-pentagon"
	RoundRectangle       ComponentDefinitionStylesShape = "round-rectangle"
	RoundTag             ComponentDefinitionStylesShape = "round-tag"
	RoundTriangle        ComponentDefinitionStylesShape = "round-triangle"
	Star                 ComponentDefinitionStylesShape = "star"
	Tag                  ComponentDefinitionStylesShape = "tag"
	Triangle             ComponentDefinitionStylesShape = "triangle"
	Vee                  ComponentDefinitionStylesShape = "vee"
)

// Defines values for ComponentDefinitionStylesTextHalign.
const (
	ComponentDefinitionStylesTextHalignCenter ComponentDefinitionStylesTextHalign = "center"
	ComponentDefinitionStylesTextHalignLeft   ComponentDefinitionStylesTextHalign = "left"
	ComponentDefinitionStylesTextHalignRight  ComponentDefinitionStylesTextHalign = "right"
)

// Defines values for ComponentDefinitionStylesTextTransform.
const (
	Lowercase ComponentDefinitionStylesTextTransform = "lowercase"
	None      ComponentDefinitionStylesTextTransform = "none"
	Uppercase ComponentDefinitionStylesTextTransform = "uppercase"
)

// Defines values for ComponentDefinitionStylesTextValign.
const (
	ComponentDefinitionStylesTextValignBottom ComponentDefinitionStylesTextValign = "bottom"
	ComponentDefinitionStylesTextValignCenter ComponentDefinitionStylesTextValign = "center"
	ComponentDefinitionStylesTextValignTop    ComponentDefinitionStylesTextValign = "top"
)

type Styles struct {
	// ActiveBgColor The colour of the indicator shown when the background is grabbed by the user. Selector needs to be *core*. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	ActiveBgColor *string `json:"active-bg-color,omitempty" yaml:"active-bg-color,omitempty"`

	// ActiveBgOpacity  The opacity of the active background indicator. Selector needs to be *core*.
	ActiveBgOpacity *string `json:"active-bg-opacity,omitempty" yaml:"active-bg-opacity,omitempty"`

	// ActiveBgSize  The opacity of the active background indicator. Selector needs to be *core*.
	ActiveBgSize *string `json:"active-bg-size,omitempty" yaml:"active-bg-size,omitempty"`

	// BackgroundBlacken Blackens the node’s body for values from 0 to 1; whitens the node’s body for values from 0 to -1.
	BackgroundBlacken *float32 `json:"background-blacken,omitempty" yaml:"background-blacken,omitempty"`

	// BackgroundColor The colour of the node’s body. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	BackgroundColor *string `json:"background-color,omitempty" yaml:"background-color,omitempty"`

	// BackgroundOpacity The opacity level of the node’s background colour
	BackgroundOpacity *float32 `json:"background-opacity,omitempty" yaml:"background-opacity,omitempty"`

	// BodyText The text to display for an element’s body. Can give a path, e.g. data(id) will label with the elements id
	BodyText *string `json:"body-text,omitempty" yaml:"body-text,omitempty"`

	// BodyTextBackgroundColor The colour of the node’s body text background. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	BodyTextBackgroundColor *string `json:"body-text-background-color,omitempty" yaml:"body-text-background-color,omitempty"`

	// BodyTextColor The colour of the node’s body text. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	BodyTextColor *string `json:"body-text-color,omitempty" yaml:"body-text-color,omitempty"`

	// BodyTextDecoration A CSS text decoration to be applied to the node’s body text.
	BodyTextDecoration *string `json:"body-text-decoration,omitempty" yaml:"body-text-decoration,omitempty"`

	// BodyTextFontSize The size of the node’s body text.
	BodyTextFontSize *string `json:"body-text-font-size,omitempty" yaml:"body-text-font-size,omitempty"`

	// BodyTextHorizontalAlign A CSS horizontal alignment to be applied to the node’s body text.
	BodyTextHorizontalAlign *string `json:"body-text-horizontal-align,omitempty" yaml:"body-text-horizontal-align,omitempty"`

	// BodyTextMaxWidth The maximum width for wrapping text in the node.
	BodyTextMaxWidth *string `json:"body-text-max-width,omitempty" yaml:"body-text-max-width,omitempty"`

	// BodyTextOpacity The opacity of the node’s body text, including its outline.
	BodyTextOpacity *float32 `json:"body-text-opacity,omitempty" yaml:"body-text-opacity,omitempty"`

	// BodyTextVerticalAlign A CSS vertical alignment to be applied to the node’s body text.
	BodyTextVerticalAlign *string `json:"body-text-vertical-align,omitempty" yaml:"body-text-vertical-align,omitempty"`

	// BodyTextWeight A CSS font weight to be applied to the node’s body text.
	BodyTextWeight *string `json:"body-text-weight,omitempty" yaml:"body-text-weight,omitempty"`

	// BodyTextWrap How to wrap the text in the node. Can be 'none', 'wrap', or 'ellipsis'.
	BodyTextWrap *string `json:"body-text-wrap,omitempty" yaml:"body-text-wrap,omitempty"`

	// BorderColor The colour of the node’s border. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	BorderColor *string `json:"border-color,omitempty" yaml:"border-color,omitempty"`

	// BorderOpacity The opacity of the node’s border
	BorderOpacity *float32 `json:"border-opacity,omitempty" yaml:"border-opacity,omitempty"`

	// BorderStyle The style of the node’s border
	BorderStyle *ComponentDefinitionStylesBorderStyle `json:"border-style,omitempty" yaml:"border-style,omitempty"`

	// BorderWidth The size of the node’s border.
	BorderWidth *float32 `json:"border-width,omitempty" yaml:"border-width,omitempty"`

	// Color The color of the element's label. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	Color *string `json:"color,omitempty" yaml:"color,omitempty"`

	// FontFamily A comma-separated list of font names to use on the label text.
	FontFamily *string `json:"font-family,omitempty" yaml:"font-family,omitempty"`

	// FontSize The size of the label text.
	FontSize *string `json:"font-size,omitempty" yaml:"font-size,omitempty"`

	// FontStyle A CSS font style to be applied to the label text.
	FontStyle *string `json:"font-style,omitempty" yaml:"font-style,omitempty"`

	// FontWeight A CSS font weight to be applied to the label text.
	FontWeight *string `json:"font-weight,omitempty" yaml:"font-weight,omitempty"`

	// Ghost Whether to use the ghost effect, a semitransparent duplicate of the element drawn at an offset.
	Ghost *ComponentDefinitionStylesGhost `json:"ghost,omitempty" yaml:"ghost,omitempty"`

	// Height The height of the node’s body
	Height *float32 `json:"height,omitempty" yaml:"height,omitempty"`

	// Label The text to display for an element’s label. Can give a path, e.g. data(id) will label with the elements id
	Label *string `json:"label,omitempty" yaml:"label,omitempty"`

	// Opacity The opacity of the element, ranging from 0 to 1. Note that the opacity of a compound node parent affects the effective opacity of its children.See https://js.cytoscape.org/#style/visibility
	Opacity *float32 `json:"opacity,omitempty" yaml:"opacity,omitempty"`

	// OutsideTextureBgColor The colour of the area outside the viewport texture when initOptions.textureOnViewport === true.  Selector needs to be *core*. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	OutsideTextureBgColor *string `json:"outside-texture-bg-color,omitempty" yaml:"outside-texture-bg-color,omitempty"`

	// OutsideTextureBgOpacity The opacity of the area outside the viewport texture. Selector needs to be *core*
	OutsideTextureBgOpacity *float32 `json:"outside-texture-bg-opacity,omitempty" yaml:"outside-texture-bg-opacity,omitempty"`

	// Padding The amount of padding around all sides of the node.
	Padding *float32 `json:"padding,omitempty" yaml:"padding,omitempty"`

	// Position The position of the node. If the position is set, the node is drawn at that position in the given dimensions. If the position is not set, the node is drawn at a random position.
	Position *struct {
		// X The x-coordinate of the node.
		X float32 `json:"x" yaml:"x"`

		// Y The y-coordinate of the node.
		Y float32 `json:"y" yaml:"y"`
	} `json:"position,omitempty" yaml:"position,omitempty"`

	// PrimaryColor Primary color of the component used for UI representation.
	PrimaryColor string `json:"primaryColor" yaml:"primaryColor"`

	// SecondaryColor Secondary color of the entity used for UI representation.
	SecondaryColor *string `json:"secondaryColor,omitempty" yaml:"secondaryColor,omitempty"`

	// SelectionBoxBorderWidth The size of the border on the selection box. Selector needs to be *core*
	SelectionBoxBorderWidth *float32 `json:"selection-box-border-width,omitempty" yaml:"selection-box-border-width,omitempty"`

	// SelectionBoxColor The background colour of the selection box used for drag selection. Selector needs to be *core*. Colours may be specified by name (e.g. red), hex (e.g. #ff0000 or #f00), RGB (e.g. rgb(255, 0, 0)), or HSL (e.g. hsl(0, 100%, 50%)).
	SelectionBoxColor *string `json:"selection-box-color,omitempty" yaml:"selection-box-color,omitempty"`

	// SelectionBoxOpacity The opacity of the selection box.  Selector needs to be *core*
	SelectionBoxOpacity *float32 `json:"selection-box-opacity,omitempty" yaml:"selection-box-opacity,omitempty"`

	// Shape The shape of the node’s body. Note that each shape fits within the specified width and height, and so you may have to adjust width and height if you desire an equilateral shape (i.e. width !== height for several equilateral shapes)
	Shape *ComponentDefinitionStylesShape `json:"shape,omitempty" yaml:"shape,omitempty"`

	// SvgColor Colored SVG of the entity used for UI representation on light background.
	SvgColor string `json:"svgColor" yaml:"svgColor"`

	// SvgComplete Complete SVG of the entity used for UI representation, often inclusive of background.
	SvgComplete string `json:"svgComplete,omitempty" yaml:"svgComplete,omitempty"`

	// SvgWhite White SVG of the entity used for UI representation on dark background.
	SvgWhite string `json:"svgWhite" yaml:"svgWhite"`

	// TextHalign The horizontal alignment of a node’s label
	TextHalign *ComponentDefinitionStylesTextHalign `json:"text-halign,omitempty" yaml:"text-halign,omitempty"`

	// TextOpacity The opacity of the label text, including its outline.
	TextOpacity *float32 `json:"text-opacity,omitempty" yaml:"text-opacity,omitempty"`

	// TextTransform A transformation to apply to the label text
	TextTransform *ComponentDefinitionStylesTextTransform `json:"text-transform,omitempty" yaml:"text-transform,omitempty"`

	// TextValign The vertical alignment of a node’s label
	TextValign *ComponentDefinitionStylesTextValign `json:"text-valign,omitempty" yaml:"text-valign,omitempty"`

	// Width The width of the node’s body or the width of an edge’s line.
	Width *float32 `json:"width,omitempty" yaml:"width,omitempty"`

	// ZIndex An integer value that affects the relative draw order of elements. In general, an element with a higher z-index will be drawn on top of an element with a lower z-index. Note that edges are under nodes despite z-index.
	ZIndex *int `json:"z-index,omitempty" yaml:"z-index,omitempty"`
}

type Component struct {
	// Kind The unique identifier (name) assigned by the registrant to this component. Example: A Kubernetes Pod is of kind 'Pod'.
	Kind string `json:"kind" yaml:"kind"`

	// Schema JSON schema of the object as defined by the registrant.
	Schema string `json:"schema" yaml:"schema"`

	// Version Version of the component produced by the registrant. Example: APIVersion of a Kubernetes Pod.
	Version string `json:"version" yaml:"version"`
}

// ComponentDefinition Components are reusable building blocks for depicting capabilities defined within models. Learn more at https://docs.meshery.io/concepts/components
type ComponentDefinition struct {
	// Capabilities Meshery manages components in accordance with their specific capabilities. This field explicitly identifies those capabilities largely by what actions a given component supports; e.g. metric-scrape, sub-interface, and so on. This field is extensible. ComponentDefinitions may define a broad array of capabilities, which are in-turn dynamically interpretted by Meshery for full lifecycle management.
	Capabilities *[]capability.Capability `json:"capabilities,omitempty" yaml:"capabilities" gorm:"type:bytes;serializer:json"`

	// Component Component and it's properties.
	Component Component `gorm:"type:bytes;serializer:json" json:"component" yaml:"component"`

	// Configuration The configuration of the component. The configuration is based on the schema defined within the component definition(component.schema).
	Configuration map[string]interface{} `gorm:"type:bytes;serializer:json" json:"configuration" yaml:"configuration"`

	// Description A written representation of the purpose and characteristics of the component.
	Description string `json:"description" yaml:"description"`

	// DisplayName Name of the component in human-readible format.
	DisplayName string `json:"displayName" yaml:"displayName"`

	// Format Format specifies the format used in the `component.schema` field. JSON is the default.
	Format ComponentDefinitionFormat `json:"format" yaml:"format"`

	// Id Uniquely identifies the entity (i.e. component) as defined in a declaration (i.e. design).
	Id uuid.UUID `json:"id" yaml:"id"`

	// Metadata Metadata contains additional information associated with the component.
	Metadata ComponentDefinition_Metadata `gorm:"type:bytes;serializer:json" json:"metadata,omitempty" yaml:"metadata"`

	// Model Reference to the specific registered model to which the component belongs and from which model version, category, and other properties may be referenced. Learn more at https://docs.meshery.io/concepts/models
	Model model.ModelDefinition `gorm:"foreignKey:ModelId;references:Id" json:"model" yaml:"model"`

	ModelId uuid.UUID `json:"-" gorm:"index:idx_component_definition_dbs_model_id,column:model_id" yaml:"-"`

	// SchemaVersion Specifies the version of the schema to which the component definition conforms.
	SchemaVersion string `json:"schemaVersion" yaml:"schemaVersion"`

	// Status Status of component, including:
	// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
	// - maintenance: model is unavailable for a period of time.
	// - enabled: model is available for use for all users of this Meshery Server.
	// - ignored: model is unavailable for use for all users of this Meshery Server.
	Status ComponentDefinitionStatus `json:"status" yaml:"status"`

	// Styles Visualization styles for a component
	Styles *Styles `json:"styles" yaml:"styles" gorm:"type:bytes;serializer:json"`

	// Version Version of the component definition.
	Version string `json:"version" yaml:"version"`
}

// ComponentDefinitionCapabilitiesEntityState defines model for ComponentDefinition.Capabilities.EntityState.
type ComponentDefinitionCapabilitiesEntityState string

// ComponentDefinitionCapabilitiesStatus Status of the capability
type ComponentDefinitionCapabilitiesStatus string

// ComponentDefinitionFormat Format specifies the format used in the `component.schema` field. JSON is the default.
type ComponentDefinitionFormat string

// ComponentDefinition_Metadata Metadata contains additional information associated with the component.
type ComponentDefinition_Metadata struct {
	// Genealogy Genealogy represents the various representational states of the component.
	Genealogy string `json:"genealogy" yaml:"genealogy"`

	// IsAnnotation Identifies whether the component is semantically meaningful or not; identifies whether the component should be treated as deployable entity or is for purposes of logical representation.
	IsAnnotation bool `json:"isAnnotation" yaml:"isAnnotation"`

	// Published 'published' controls whether the component should be registered in Meshery Registry. When the same 'published' property in Models, is set to 'false', the Model property takes precedence with all Entities in the Model not being registered.
	Published            bool                   `json:"published" yaml:"published"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ComponentDefinitionModelMetadataCapabilitiesEntityState defines model for ComponentDefinition.Model.Metadata.Capabilities.EntityState.
type ComponentDefinitionModelMetadataCapabilitiesEntityState string

// ComponentDefinitionModelMetadataCapabilitiesStatus Status of the capability
type ComponentDefinitionModelMetadataCapabilitiesStatus string

// ComponentDefinitionStatus Status of component, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type ComponentDefinitionStatus string

// ComponentDefinitionStylesBorderStyle The style of the node’s border
type ComponentDefinitionStylesBorderStyle string

// ComponentDefinitionStylesGhost Whether to use the ghost effect, a semitransparent duplicate of the element drawn at an offset.
type ComponentDefinitionStylesGhost string

// ComponentDefinitionStylesShape The shape of the node’s body. Note that each shape fits within the specified width and height, and so you may have to adjust width and height if you desire an equilateral shape (i.e. width !== height for several equilateral shapes)
type ComponentDefinitionStylesShape string

// ComponentDefinitionStylesTextHalign The horizontal alignment of a node’s label
type ComponentDefinitionStylesTextHalign string

// ComponentDefinitionStylesTextTransform A transformation to apply to the label text
type ComponentDefinitionStylesTextTransform string

// ComponentDefinitionStylesTextValign The vertical alignment of a node’s label
type ComponentDefinitionStylesTextValign string

// Getter for additional properties for ComponentDefinition_Metadata. Returns the specified
// element and whether it was found
func (a ComponentDefinition_Metadata) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ComponentDefinition_Metadata
func (a *ComponentDefinition_Metadata) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ComponentDefinition_Metadata to handle AdditionalProperties
func (a *ComponentDefinition_Metadata) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["genealogy"]; found {
		err = json.Unmarshal(raw, &a.Genealogy)
		if err != nil {
			return fmt.Errorf("error reading 'genealogy': %w", err)
		}
		delete(object, "genealogy")
	}

	if raw, found := object["isAnnotation"]; found {
		err = json.Unmarshal(raw, &a.IsAnnotation)
		if err != nil {
			return fmt.Errorf("error reading 'isAnnotation': %w", err)
		}
		delete(object, "isAnnotation")
	}

	if raw, found := object["published"]; found {
		err = json.Unmarshal(raw, &a.Published)
		if err != nil {
			return fmt.Errorf("error reading 'published': %w", err)
		}
		delete(object, "published")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ComponentDefinition_Metadata to handle AdditionalProperties
func (a ComponentDefinition_Metadata) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["genealogy"], err = json.Marshal(a.Genealogy)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'genealogy': %w", err)
	}

	object["isAnnotation"], err = json.Marshal(a.IsAnnotation)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'isAnnotation': %w", err)
	}

	object["published"], err = json.Marshal(a.Published)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'published': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}