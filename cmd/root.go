/*
Copyright © 2026 kotob-project contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	toLang   string
	fromLang string
	apiKey   string
	model    string
	system   string
	asJson   bool
)

var rootCmd = &cobra.Command{
	Use:   "kotob [flags] [text]",
	Short: "A lightweight CLI translation tool powered by Gemini API",
	Long: `Kotob is a lightweight CLI translation tool built with Go,
leveraging the Google Gemini API for fast and accurate translations.`,

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: 翻訳の呼び出し
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVarP(&toLang, "to", "t", "", "target language (defaults to en ⇔ ja if unspecified)")
	rootCmd.Flags().StringVarP(&fromLang, "from", "f", "auto", "source language (default auto)")
	rootCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "Gemini API key for the session")
	rootCmd.Flags().StringVarP(&model, "model", "m", "gemini-2.0-flash-lite", "AI model to use")
	rootCmd.Flags().StringVarP(&system, "system", "s", "", "custom system prompt for the AI")
	rootCmd.Flags().BoolVarP(&asJson, "json", "j", false, "output result as a JSON object")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetEnvPrefix("KOTOB")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// 設定ファイル関連
	viper.SetConfigName("kotob")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if home, err := os.UserHomeDir(); err == nil {
		viper.AddConfigPath(home + "/.config/kotob")
	}

	_ = viper.ReadInConfig()

	// 補完処理
	if apiKey == "" {
		apiKey = viper.GetString("api-key")
	}
	if model == "" || model == "gemini-2.0-flash-lite" {
		vModel := viper.GetString("model")
		if vModel != "" {
			model = vModel
		}
	}
	if toLang == "" {
		toLang = viper.GetString("to")
	}
	if fromLang == "" || fromLang == "auto" {
		vfromLang := viper.GetString("from")
		if vfromLang != "" {
			fromLang = vfromLang
		}
	}
	if system == "" {
		system = viper.GetString("system")
	}
	if !asJson {
		asJson = viper.GetBool("json")
	}
}
