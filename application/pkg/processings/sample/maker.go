package sample

import (
    "application/pkg/processings/processing"
)

type sampleMaker string

func Init() {
    var s sampleMaker = "sample"
    processing.Register(s)
}

func (m sampleMaker) Name() string {
    return string(m)
}

func (m sampleMaker) Create() processing.Processing {
    return &sample{}
}

