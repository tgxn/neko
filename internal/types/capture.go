package types

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/PaesslerAG/gval"
	"github.com/pion/webrtc/v3/pkg/media"

	"demodesk/neko/internal/types/codec"
)

var (
	ErrCapturePipelineAlreadyExists = errors.New("capture pipeline already exists")
)

type Sample media.Sample

type BroadcastManager interface {
	Start(url string) error
	Stop()
	Started() bool
	Url() string
}

type ScreencastManager interface {
	Enabled() bool
	Started() bool
	Image() ([]byte, error)
}

type StreamManager interface {
	Codec() codec.RTPCodec

	// starts pipeline if was not running before and returns register function
	NewListener(listener *func(sample Sample)) (addListener func(), err error)
	// stops pipeline if it was last listener
	RemoveListener(listener *func(sample Sample))

	ListenersCount() int
	Started() bool
}

type CaptureManager interface {
	Start()
	Shutdown() error

	Broadcast() BroadcastManager
	Screencast() ScreencastManager
	Audio() StreamManager
	Video(videoID string) (StreamManager, bool)
	VideoIDs() []string
}

type VideoConfig struct {
	Width       string            `mapstructure:"width"`        // expression
	Height      string            `mapstructure:"height"`       // expression
	Fps         string            `mapstructure:"fps"`          // expression
	GstPrefix   string            `mapstructure:"gst_prefix"`   // pipeline prefix, starts with !
	GstEncoder  string            `mapstructure:"gst_encoder"`  // gst encoder name
	GstParams   map[string]string `mapstructure:"gst_params"`   // map of expressions
	GstSuffix   string            `mapstructure:"gst_suffix"`   // pipeline suffix, starts with !
	GstPipeline string            `mapstructure:"gst_pipeline"` // whole pipeline as a string
}

func (config *VideoConfig) GetPipeline(screen ScreenSize) (string, error) {
	values := map[string]interface{}{
		"width":  screen.Width,
		"height": screen.Height,
		"fps":    screen.Rate,
	}

	language := []gval.Language{
		gval.Function("round", func(args ...interface{}) (interface{}, error) {
			return (int)(math.Round(args[0].(float64))), nil
		}),
	}

	// get fps pipeline
	fpsPipeline := "! video/x-raw ! videoconvert ! queue"
	if config.Fps != "" {
		eval, err := gval.Full(language...).NewEvaluable(config.Fps)
		if err != nil {
			return "", err
		}

		val, err := eval.EvalFloat64(context.Background(), values)
		if err != nil {
			return "", err
		}

		fpsPipeline = fmt.Sprintf("! video/x-raw,framerate=%d/100 ! videoconvert ! queue", int(val*100))
	}

	// get scale pipeline
	scalePipeline := ""
	if config.Width != "" && config.Height != "" {
		eval, err := gval.Full(language...).NewEvaluable(config.Width)
		if err != nil {
			return "", err
		}

		w, err := eval.EvalInt(context.Background(), values)
		if err != nil {
			return "", err
		}

		eval, err = gval.Full(language...).NewEvaluable(config.Height)
		if err != nil {
			return "", err
		}

		h, err := eval.EvalInt(context.Background(), values)
		if err != nil {
			return "", err
		}

		scalePipeline = fmt.Sprintf("! videoscale ! video/x-raw,width=%d,height=%d ! queue", w, h)
	}

	// get encoder pipeline
	encPipeline := fmt.Sprintf("! %s", config.GstEncoder)
	for key, expr := range config.GstParams {
		if expr == "" {
			continue
		}

		val, err := gval.Evaluate(expr, values, language...)
		if err != nil {
			return "", err
		}

		if val != nil {
			encPipeline += fmt.Sprintf(" %s=%v", key, val)
		} else {
			encPipeline += fmt.Sprintf(" %s=%s", key, expr)
		}
	}

	// join strings with space
	return strings.Join([]string{
		fpsPipeline,
		scalePipeline,
		config.GstPrefix,
		encPipeline,
		config.GstSuffix,
	}[:], " "), nil
}
