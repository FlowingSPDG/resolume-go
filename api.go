package resolume

import (
	"context"
	"fmt"
)

// GetParameterByID retrieves a parameter given its unique id
func (c *Client) GetParameterByID(ctx context.Context, parameterID int64) (interface{}, error) {
	endpoint := fmt.Sprintf("/parameter/by-id/%d", parameterID)
	var param interface{}
	if err := c.get(ctx, endpoint, &param); err != nil {
		return nil, err
	}
	return param, nil
}

// SetParameterByID updates a parameter given its unique id
func (c *Client) SetParameterByID(ctx context.Context, parameterID int64, parameter interface{}) error {
	endpoint := fmt.Sprintf("/parameter/by-id/%d", parameterID)
	return c.put(ctx, endpoint, parameter, nil)
}

// ResetParameterByID resets a parameter with the matching unique id
func (c *Client) ResetParameterByID(ctx context.Context, parameterID int64, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/parameter/by-id/%d/reset", parameterID)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// GetComposition retrieves the complete composition
func (c *Client) GetComposition(ctx context.Context) (*Composition, error) {
	endpoint := "/composition"
	var composition Composition
	if err := c.get(ctx, endpoint, &composition); err != nil {
		return nil, err
	}
	return &composition, nil
}

// ReplaceComposition updates the complete composition
func (c *Client) ReplaceComposition(ctx context.Context, composition *Composition) error {
	endpoint := "/composition"
	return c.put(ctx, endpoint, composition, nil)
}

// CompositionAction executes undo or redo actions
func (c *Client) CompositionAction(ctx context.Context, action string) error {
	if action != "undo" && action != "redo" {
		return fmt.Errorf("invalid action: %s (must be 'undo' or 'redo')", action)
	}
	endpoint := "/composition/action"
	return c.post(ctx, endpoint, action, nil)
}

// DisconnectAllClips disconnects all clips in the composition
func (c *Client) DisconnectAllClips(ctx context.Context) error {
	endpoint := "/composition/disconnect-all"
	return c.post(ctx, endpoint, nil, nil)
}

// SetEffectDisplayName changes the display name of an effect
func (c *Client) SetEffectDisplayName(ctx context.Context, effectID int64, displayName string) error {
	endpoint := fmt.Sprintf("/composition/effects/by-id/%d/set-display-name", effectID)
	return c.post(ctx, endpoint, displayName, nil)
}

// MoveEffect moves an effect to the end of the composition
func (c *Client) MoveEffect(ctx context.Context, effectURI string) error {
	endpoint := "/composition/effects/video/move"
	return c.post(ctx, endpoint, effectURI, nil)
}

// MoveEffectToOffset moves an effect to a specific offset in the composition
func (c *Client) MoveEffectToOffset(ctx context.Context, offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/effects/video/move/%d", offset)
	return c.post(ctx, endpoint, effectURI, nil)
}

// AddEffect adds an effect to the entire composition
func (c *Client) AddEffect(ctx context.Context, effectURI string) error {
	endpoint := "/composition/effects/video/add"
	return c.post(ctx, endpoint, effectURI, nil)
}

// AddEffectAtOffset adds an effect to the composition at a specific offset
func (c *Client) AddEffectAtOffset(ctx context.Context, offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/effects/video/add/%d", offset)
	return c.post(ctx, endpoint, effectURI, nil)
}

// DeleteEffect removes an effect from the composition
func (c *Client) DeleteEffect(ctx context.Context, offset int64) error {
	endpoint := fmt.Sprintf("/composition/effects/video/%d", offset)
	return c.delete(ctx, endpoint)
}

// ResetCompositionParameter resets a parameter in the composition to its default value
func (c *Client) ResetCompositionParameter(ctx context.Context, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// GetColumn retrieves column properties by index
func (c *Client) GetColumn(ctx context.Context, columnIndex int64) (*Column, error) {
	endpoint := fmt.Sprintf("/composition/columns/%d", columnIndex)
	var column Column
	if err := c.get(ctx, endpoint, &column); err != nil {
		return nil, err
	}
	return &column, nil
}

// ReplaceColumn updates a specific column by index
func (c *Client) ReplaceColumn(ctx context.Context, columnIndex int64, column *Column) error {
	endpoint := fmt.Sprintf("/composition/columns/%d", columnIndex)
	return c.put(ctx, endpoint, column, nil)
}

// DeleteColumn removes a column by index
func (c *Client) DeleteColumn(ctx context.Context, columnIndex int64) error {
	endpoint := fmt.Sprintf("/composition/columns/%d", columnIndex)
	return c.delete(ctx, endpoint)
}

// DuplicateColumn duplicates the given column
func (c *Client) DuplicateColumn(ctx context.Context, columnIndex int64) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/duplicate", columnIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// AddColumn adds a new column to the composition
func (c *Client) AddColumn(ctx context.Context, beforeColumnURI string) error {
	endpoint := "/composition/columns/add"
	return c.post(ctx, endpoint, beforeColumnURI, nil)
}

// ResetColumnParameter resets a parameter in a column to its default value
func (c *Client) ResetColumnParameter(ctx context.Context, columnIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/%s/reset", columnIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// ConnectColumn connects the column by index
func (c *Client) ConnectColumn(ctx context.Context, columnIndex int64, connect *bool) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/connect", columnIndex)
	return c.post(ctx, endpoint, connect, nil)
}

// SelectColumn selects the column by index
func (c *Client) SelectColumn(ctx context.Context, columnIndex int64) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/select", columnIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// GetLayer retrieves layer properties and clip info by index
func (c *Client) GetLayer(ctx context.Context, layerIndex int64) (*Layer, error) {
	endpoint := fmt.Sprintf("/composition/layers/%d", layerIndex)
	var layer Layer
	if err := c.get(ctx, endpoint, &layer); err != nil {
		return nil, err
	}
	return &layer, nil
}

// ReplaceLayer updates specified layer and/or clips by index
func (c *Client) ReplaceLayer(ctx context.Context, layerIndex int64, layer *Layer) error {
	endpoint := fmt.Sprintf("/composition/layers/%d", layerIndex)
	return c.put(ctx, endpoint, layer, nil)
}

// DeleteLayer removes a layer by index
func (c *Client) DeleteLayer(ctx context.Context, layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d", layerIndex)
	return c.delete(ctx, endpoint)
}

// DuplicateLayer duplicates the given layer
func (c *Client) DuplicateLayer(ctx context.Context, layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/duplicate", layerIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// AddLayer adds a new layer to the composition
func (c *Client) AddLayer(ctx context.Context, beforeLayerURI string) error {
	endpoint := "/composition/layers/add"
	return c.post(ctx, endpoint, beforeLayerURI, nil)
}

// ResetLayerParameter resets a parameter in a layer to its default value
func (c *Client) ResetLayerParameter(ctx context.Context, layerIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/%s/reset", layerIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// SelectLayer selects the layer by index
func (c *Client) SelectLayer(ctx context.Context, layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/select", layerIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// SelectLayerClip selects the clip by index
func (c *Client) SelectLayerClip(ctx context.Context, layerIndex, clipIndex int) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d/select", layerIndex, clipIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// ClearLayer disconnects any playing clips in the layer by index
func (c *Client) ClearLayer(ctx context.Context, layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clear", layerIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// ClearLayerClips clears all clips in the layer by index
func (c *Client) ClearLayerClips(ctx context.Context, layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clearclips", layerIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// GetSelectedLayer retrieves layer properties and clip info for the selected layers
func (c *Client) GetSelectedLayer(ctx context.Context) (*Layer, error) {
	endpoint := "/composition/layers/selected"
	var layer Layer
	if err := c.get(ctx, endpoint, &layer); err != nil {
		return nil, err
	}
	return &layer, nil
}

// ReplaceSelectedLayer updates selected layer and/or clips
func (c *Client) ReplaceSelectedLayer(ctx context.Context, layer *Layer) error {
	endpoint := "/composition/layers/selected"
	return c.put(ctx, endpoint, layer, nil)
}

// DuplicateSelectedLayer duplicates the selected layer
func (c *Client) DuplicateSelectedLayer(ctx context.Context) error {
	endpoint := "/composition/layers/selected/duplicate"
	return c.post(ctx, endpoint, nil, nil)
}

// AddEffectToSelectedLayer adds an effect to the selected layer
func (c *Client) AddEffectToSelectedLayer(ctx context.Context, effectURI string) error {
	endpoint := "/composition/layers/selected/effects/video/add"
	return c.post(ctx, endpoint, effectURI, nil)
}

// AddEffectToSelectedLayerAtOffset adds an effect at the given offset to the selected layer
func (c *Client) AddEffectToSelectedLayerAtOffset(ctx context.Context, offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/layers/selected/effects/video/add/%d", offset)
	return c.post(ctx, endpoint, effectURI, nil)
}

// DeleteSelectedLayerEffect removes an effect from the selected layer
func (c *Client) DeleteSelectedLayerEffect(ctx context.Context, offset int64) error {
	endpoint := fmt.Sprintf("/composition/layers/selected/effects/video/%d", offset)
	return c.delete(ctx, endpoint)
}

// ResetSelectedLayerParameter resets a parameter in the selected layer to its default value
func (c *Client) ResetSelectedLayerParameter(ctx context.Context, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layers/selected/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// ClearSelectedLayer disconnects any playing clips in the selected layer
func (c *Client) ClearSelectedLayer(ctx context.Context) error {
	endpoint := "/composition/layers/selected/clear"
	return c.post(ctx, endpoint, nil, nil)
}

// ClearSelectedLayerClips clears all clips in the selected layer
func (c *Client) ClearSelectedLayerClips(ctx context.Context) error {
	endpoint := "/composition/layers/selected/clearclips"
	return c.post(ctx, endpoint, nil, nil)
}

// GetLayerGroup retrieves layer group properties and layer info by index
func (c *Client) GetLayerGroup(ctx context.Context, layerGroupIndex int64) (*LayerGroup, error) {
	endpoint := fmt.Sprintf("/composition/layergroups/%d", layerGroupIndex)
	var layerGroup LayerGroup
	if err := c.get(ctx, endpoint, &layerGroup); err != nil {
		return nil, err
	}
	return &layerGroup, nil
}

// ReplaceLayerGroup updates specified layer group and/or layers by index
func (c *Client) ReplaceLayerGroup(ctx context.Context, layerGroupIndex int64, layerGroup *LayerGroup) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d", layerGroupIndex)
	return c.put(ctx, endpoint, layerGroup, nil)
}

// DeleteLayerGroup removes a layer group by index
func (c *Client) DeleteLayerGroup(ctx context.Context, layerGroupIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d", layerGroupIndex)
	return c.delete(ctx, endpoint)
}

// DuplicateLayerGroup duplicates the given layer group
func (c *Client) DuplicateLayerGroup(ctx context.Context, layerGroupIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/duplicate", layerGroupIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// MoveLayerToGroup adds an existing layer to an existing layer group
func (c *Client) MoveLayerToGroup(ctx context.Context, layerGroupIndex int64, layerURI string) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/move-layer", layerGroupIndex)
	return c.post(ctx, endpoint, layerURI, nil)
}

// AddLayerToGroup adds a new layer to an existing layer group
func (c *Client) AddLayerToGroup(ctx context.Context, layerGroupIndex int64, beforeLayerURI string) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/add-layer", layerGroupIndex)
	return c.post(ctx, endpoint, beforeLayerURI, nil)
}

// AddLayerGroup adds a new layer group to the composition
func (c *Client) AddLayerGroup(ctx context.Context, beforeLayerOrGroupURI string) error {
	endpoint := "/composition/layergroups/add"
	return c.post(ctx, endpoint, beforeLayerOrGroupURI, nil)
}

// ResetLayerGroupParameter resets a parameter in a layer group to its default value
func (c *Client) ResetLayerGroupParameter(ctx context.Context, layerGroupIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/%s/reset", layerGroupIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// SelectLayerGroup selects the layer group by index
func (c *Client) SelectLayerGroup(ctx context.Context, layerGroupIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/select", layerGroupIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// GetSelectedLayerGroup retrieves selected layer group properties and layer info
func (c *Client) GetSelectedLayerGroup(ctx context.Context) (*LayerGroup, error) {
	endpoint := "/composition/layergroups/selected"
	var layerGroup LayerGroup
	if err := c.get(ctx, endpoint, &layerGroup); err != nil {
		return nil, err
	}
	return &layerGroup, nil
}

// ReplaceSelectedLayerGroup updates selected layer group and/or layers
func (c *Client) ReplaceSelectedLayerGroup(ctx context.Context, layerGroup *LayerGroup) error {
	endpoint := "/composition/layergroups/selected"
	return c.put(ctx, endpoint, layerGroup, nil)
}

// DeleteSelectedLayerGroup removes the selected layer group
func (c *Client) DeleteSelectedLayerGroup(ctx context.Context) error {
	endpoint := "/composition/layergroups/selected"
	return c.delete(ctx, endpoint)
}

// DuplicateSelectedLayerGroup duplicates the selected layer group
func (c *Client) DuplicateSelectedLayerGroup(ctx context.Context) error {
	endpoint := "/composition/layergroups/selected/duplicate"
	return c.post(ctx, endpoint, nil, nil)
}

// MoveLayerToSelectedGroup adds an existing layer to the selected layer group
func (c *Client) MoveLayerToSelectedGroup(ctx context.Context, layerURI string) error {
	endpoint := "/composition/layergroups/selected/move-layer"
	return c.post(ctx, endpoint, layerURI, nil)
}

// AddLayerToSelectedGroup adds a new layer to the selected layer group
func (c *Client) AddLayerToSelectedGroup(ctx context.Context, beforeLayerURI string) error {
	endpoint := "/composition/layergroups/selected/add-layer"
	return c.post(ctx, endpoint, beforeLayerURI, nil)
}

// ResetSelectedLayerGroupParameter resets a parameter in the selected layer group to its default value
func (c *Client) ResetSelectedLayerGroupParameter(ctx context.Context, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layergroups/selected/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// GetDeck retrieves deck properties by index
func (c *Client) GetDeck(ctx context.Context, deckIndex int64) (*Deck, error) {
	endpoint := fmt.Sprintf("/composition/decks/%d", deckIndex)
	var deck Deck
	if err := c.get(ctx, endpoint, &deck); err != nil {
		return nil, err
	}
	return &deck, nil
}

// ReplaceDeck updates a specific deck by index
func (c *Client) ReplaceDeck(ctx context.Context, deckIndex int64, deck *Deck) error {
	endpoint := fmt.Sprintf("/composition/decks/%d", deckIndex)
	return c.put(ctx, endpoint, deck, nil)
}

// DeleteDeck removes a deck by index
func (c *Client) DeleteDeck(ctx context.Context, deckIndex int64) error {
	endpoint := fmt.Sprintf("/composition/decks/%d", deckIndex)
	return c.delete(ctx, endpoint)
}

// DuplicateDeck duplicates the given deck
func (c *Client) DuplicateDeck(ctx context.Context, deckIndex int64) error {
	endpoint := fmt.Sprintf("/composition/decks/%d/duplicate", deckIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// AddDeck adds a new deck to the composition
func (c *Client) AddDeck(ctx context.Context, beforeDeckURI string) error {
	endpoint := "/composition/decks/add"
	return c.post(ctx, endpoint, beforeDeckURI, nil)
}

// ResetDeckParameter resets a parameter in a deck to its default value
func (c *Client) ResetDeckParameter(ctx context.Context, deckIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/decks/%d/%s/reset", deckIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// SelectDeck selects the deck by index
func (c *Client) SelectDeck(ctx context.Context, deckIndex int64) error {
	endpoint := fmt.Sprintf("/composition/decks/%d/select", deckIndex)
	return c.post(ctx, endpoint, nil, nil)
}

// GetDeckByID retrieves deck properties by id
func (c *Client) GetDeckByID(ctx context.Context, deckID int64) (*Deck, error) {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d", deckID)
	var deck Deck
	if err := c.get(ctx, endpoint, &deck); err != nil {
		return nil, err
	}
	return &deck, nil
}

// ReplaceDeckByID updates specific deck by id
func (c *Client) ReplaceDeckByID(ctx context.Context, deckID int64, deck *Deck) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d", deckID)
	return c.put(ctx, endpoint, deck, nil)
}

// DeleteDeckByID removes specified deck by id
func (c *Client) DeleteDeckByID(ctx context.Context, deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d", deckID)
	return c.delete(ctx, endpoint)
}

// DuplicateDeckByID duplicates the given deck
func (c *Client) DuplicateDeckByID(ctx context.Context, deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/duplicate", deckID)
	return c.post(ctx, endpoint, nil, nil)
}

// CloseDeckByID closes the given deck
func (c *Client) CloseDeckByID(ctx context.Context, deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/close", deckID)
	return c.post(ctx, endpoint, nil, nil)
}

// OpenDeckByID re-opens the given deck
func (c *Client) OpenDeckByID(ctx context.Context, deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/open", deckID)
	return c.post(ctx, endpoint, nil, nil)
}

// ResetDeckParameterByID resets a parameter in a deck to its default value
func (c *Client) ResetDeckParameterByID(ctx context.Context, deckID int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/%s/reset", deckID, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// SelectDeckByID selects the deck by id
func (c *Client) SelectDeckByID(ctx context.Context, deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/select", deckID)
	return c.post(ctx, endpoint, nil, nil)
}

// GetClipByPosition retrieves a clip by its position in the clip grid
func (c *Client) GetClipByPosition(ctx context.Context, layerIndex, clipIndex int64) (*Clip, error) {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d", layerIndex, clipIndex)
	var clip Clip
	if err := c.get(ctx, endpoint, &clip); err != nil {
		return nil, err
	}
	return &clip, nil
}

// ReplaceClipByPosition updates clip and/or its effects by position in the clip grid
func (c *Client) ReplaceClipByPosition(ctx context.Context, layerIndex, clipIndex int64, clip *Clip) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d", layerIndex, clipIndex)
	return c.put(ctx, endpoint, clip, nil)
}

// GetSelectedClip retrieves the selected clip
func (c *Client) GetSelectedClip(ctx context.Context) (*Clip, error) {
	endpoint := "/composition/clips/selected"
	var clip Clip
	if err := c.get(ctx, endpoint, &clip); err != nil {
		return nil, err
	}
	return &clip, nil
}

// ReplaceSelectedClip updates selected clip and/or its effects
func (c *Client) ReplaceSelectedClip(ctx context.Context, clip *Clip) error {
	endpoint := "/composition/clips/selected"
	return c.put(ctx, endpoint, clip, nil)
}

// AddEffectToSelectedClip adds an effect to the selected clip
func (c *Client) AddEffectToSelectedClip(ctx context.Context, effectURI string) error {
	endpoint := "/composition/clips/selected/effects/video/add"
	return c.post(ctx, endpoint, effectURI, nil)
}

// AddEffectToSelectedClipAtOffset adds an effect at the given offset to the selected clip
func (c *Client) AddEffectToSelectedClipAtOffset(ctx context.Context, offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/clips/selected/effects/video/add/%d", offset)
	return c.post(ctx, endpoint, effectURI, nil)
}

// DeleteSelectedClipEffect removes an effect from the selected clip
func (c *Client) DeleteSelectedClipEffect(ctx context.Context, offset int64) error {
	endpoint := fmt.Sprintf("/composition/clips/selected/effects/video/%d", offset)
	return c.delete(ctx, endpoint)
}

// ResetSelectedClipParameter resets a parameter in the selected clip to its default value
func (c *Client) ResetSelectedClipParameter(ctx context.Context, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/clips/selected/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(ctx, endpoint, body, nil)
}

// ConnectSelectedClip connects the selected clip
func (c *Client) ConnectSelectedClip(ctx context.Context, connect *bool) error {
	endpoint := "/composition/clips/selected/connect"
	return c.post(ctx, endpoint, connect, nil)
}

// OpenSelectedClip loads a file or opens a source into the selected clip
func (c *Client) OpenSelectedClip(ctx context.Context, uri string) error {
	endpoint := "/composition/clips/selected/open"
	return c.post(ctx, endpoint, uri, nil)
}

// ClearSelectedClip clears the selected clip
func (c *Client) ClearSelectedClip(ctx context.Context) error {
	endpoint := "/composition/clips/selected/clear"
	return c.post(ctx, endpoint, nil, nil)
}

// GetClipByID retrieves a clip by id
func (c *Client) GetClipByID(ctx context.Context, clipID int64) (*Clip, error) {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d", clipID)
	var clip Clip
	if err := c.get(ctx, endpoint, &clip); err != nil {
		return nil, err
	}
	return &clip, nil
}

// ReplaceClipByID updates clip and/or its effects by id
func (c *Client) ReplaceClipByID(ctx context.Context, clipID int64, clip *Clip) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d", clipID)
	return c.put(ctx, endpoint, clip, nil)
}

// SelectClipByID selects the clip by id
func (c *Client) SelectClipByID(ctx context.Context, clipID int64) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/select", clipID)
	return c.post(ctx, endpoint, nil, nil)
}

// ConnectClipByID connects the clip by id
func (c *Client) ConnectClipByID(ctx context.Context, clipID int64, connect *bool) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/connect", clipID)
	return c.post(ctx, endpoint, connect, nil)
}

// OpenClipByID loads a file or opens a source into the clip with the given id
func (c *Client) OpenClipByID(ctx context.Context, clipID int64, uri string) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/open", clipID)
	return c.post(ctx, endpoint, uri, nil)
}

// ClearClipByID clears the clip with the given id
func (c *Client) ClearClipByID(ctx context.Context, clipID int64) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/clear", clipID)
	return c.post(ctx, endpoint, nil, nil)
}
