package mongo

import (
	"github.com/czyt/kilot/cmd/kilot/kratos/internal/templateContext"
	"github.com/urfave/cli/v2"
)

const (
	userOdmFlag      = "use-odm"
	enableSoftDelete = "with-soft-delete"
	formatFlag       = "format"
	modelSuffixFlag  = "model-suffix"
	modelPrefixFlag  = "model-prefix"
	bizLayerCodeDir  = "layer-biz-dir"
	dataLayerCodeDir = "layer-data-dir"
	modelNamesFlag   = "model-names"
	useCaseFlag      = "use-case-name"
	modelOutputFlag  = "model-output-dir"
)

var (
	tplContext templateContext.MongoContext
)

func Options() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        userOdmFlag,
			Value:       true,
			Usage:       "this flag set whether to use mongo odm (use mgm).",
			Destination: &tplContext.UseOdm,
		},
		&cli.BoolFlag{
			Name:        enableSoftDelete,
			Aliases:     []string{"sd"},
			Value:       false,
			Usage:       "this flag set whether to generate soft delete feature code.",
			Destination: &tplContext.WithSoftDelete,
		},
		&cli.BoolFlag{
			Name:        formatFlag,
			Aliases:     []string{"f"},
			Value:       true,
			Usage:       "set whether to format generated code before write to file.",
			Destination: &tplContext.FormatCode,
		},
		&cli.StringFlag{
			Name:        modelSuffixFlag,
			Aliases:     []string{"s"},
			Value:       "",
			Usage:       "set model name suffix `Suffix` .default is empty.",
			Destination: &tplContext.ModelSuffix,
		},
		&cli.StringFlag{
			Name:        modelPrefixFlag,
			Aliases:     []string{"p"},
			Value:       "",
			Usage:       "set model name prefix `Prefix`.default is empty.",
			Destination: &tplContext.ModelPrefix,
		},
		&cli.StringFlag{
			Name:        bizLayerCodeDir,
			Aliases:     []string{"bd"},
			Value:       "biz",
			Usage:       "set biz layer code store Dir `DIR`.default is `biz`.",
			Destination: &tplContext.BizLayerCodeDir,
		},
		&cli.StringFlag{
			Name:        dataLayerCodeDir,
			Aliases:     []string{"dd"},
			Value:       "data",
			Usage:       "set data layer code store Dir `DIR`.default is `data`.",
			Destination: &tplContext.DataLayerCodeDir,
		},
		&cli.StringFlag{
			Name:        modelOutputFlag,
			Aliases:     []string{"d"},
			Value:       "",
			Usage:       "set model save dir `Dir`.if not set tool working dir will be used.",
			Destination: &tplContext.ModelOutputDir,
		},
		&cli.StringFlag{
			Name:        useCaseFlag,
			Aliases:     []string{"uc"},
			Value:       "",
			Usage:       "set useCase name if you want to generate useCase .if empty,useCase will not generated.",
			Destination: &tplContext.UseCase,
		},
		&cli.StringSliceFlag{
			Name:        modelNamesFlag,
			Aliases:     []string{"n"},
			Value:       nil,
			Usage:       "set model names `ModelName` for mongo code generate.multi model names supported.",
			Required:    true,
			Destination: &tplContext.ModelNames,
		},
	}
}
