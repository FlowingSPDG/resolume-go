package resolume

import "fmt"

// GetParameterByID retrieves a parameter given its unique id
func (c *Client) GetParameterByID(parameterID int64) (interface{}, error) {
	endpoint := fmt.Sprintf("/parameter/by-id/%d", parameterID)
	var param interface{}
	if err := c.get(endpoint, &param); err != nil {
		return nil, err
	}
	return param, nil
}

// SetParameterByID updates a parameter given its unique id
func (c *Client) SetParameterByID(parameterID int64, parameter interface{}) error {
	endpoint := fmt.Sprintf("/parameter/by-id/%d", parameterID)
	return c.put(endpoint, parameter, nil)
}

// ResetParameterByID resets a parameter with the matching unique id
func (c *Client) ResetParameterByID(parameterID int64, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/parameter/by-id/%d/reset", parameterID)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// GetComposition retrieves the complete composition
func (c *Client) GetComposition() (*Composition, error) {
	endpoint := "/composition"
	var composition Composition
	if err := c.get(endpoint, &composition); err != nil {
		return nil, err
	}
	return &composition, nil
}

// ReplaceComposition updates the complete composition
func (c *Client) ReplaceComposition(composition *Composition) error {
	endpoint := "/composition"
	return c.put(endpoint, composition, nil)
}

// CompositionAction executes undo or redo actions
func (c *Client) CompositionAction(action string) error {
	if action != "undo" && action != "redo" {
		return fmt.Errorf("invalid action: %s (must be 'undo' or 'redo')", action)
	}
	endpoint := "/composition/action"
	return c.post(endpoint, action, nil)
}

// DisconnectAllClips disconnects all clips in the composition
func (c *Client) DisconnectAllClips() error {
	endpoint := "/composition/disconnect-all"
	return c.post(endpoint, nil, nil)
}

// SetEffectDisplayName changes the display name of an effect
func (c *Client) SetEffectDisplayName(effectID int64, displayName string) error {
	endpoint := fmt.Sprintf("/composition/effects/by-id/%d/set-display-name", effectID)
	return c.post(endpoint, displayName, nil)
}

// MoveEffect moves an effect to the end of the composition
func (c *Client) MoveEffect(effectURI string) error {
	endpoint := "/composition/effects/video/move"
	return c.post(endpoint, effectURI, nil)
}

// MoveEffectToOffset moves an effect to a specific offset in the composition
func (c *Client) MoveEffectToOffset(offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/effects/video/move/%d", offset)
	return c.post(endpoint, effectURI, nil)
}

// AddEffect adds an effect to the entire composition
func (c *Client) AddEffect(effectURI string) error {
	endpoint := "/composition/effects/video/add"
	return c.post(endpoint, effectURI, nil)
}

// AddEffectAtOffset adds an effect to the composition at a specific offset
func (c *Client) AddEffectAtOffset(offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/effects/video/add/%d", offset)
	return c.post(endpoint, effectURI, nil)
}

// DeleteEffect removes an effect from the composition
func (c *Client) DeleteEffect(offset int64) error {
	endpoint := fmt.Sprintf("/composition/effects/video/%d", offset)
	return c.delete(endpoint)
}

// ResetCompositionParameter resets a parameter in the composition to its default value
func (c *Client) ResetCompositionParameter(parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// GetColumn retrieves column properties by index
func (c *Client) GetColumn(columnIndex int64) (*Column, error) {
	endpoint := fmt.Sprintf("/composition/columns/%d", columnIndex)
	var column Column
	if err := c.get(endpoint, &column); err != nil {
		return nil, err
	}
	return &column, nil
}

// ReplaceColumn updates a specific column by index
func (c *Client) ReplaceColumn(columnIndex int64, column *Column) error {
	endpoint := fmt.Sprintf("/composition/columns/%d", columnIndex)
	return c.put(endpoint, column, nil)
}

// DeleteColumn removes a column by index
func (c *Client) DeleteColumn(columnIndex int64) error {
	endpoint := fmt.Sprintf("/composition/columns/%d", columnIndex)
	return c.delete(endpoint)
}

// DuplicateColumn duplicates the given column
func (c *Client) DuplicateColumn(columnIndex int64) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/duplicate", columnIndex)
	return c.post(endpoint, nil, nil)
}

// AddColumn adds a new column to the composition
func (c *Client) AddColumn(beforeColumnURI string) error {
	endpoint := "/composition/columns/add"
	return c.post(endpoint, beforeColumnURI, nil)
}

// ResetColumnParameter resets a parameter in a column to its default value
func (c *Client) ResetColumnParameter(columnIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/%s/reset", columnIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// ConnectColumn connects the column by index
func (c *Client) ConnectColumn(columnIndex int64, connect *bool) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/connect", columnIndex)
	return c.post(endpoint, connect, nil)
}

// SelectColumn selects the column by index
func (c *Client) SelectColumn(columnIndex int64) error {
	endpoint := fmt.Sprintf("/composition/columns/%d/select", columnIndex)
	return c.post(endpoint, nil, nil)
}

// GetLayer retrieves layer properties and clip info by index
func (c *Client) GetLayer(layerIndex int64) (*Layer, error) {
	endpoint := fmt.Sprintf("/composition/layers/%d", layerIndex)
	var layer Layer
	if err := c.get(endpoint, &layer); err != nil {
		return nil, err
	}
	return &layer, nil
}

// ReplaceLayer updates specified layer and/or clips by index
func (c *Client) ReplaceLayer(layerIndex int64, layer *Layer) error {
	endpoint := fmt.Sprintf("/composition/layers/%d", layerIndex)
	return c.put(endpoint, layer, nil)
}

// DeleteLayer removes a layer by index
func (c *Client) DeleteLayer(layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d", layerIndex)
	return c.delete(endpoint)
}

// DuplicateLayer duplicates the given layer
func (c *Client) DuplicateLayer(layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/duplicate", layerIndex)
	return c.post(endpoint, nil, nil)
}

// AddLayer adds a new layer to the composition
func (c *Client) AddLayer(beforeLayerURI string) error {
	endpoint := "/composition/layers/add"
	return c.post(endpoint, beforeLayerURI, nil)
}

// ResetLayerParameter resets a parameter in a layer to its default value
func (c *Client) ResetLayerParameter(layerIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/%s/reset", layerIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// SelectLayer selects the layer by index
func (c *Client) SelectLayer(layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/select", layerIndex)
	return c.post(endpoint, nil, nil)
}

// ClearLayer disconnects any playing clips in the layer by index
func (c *Client) ClearLayer(layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clear", layerIndex)
	return c.post(endpoint, nil, nil)
}

// ClearLayerClips clears all clips in the layer by index
func (c *Client) ClearLayerClips(layerIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clearclips", layerIndex)
	return c.post(endpoint, nil, nil)
}

// GetSelectedLayer retrieves layer properties and clip info for the selected layers
func (c *Client) GetSelectedLayer() (*Layer, error) {
	endpoint := "/composition/layers/selected"
	var layer Layer
	if err := c.get(endpoint, &layer); err != nil {
		return nil, err
	}
	return &layer, nil
}

// ReplaceSelectedLayer updates selected layer and/or clips
func (c *Client) ReplaceSelectedLayer(layer *Layer) error {
	endpoint := "/composition/layers/selected"
	return c.put(endpoint, layer, nil)
}

// DuplicateSelectedLayer duplicates the selected layer
func (c *Client) DuplicateSelectedLayer() error {
	endpoint := "/composition/layers/selected/duplicate"
	return c.post(endpoint, nil, nil)
}

// AddEffectToSelectedLayer adds an effect to the selected layer
func (c *Client) AddEffectToSelectedLayer(effectURI string) error {
	endpoint := "/composition/layers/selected/effects/video/add"
	return c.post(endpoint, effectURI, nil)
}

// AddEffectToSelectedLayerAtOffset adds an effect at the given offset to the selected layer
func (c *Client) AddEffectToSelectedLayerAtOffset(offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/layers/selected/effects/video/add/%d", offset)
	return c.post(endpoint, effectURI, nil)
}

// DeleteSelectedLayerEffect removes an effect from the selected layer
func (c *Client) DeleteSelectedLayerEffect(offset int64) error {
	endpoint := fmt.Sprintf("/composition/layers/selected/effects/video/%d", offset)
	return c.delete(endpoint)
}

// ResetSelectedLayerParameter resets a parameter in the selected layer to its default value
func (c *Client) ResetSelectedLayerParameter(parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layers/selected/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// ClearSelectedLayer disconnects any playing clips in the selected layer
func (c *Client) ClearSelectedLayer() error {
	endpoint := "/composition/layers/selected/clear"
	return c.post(endpoint, nil, nil)
}

// ClearSelectedLayerClips clears all clips in the selected layer
func (c *Client) ClearSelectedLayerClips() error {
	endpoint := "/composition/layers/selected/clearclips"
	return c.post(endpoint, nil, nil)
}

// GetLayerGroup retrieves layer group properties and layer info by index
func (c *Client) GetLayerGroup(layerGroupIndex int64) (*LayerGroup, error) {
	endpoint := fmt.Sprintf("/composition/layergroups/%d", layerGroupIndex)
	var layerGroup LayerGroup
	if err := c.get(endpoint, &layerGroup); err != nil {
		return nil, err
	}
	return &layerGroup, nil
}

// ReplaceLayerGroup updates specified layer group and/or layers by index
func (c *Client) ReplaceLayerGroup(layerGroupIndex int64, layerGroup *LayerGroup) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d", layerGroupIndex)
	return c.put(endpoint, layerGroup, nil)
}

// DeleteLayerGroup removes a layer group by index
func (c *Client) DeleteLayerGroup(layerGroupIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d", layerGroupIndex)
	return c.delete(endpoint)
}

// DuplicateLayerGroup duplicates the given layer group
func (c *Client) DuplicateLayerGroup(layerGroupIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/duplicate", layerGroupIndex)
	return c.post(endpoint, nil, nil)
}

// MoveLayerToGroup adds an existing layer to an existing layer group
func (c *Client) MoveLayerToGroup(layerGroupIndex int64, layerURI string) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/move-layer", layerGroupIndex)
	return c.post(endpoint, layerURI, nil)
}

// AddLayerToGroup adds a new layer to an existing layer group
func (c *Client) AddLayerToGroup(layerGroupIndex int64, beforeLayerURI string) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/add-layer", layerGroupIndex)
	return c.post(endpoint, beforeLayerURI, nil)
}

// AddLayerGroup adds a new layer group to the composition
func (c *Client) AddLayerGroup(beforeLayerOrGroupURI string) error {
	endpoint := "/composition/layergroups/add"
	return c.post(endpoint, beforeLayerOrGroupURI, nil)
}

// ResetLayerGroupParameter resets a parameter in a layer group to its default value
func (c *Client) ResetLayerGroupParameter(layerGroupIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/%s/reset", layerGroupIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// SelectLayerGroup selects the layer group by index
func (c *Client) SelectLayerGroup(layerGroupIndex int64) error {
	endpoint := fmt.Sprintf("/composition/layergroups/%d/select", layerGroupIndex)
	return c.post(endpoint, nil, nil)
}

// GetSelectedLayerGroup retrieves selected layer group properties and layer info
func (c *Client) GetSelectedLayerGroup() (*LayerGroup, error) {
	endpoint := "/composition/layergroups/selected"
	var layerGroup LayerGroup
	if err := c.get(endpoint, &layerGroup); err != nil {
		return nil, err
	}
	return &layerGroup, nil
}

// ReplaceSelectedLayerGroup updates selected layer group and/or layers
func (c *Client) ReplaceSelectedLayerGroup(layerGroup *LayerGroup) error {
	endpoint := "/composition/layergroups/selected"
	return c.put(endpoint, layerGroup, nil)
}

// DeleteSelectedLayerGroup removes the selected layer group
func (c *Client) DeleteSelectedLayerGroup() error {
	endpoint := "/composition/layergroups/selected"
	return c.delete(endpoint)
}

// DuplicateSelectedLayerGroup duplicates the selected layer group
func (c *Client) DuplicateSelectedLayerGroup() error {
	endpoint := "/composition/layergroups/selected/duplicate"
	return c.post(endpoint, nil, nil)
}

// MoveLayerToSelectedGroup adds an existing layer to the selected layer group
func (c *Client) MoveLayerToSelectedGroup(layerURI string) error {
	endpoint := "/composition/layergroups/selected/move-layer"
	return c.post(endpoint, layerURI, nil)
}

// AddLayerToSelectedGroup adds a new layer to the selected layer group
func (c *Client) AddLayerToSelectedGroup(beforeLayerURI string) error {
	endpoint := "/composition/layergroups/selected/add-layer"
	return c.post(endpoint, beforeLayerURI, nil)
}

// ResetSelectedLayerGroupParameter resets a parameter in the selected layer group to its default value
func (c *Client) ResetSelectedLayerGroupParameter(parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/layergroups/selected/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// GetDeck retrieves deck properties by index
func (c *Client) GetDeck(deckIndex int64) (*Deck, error) {
	endpoint := fmt.Sprintf("/composition/decks/%d", deckIndex)
	var deck Deck
	if err := c.get(endpoint, &deck); err != nil {
		return nil, err
	}
	return &deck, nil
}

// ReplaceDeck updates a specific deck by index
func (c *Client) ReplaceDeck(deckIndex int64, deck *Deck) error {
	endpoint := fmt.Sprintf("/composition/decks/%d", deckIndex)
	return c.put(endpoint, deck, nil)
}

// DeleteDeck removes a deck by index
func (c *Client) DeleteDeck(deckIndex int64) error {
	endpoint := fmt.Sprintf("/composition/decks/%d", deckIndex)
	return c.delete(endpoint)
}

// DuplicateDeck duplicates the given deck
func (c *Client) DuplicateDeck(deckIndex int64) error {
	endpoint := fmt.Sprintf("/composition/decks/%d/duplicate", deckIndex)
	return c.post(endpoint, nil, nil)
}

// AddDeck adds a new deck to the composition
func (c *Client) AddDeck(beforeDeckURI string) error {
	endpoint := "/composition/decks/add"
	return c.post(endpoint, beforeDeckURI, nil)
}

// ResetDeckParameter resets a parameter in a deck to its default value
func (c *Client) ResetDeckParameter(deckIndex int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/decks/%d/%s/reset", deckIndex, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// SelectDeck selects the deck by index
func (c *Client) SelectDeck(deckIndex int64) error {
	endpoint := fmt.Sprintf("/composition/decks/%d/select", deckIndex)
	return c.post(endpoint, nil, nil)
}

// GetDeckByID retrieves deck properties by id
func (c *Client) GetDeckByID(deckID int64) (*Deck, error) {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d", deckID)
	var deck Deck
	if err := c.get(endpoint, &deck); err != nil {
		return nil, err
	}
	return &deck, nil
}

// ReplaceDeckByID updates specific deck by id
func (c *Client) ReplaceDeckByID(deckID int64, deck *Deck) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d", deckID)
	return c.put(endpoint, deck, nil)
}

// DeleteDeckByID removes specified deck by id
func (c *Client) DeleteDeckByID(deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d", deckID)
	return c.delete(endpoint)
}

// DuplicateDeckByID duplicates the given deck
func (c *Client) DuplicateDeckByID(deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/duplicate", deckID)
	return c.post(endpoint, nil, nil)
}

// CloseDeckByID closes the given deck
func (c *Client) CloseDeckByID(deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/close", deckID)
	return c.post(endpoint, nil, nil)
}

// OpenDeckByID re-opens the given deck
func (c *Client) OpenDeckByID(deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/open", deckID)
	return c.post(endpoint, nil, nil)
}

// ResetDeckParameterByID resets a parameter in a deck to its default value
func (c *Client) ResetDeckParameterByID(deckID int64, parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/%s/reset", deckID, parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// SelectDeckByID selects the deck by id
func (c *Client) SelectDeckByID(deckID int64) error {
	endpoint := fmt.Sprintf("/composition/decks/by-id/%d/select", deckID)
	return c.post(endpoint, nil, nil)
}

// GetClipByPosition retrieves a clip by its position in the clip grid
func (c *Client) GetClipByPosition(layerIndex, clipIndex int64) (*Clip, error) {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d", layerIndex, clipIndex)
	var clip Clip
	if err := c.get(endpoint, &clip); err != nil {
		return nil, err
	}
	return &clip, nil
}

// ReplaceClipByPosition updates clip and/or its effects by position in the clip grid
func (c *Client) ReplaceClipByPosition(layerIndex, clipIndex int64, clip *Clip) error {
	endpoint := fmt.Sprintf("/composition/layers/%d/clips/%d", layerIndex, clipIndex)
	return c.put(endpoint, clip, nil)
}

// GetSelectedClip retrieves the selected clip
func (c *Client) GetSelectedClip() (*Clip, error) {
	endpoint := "/composition/clips/selected"
	var clip Clip
	if err := c.get(endpoint, &clip); err != nil {
		return nil, err
	}
	return &clip, nil
}

// ReplaceSelectedClip updates selected clip and/or its effects
func (c *Client) ReplaceSelectedClip(clip *Clip) error {
	endpoint := "/composition/clips/selected"
	return c.put(endpoint, clip, nil)
}

// AddEffectToSelectedClip adds an effect to the selected clip
func (c *Client) AddEffectToSelectedClip(effectURI string) error {
	endpoint := "/composition/clips/selected/effects/video/add"
	return c.post(endpoint, effectURI, nil)
}

// AddEffectToSelectedClipAtOffset adds an effect at the given offset to the selected clip
func (c *Client) AddEffectToSelectedClipAtOffset(offset int64, effectURI string) error {
	endpoint := fmt.Sprintf("/composition/clips/selected/effects/video/add/%d", offset)
	return c.post(endpoint, effectURI, nil)
}

// DeleteSelectedClipEffect removes an effect from the selected clip
func (c *Client) DeleteSelectedClipEffect(offset int64) error {
	endpoint := fmt.Sprintf("/composition/clips/selected/effects/video/%d", offset)
	return c.delete(endpoint)
}

// ResetSelectedClipParameter resets a parameter in the selected clip to its default value
func (c *Client) ResetSelectedClipParameter(parameter string, resetAnimation bool) error {
	endpoint := fmt.Sprintf("/composition/clips/selected/%s/reset", parameter)
	body := ResetParameter{
		ResetAnimation: resetAnimation,
	}
	return c.post(endpoint, body, nil)
}

// ConnectSelectedClip connects the selected clip
func (c *Client) ConnectSelectedClip(connect *bool) error {
	endpoint := "/composition/clips/selected/connect"
	return c.post(endpoint, connect, nil)
}

// OpenSelectedClip loads a file or opens a source into the selected clip
func (c *Client) OpenSelectedClip(uri string) error {
	endpoint := "/composition/clips/selected/open"
	return c.post(endpoint, uri, nil)
}

// ClearSelectedClip clears the selected clip
func (c *Client) ClearSelectedClip() error {
	endpoint := "/composition/clips/selected/clear"
	return c.post(endpoint, nil, nil)
}

// GetClipByID retrieves a clip by id
func (c *Client) GetClipByID(clipID int64) (*Clip, error) {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d", clipID)
	var clip Clip
	if err := c.get(endpoint, &clip); err != nil {
		return nil, err
	}
	return &clip, nil
}

// ReplaceClipByID updates clip and/or its effects by id
func (c *Client) ReplaceClipByID(clipID int64, clip *Clip) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d", clipID)
	return c.put(endpoint, clip, nil)
}

// SelectClipByID selects the clip by id
func (c *Client) SelectClipByID(clipID int64) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/select", clipID)
	return c.post(endpoint, nil, nil)
}

// ConnectClipByID connects the clip by id
func (c *Client) ConnectClipByID(clipID int64, connect *bool) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/connect", clipID)
	return c.post(endpoint, connect, nil)
}

// OpenClipByID loads a file or opens a source into the clip with the given id
func (c *Client) OpenClipByID(clipID int64, uri string) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/open", clipID)
	return c.post(endpoint, uri, nil)
}

// ClearClipByID clears the clip with the given id
func (c *Client) ClearClipByID(clipID int64) error {
	endpoint := fmt.Sprintf("/composition/clips/by-id/%d/clear", clipID)
	return c.post(endpoint, nil, nil)
}
