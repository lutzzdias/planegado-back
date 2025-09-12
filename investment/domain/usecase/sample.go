package usecase

import (
	"planegado/investment/domain/entity"
	"strings"
)

type SampleUsecase interface {
	GetSample(value string) entity.Sample
}

type sampleUsecase struct{}

func NewSampleUsecase() SampleUsecase {
	return &sampleUsecase{}
}

func (usecase *sampleUsecase) GetSample(value string) entity.Sample {
	data := strings.ToTitle(value)
	return entity.Sample{Value: data}
}
