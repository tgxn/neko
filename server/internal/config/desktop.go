package config

import (
	"os"
	"regexp"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/m1k1o/neko/server/pkg/types"
)

type Desktop struct {
	Display string

	ScreenSize types.ScreenSize

	UseInputDriver bool
	InputSocket    string

	Unminimize        bool
	UploadDrop        bool
	FileChooserDialog bool
}

func (Desktop) Init(cmd *cobra.Command) error {
	cmd.PersistentFlags().String("desktop.display", "", "X display to use for desktop sharing")
	if err := viper.BindPFlag("desktop.display", cmd.PersistentFlags().Lookup("desktop.display")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("desktop.screen", "1280x720@30", "default screen size and framerate")
	if err := viper.BindPFlag("desktop.screen", cmd.PersistentFlags().Lookup("desktop.screen")); err != nil {
		return err
	}

	cmd.PersistentFlags().Bool("desktop.input.enabled", true, "whether custom xf86 input driver should be used to handle touchscreen")
	if err := viper.BindPFlag("desktop.input.enabled", cmd.PersistentFlags().Lookup("desktop.input.enabled")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("desktop.input.socket", "/tmp/xf86-input-neko.sock", "socket path for custom xf86 input driver connection")
	if err := viper.BindPFlag("desktop.input.socket", cmd.PersistentFlags().Lookup("desktop.input.socket")); err != nil {
		return err
	}

	cmd.PersistentFlags().Bool("desktop.unminimize", true, "automatically unminimize window when it is minimized")
	if err := viper.BindPFlag("desktop.unminimize", cmd.PersistentFlags().Lookup("desktop.unminimize")); err != nil {
		return err
	}

	cmd.PersistentFlags().Bool("desktop.upload_drop", true, "whether drop upload is enabled")
	if err := viper.BindPFlag("desktop.upload_drop", cmd.PersistentFlags().Lookup("desktop.upload_drop")); err != nil {
		return err
	}

	cmd.PersistentFlags().Bool("desktop.file_chooser_dialog", false, "whether to handle file chooser dialog externally")
	if err := viper.BindPFlag("desktop.file_chooser_dialog", cmd.PersistentFlags().Lookup("desktop.file_chooser_dialog")); err != nil {
		return err
	}

	return nil
}

func (Desktop) InitV2(cmd *cobra.Command) error {
	cmd.PersistentFlags().String("screen", "", "V2: default screen resolution and framerate")
	if err := viper.BindPFlag("screen", cmd.PersistentFlags().Lookup("screen")); err != nil {
		return err
	}

	return nil
}

func (s *Desktop) Set() {
	s.Display = viper.GetString("desktop.display")

	// Display is provided by env variable unless explicitly set
	if s.Display == "" {
		s.Display = os.Getenv("DISPLAY")
	}

	s.ScreenSize = types.ScreenSize{
		Width:  1280,
		Height: 720,
		Rate:   30,
	}

	r := regexp.MustCompile(`([0-9]{1,4})x([0-9]{1,4})@([0-9]{1,3})`)
	res := r.FindStringSubmatch(viper.GetString("desktop.screen"))

	if len(res) > 0 {
		width, err1 := strconv.ParseInt(res[1], 10, 64)
		height, err2 := strconv.ParseInt(res[2], 10, 64)
		rate, err3 := strconv.ParseInt(res[3], 10, 64)

		if err1 == nil && err2 == nil && err3 == nil {
			s.ScreenSize.Width = int(width)
			s.ScreenSize.Height = int(height)
			s.ScreenSize.Rate = int16(rate)
		}
	}

	s.UseInputDriver = viper.GetBool("desktop.input.enabled")
	s.InputSocket = viper.GetString("desktop.input.socket")
	s.Unminimize = viper.GetBool("desktop.unminimize")
	s.UploadDrop = viper.GetBool("desktop.upload_drop")
	s.FileChooserDialog = viper.GetBool("desktop.file_chooser_dialog")
}

func (s *Desktop) SetV2() {
	enableLegacy := false

	if viper.IsSet("screen") {
		r := regexp.MustCompile(`([0-9]{1,4})x([0-9]{1,4})@([0-9]{1,3})`)
		res := r.FindStringSubmatch(viper.GetString("screen"))

		if len(res) > 0 {
			width, err1 := strconv.ParseInt(res[1], 10, 64)
			height, err2 := strconv.ParseInt(res[2], 10, 64)
			rate, err3 := strconv.ParseInt(res[3], 10, 64)

			if err1 == nil && err2 == nil && err3 == nil {
				s.ScreenSize.Width = int(width)
				s.ScreenSize.Height = int(height)
				s.ScreenSize.Rate = int16(rate)
			}
		}
		log.Warn().Msg("you are using v2 configuration 'NEKO_SCREEN' which is deprecated, please use 'NEKO_DESKTOP_SCREEN' instead")
		enableLegacy = true
	}

	// set legacy flag if any V2 configuration was used
	if !viper.IsSet("legacy") && enableLegacy {
		log.Warn().Msg("legacy configuration is enabled because at least one V2 configuration was used, please migrate to V3 configuration, or set 'NEKO_LEGACY=true' to acknowledge this message")
		viper.Set("legacy", true)
	}
}
