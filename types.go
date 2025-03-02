package resolume

// ProductInfo represents information about the Resolume product
type ProductInfo struct {
	Name     string `json:"name"`
	Major    int64  `json:"major"`
	Minor    int64  `json:"minor"`
	Micro    int64  `json:"micro"`
	Revision int64  `json:"revision"`
}

// Effect represents an effect to be used on clips/layers/composition
type Effect struct {
	IDString string `json:"idstring"`
	Name     string `json:"name"`
	Presets  []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"presets,omitempty"`
}

// Effects represents available effects for clips/layers/composition
type Effects struct {
	Video []Effect `json:"video"`
}

// Source represents a source that can be used in clips
type Source struct {
	IDString string `json:"idstring"`
	Name     string `json:"name"`
	Presets  []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"presets,omitempty"`
}

// Sources represents available sources for clips
type Sources struct {
	Video []Source `json:"video"`
}

// ParameterView represents semantic information about how to display a parameter
type ParameterView struct {
	Suffix       string  `json:"suffix,omitempty"`
	Step         float64 `json:"step,omitempty"`
	Multiplier   float64 `json:"multiplier,omitempty"`
	DisplayUnits string  `json:"display_units,omitempty"`
	ControlType  string  `json:"control_type,omitempty"`
}

// BooleanParameter represents a parameter containing a true or false value
type BooleanParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	Value     bool          `json:"value"`
	View      ParameterView `json:"view,omitempty"`
}

// ChoiceParameter represents a multiple-choice parameter
type ChoiceParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	Value     string        `json:"value"`
	Index     int32         `json:"index"`
	Options   []string      `json:"options"`
	View      ParameterView `json:"view,omitempty"`
}

// ColorParameter represents a parameter containing color data
type ColorParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	Value     string        `json:"value"`
	Palette   []string      `json:"palette,omitempty"`
	View      ParameterView `json:"view,omitempty"`
}

// EventParameter represents a parameter that handles events but does not contain a value
type EventParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	View      ParameterView `json:"view,omitempty"`
}

// IntegerParameter represents a parameter containing numeric data
type IntegerParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	Value     int64         `json:"value"`
	View      ParameterView `json:"view,omitempty"`
}

// RangeParameter represents a parameter containing a floating-point value with min/max
type RangeParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	Min       float64       `json:"min"`
	Max       float64       `json:"max"`
	In        float64       `json:"in"`
	Out       float64       `json:"out"`
	Value     float64       `json:"value"`
	View      ParameterView `json:"view,omitempty"`
}

// StringParameter represents a parameter containing string data
type StringParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	Value     string        `json:"value"`
	View      ParameterView `json:"view,omitempty"`
}

// TextParameter represents a parameter containing possibly multiline string data
type TextParameter struct {
	ID        int64         `json:"id"`
	ValueType string        `json:"valuetype"`
	Value     string        `json:"value"`
	View      ParameterView `json:"view,omitempty"`
}

// ResetParameter represents options for resetting a parameter
type ResetParameter struct {
	ResetAnimation bool `json:"resetanimation"`
}

// ParameterCollection represents an unstructured collection of parameters
type ParameterCollection map[string]interface{}

// AudioEffect represents a single audio effect in a chain
type AudioEffect struct {
	ID       int64               `json:"id"`
	Name     string              `json:"name"`
	Bypassed *BooleanParameter   `json:"bypassed,omitempty"`
	Params   ParameterCollection `json:"params,omitempty"`
}

// VideoEffect represents a single video effect in a chain
type VideoEffect struct {
	ID          int64               `json:"id"`
	Name        string              `json:"name"`
	DisplayName string              `json:"display_name"`
	Bypassed    *BooleanParameter   `json:"bypassed,omitempty"`
	Mixer       ParameterCollection `json:"mixer,omitempty"`
	Params      ParameterCollection `json:"params,omitempty"`
	Effect      ParameterCollection `json:"effect,omitempty"`
}

// AudioFileInfo represents meta information for an audio file
type AudioFileInfo struct {
	Path        string  `json:"path"`
	Exists      bool    `json:"exists"`
	Duration    string  `json:"duration"`
	DurationMs  float64 `json:"duration_ms"`
	SampleRate  float64 `json:"sample_rate"`
	NumChannels int32   `json:"num_channels"`
	BPM         float64 `json:"bpm"`
}

// AudioTrack represents an audio track as part of a clip/layer/group/composition
type AudioTrack struct {
	Volume  *RangeParameter `json:"volume,omitempty"`
	Pan     *RangeParameter `json:"pan,omitempty"`
	Effects []AudioEffect   `json:"effects,omitempty"`
}

// AudioTrackClip represents an audio track specifically for clips
type AudioTrackClip struct {
	AudioTrack
	Description string         `json:"description,omitempty"`
	FileInfo    *AudioFileInfo `json:"fileinfo,omitempty"`
}

// AutoPilot represents options to control automatic clip transitions
type AutoPilot struct {
	Target *ChoiceParameter `json:"target,omitempty"`
}

// TransportBPMSync represents BPM sync transport controls
type TransportBPMSync struct {
	Position *RangeParameter `json:"position,omitempty"`
	Controls struct {
		PlayDirection *ChoiceParameter `json:"playdirection,omitempty"`
		PlayMode      *ChoiceParameter `json:"playmode,omitempty"`
		PlayModeAway  *ChoiceParameter `json:"playmodeaway,omitempty"`
		Duration      *RangeParameter  `json:"duration,omitempty"`
		Speed         *RangeParameter  `json:"speed,omitempty"`
		BPM           *RangeParameter  `json:"bpm,omitempty"`
		SyncMode      *ChoiceParameter `json:"syncmode,omitempty"`
		BeatLoop      *ChoiceParameter `json:"beatloop,omitempty"`
	} `json:"controls"`
}

// TransportTimeline represents timeline transport controls
type TransportTimeline struct {
	Position *RangeParameter `json:"position,omitempty"`
	Controls struct {
		PlayDirection *ChoiceParameter `json:"playdirection,omitempty"`
		PlayMode      *ChoiceParameter `json:"playmode,omitempty"`
		PlayModeAway  *ChoiceParameter `json:"playmodeaway,omitempty"`
		Duration      *RangeParameter  `json:"duration,omitempty"`
		Speed         *RangeParameter  `json:"speed,omitempty"`
	} `json:"controls"`
}

// FrameRate represents frame rate expressed as a ratio
type FrameRate struct {
	Num   int32 `json:"num"`
	Denom int32 `json:"denom"`
}

// VideoFileInfo represents meta information for a video file
type VideoFileInfo struct {
	Path       string    `json:"path"`
	Exists     bool      `json:"exists"`
	Duration   string    `json:"duration"`
	DurationMs float64   `json:"duration_ms"`
	FrameRate  FrameRate `json:"framerate"`
	Width      int32     `json:"width"`
	Height     int32     `json:"height"`
}

// VideoTrack represents a video track as part of a clip/layer/group/composition
type VideoTrack struct {
	Width   *RangeParameter     `json:"width,omitempty"`
	Height  *RangeParameter     `json:"height,omitempty"`
	Opacity *RangeParameter     `json:"opacity,omitempty"`
	Mixer   ParameterCollection `json:"mixer,omitempty"`
	Effects []VideoEffect       `json:"effects,omitempty"`
}

// VideoTrackLayer represents a video track specifically for layers
type VideoTrackLayer struct {
	VideoTrack
	AutoSize *ChoiceParameter `json:"autosize,omitempty"`
}

// VideoTrackClip represents a video track specifically for clips
type VideoTrackClip struct {
	VideoTrack
	Description  string              `json:"description,omitempty"`
	FileInfo     *VideoFileInfo      `json:"fileinfo,omitempty"`
	Resize       *ChoiceParameter    `json:"resize,omitempty"`
	R            *BooleanParameter   `json:"r,omitempty"`
	G            *BooleanParameter   `json:"g,omitempty"`
	B            *BooleanParameter   `json:"b,omitempty"`
	A            *BooleanParameter   `json:"a,omitempty"`
	SourceParams ParameterCollection `json:"sourceparams,omitempty"`
}

// LayerTransition describes the transition between clips within a layer
type LayerTransition struct {
	Duration  *RangeParameter  `json:"duration,omitempty"`
	BlendMode *ChoiceParameter `json:"blend_mode,omitempty"`
}

// Clip represents a single clip in the composition
type Clip struct {
	ID                  int64               `json:"id"`
	Name                *StringParameter    `json:"name,omitempty"`
	ColorID             *ChoiceParameter    `json:"colorid,omitempty"`
	Selected            *BooleanParameter   `json:"selected,omitempty"`
	Connected           *ChoiceParameter    `json:"connected,omitempty"`
	Target              *ChoiceParameter    `json:"target,omitempty"`
	TriggerStyle        *ChoiceParameter    `json:"triggerstyle,omitempty"`
	IgnoreColumnTrigger *ChoiceParameter    `json:"ignorecolumntrigger,omitempty"`
	FaderStart          *ChoiceParameter    `json:"faderstart,omitempty"`
	BeatSnap            *ChoiceParameter    `json:"beatsnap,omitempty"`
	TransportType       *ChoiceParameter    `json:"transporttype,omitempty"`
	Transport           interface{}         `json:"transport,omitempty"` // Can be TransportTimeline or TransportBPMSync
	Dashboard           ParameterCollection `json:"dashboard,omitempty"`
	Audio               *AudioTrackClip     `json:"audio,omitempty"`
	Video               *VideoTrackClip     `json:"video,omitempty"`
	Thumbnail           *struct {
		Size       int64  `json:"size"`
		LastUpdate string `json:"last_update"`
		IsDefault  bool   `json:"is_default"`
	} `json:"thumbnail,omitempty"`
}

// Column represents a column within a deck
type Column struct {
	ID        int64             `json:"id"`
	Name      *StringParameter  `json:"name,omitempty"`
	ColorID   *ChoiceParameter  `json:"colorid,omitempty"`
	Connected *ChoiceParameter  `json:"connected,omitempty"`
	Selected  *BooleanParameter `json:"selected,omitempty"`
}

// CrossFader represents cross fade between two clips
type CrossFader struct {
	ID        int64               `json:"id"`
	Phase     *RangeParameter     `json:"phase,omitempty"`
	Behaviour *ChoiceParameter    `json:"behaviour,omitempty"`
	Curve     *ChoiceParameter    `json:"curve,omitempty"`
	SideA     *EventParameter     `json:"sidea,omitempty"`
	SideB     *EventParameter     `json:"sideb,omitempty"`
	Mixer     ParameterCollection `json:"mixer,omitempty"`
}

// Deck represents a deck containing layers and clips
type Deck struct {
	ID       int64             `json:"id"`
	Closed   bool              `json:"closed"`
	Name     *StringParameter  `json:"name,omitempty"`
	ColorID  *ChoiceParameter  `json:"colorid,omitempty"`
	Selected *BooleanParameter `json:"selected,omitempty"`
	ScrollX  *IntegerParameter `json:"scrollx,omitempty"`
}

// Layer represents a container for clips
type Layer struct {
	ID                  int64               `json:"id"`
	Name                *StringParameter    `json:"name,omitempty"`
	ColorID             *ChoiceParameter    `json:"colorid,omitempty"`
	Selected            *BooleanParameter   `json:"selected,omitempty"`
	Bypassed            *BooleanParameter   `json:"bypassed,omitempty"`
	Solo                *BooleanParameter   `json:"solo,omitempty"`
	CrossFaderGroup     *ChoiceParameter    `json:"crossfadergroup,omitempty"`
	Master              *RangeParameter     `json:"master,omitempty"`
	MaskMode            *ChoiceParameter    `json:"maskmode,omitempty"`
	IgnoreColumnTrigger *BooleanParameter   `json:"ignorecolumntrigger,omitempty"`
	FaderStart          *BooleanParameter   `json:"faderstart,omitempty"`
	Dashboard           ParameterCollection `json:"dashboard,omitempty"`
	Audio               *AudioTrack         `json:"audio,omitempty"`
	Video               *VideoTrackLayer    `json:"video,omitempty"`
	Transition          *LayerTransition    `json:"transition,omitempty"`
	Clips               []Clip              `json:"clips,omitempty"`
	AutoPilot           *AutoPilot          `json:"autopilot,omitempty"`
}

// LayerGroup represents a collection of layers
type LayerGroup struct {
	ID                  int64               `json:"id"`
	Name                *StringParameter    `json:"name,omitempty"`
	ColorID             *ChoiceParameter    `json:"colorid,omitempty"`
	Selected            *BooleanParameter   `json:"selected,omitempty"`
	Bypassed            *BooleanParameter   `json:"bypassed,omitempty"`
	Solo                *BooleanParameter   `json:"solo,omitempty"`
	CrossFaderGroup     *ChoiceParameter    `json:"crossfadergroup,omitempty"`
	Master              *RangeParameter     `json:"master,omitempty"`
	Speed               *RangeParameter     `json:"speed,omitempty"`
	IgnoreColumnTrigger *BooleanParameter   `json:"ignorecolumntrigger,omitempty"`
	Dashboard           ParameterCollection `json:"dashboard,omitempty"`
	Audio               *AudioTrack         `json:"audio,omitempty"`
	Video               *VideoTrack         `json:"video,omitempty"`
	Layers              []Layer             `json:"layers,omitempty"`
}

// TempoController represents the controller for tempo-related aspects
type TempoController struct {
	Tempo     *RangeParameter `json:"tempo,omitempty"`
	TempoPull *EventParameter `json:"tempo_pull,omitempty"`
	TempoPush *EventParameter `json:"tempo_push,omitempty"`
	TempoTap  *EventParameter `json:"tempo_tap,omitempty"`
	Resync    *EventParameter `json:"resync,omitempty"`
}

// Composition represents the complete composition
type Composition struct {
	Name             *StringParameter    `json:"name,omitempty"`
	Selected         *BooleanParameter   `json:"selected,omitempty"`
	Bypassed         *BooleanParameter   `json:"bypassed,omitempty"`
	Master           *RangeParameter     `json:"master,omitempty"`
	Speed            *RangeParameter     `json:"speed,omitempty"`
	ClipTarget       *ChoiceParameter    `json:"cliptarget,omitempty"`
	ClipTriggerStyle *ChoiceParameter    `json:"cliptriggerstyle,omitempty"`
	ClipBeatSnap     *ChoiceParameter    `json:"clipbeatsnap,omitempty"`
	Dashboard        ParameterCollection `json:"dashboard,omitempty"`
	Audio            *AudioTrack         `json:"audio,omitempty"`
	Video            *VideoTrack         `json:"video,omitempty"`
	CrossFader       *CrossFader         `json:"crossfader,omitempty"`
	Decks            []Deck              `json:"decks,omitempty"`
	Layers           []Layer             `json:"layers,omitempty"`
	Columns          []Column            `json:"columns,omitempty"`
	LayerGroups      []LayerGroup        `json:"layergroups,omitempty"`
	TempoController  *TempoController    `json:"tempo_controller,omitempty"`
}
