package venc

import (
    "errors"
)

func checkParamBitrate(params *Parameters) error {
    if params.BitControlParams.Bitrate > maxBitrate {
        return errors.New("Bitrate is too large")
    }
    return nil
}

func checkParamStatTime(params *Parameters) error {
    if  params.BitControlParams.StatTime < 1 ||
        params.BitControlParams.StatTime > 60 {
        return errors.New("Stattime should be [1; 60]")
    }
    return nil
}

func checkParamFluctuate(params *Parameters) error {
    if  params.BitControlParams.Fluctuate < 1 ||
        params.BitControlParams.Fluctuate > 5 {
        return errors.New("Fluctuate should be [1; 5]")
    }
    return nil
}

func checkParamQFactor(params *Parameters) error {
    if  params.BitControlParams.QFactor < 1 ||
        params.BitControlParams.QFactor > 99 {
        return errors.New("QFactor should be [1; 99]")
    }
    return nil
}

func checkParamMaxQFactor(params *Parameters) error {
    if  params.BitControlParams.MaxQFactor < 1 ||
        params.BitControlParams.MaxQFactor > 99 {
        return errors.New("MaxQFactor should be [1; 99]")
    }
    return nil
}

func checkParamMinQFactor(params *Parameters) error {
    if  params.BitControlParams.MinQFactor < 1 ||
        params.BitControlParams.MinQFactor > params.BitControlParams.MaxQFactor { //TODO > or >=
        return errors.New("MinQFactor should be [1; MaxQFactor)")
    }
    return nil
}

func checkParamIQp(params *Parameters) error {
    if  params.BitControlParams.IQp < 1 ||
        params.BitControlParams.IQp > 51 {
        return errors.New("IQp should be [1; 51]")
    }
    return nil
}

func checkParamPQp(params *Parameters) error {
    if  params.BitControlParams.PQp < 1 ||
        params.BitControlParams.PQp > 51 {
        return errors.New("PQp should be [1; 51]")
    }
    return nil
}

func checkParamBQp(params *Parameters) error {
    if  params.BitControlParams.BQp < 1 ||
        params.BitControlParams.BQp > 51 {
        return errors.New("BQp should be [1; 51]")
    }
    return nil
}

func checkParamMinQp(params *Parameters) error {
    if  params.BitControlParams.MinQp < 1 ||
        params.BitControlParams.MinQp > 51 {
        return errors.New("MinQp should be [1; 51]")
    }
    return nil
}

func checkParamMaxQp(params *Parameters) error {
    if  params.BitControlParams.MaxQp < params.BitControlParams.MinQp ||
        params.BitControlParams.MaxQp > 51 {
        return errors.New("IQp should be [MinQp; 51]")
    }

    return nil
}

func checkParamMinIQp(params *Parameters) error {
    if  params.BitControlParams.MinIQp < params.BitControlParams.MinQp ||
        params.BitControlParams.MinIQp > params.BitControlParams.MaxQp {
        return errors.New("IQp should be [MinQp; MaxQp]")
    }
    return nil
}

func checkParamIPQpDelta(params *Parameters) error {
    return nil
}

func checkParamSPInterval(params *Parameters) error {
    return nil
}

func checkParamSPQpDelta(params *Parameters) error {
    return nil
}

func checkParamBgInterval(params *Parameters) error {
    return nil
}

func checkParamBgQpDelta(params *Parameters) error {
    return nil
}

func checkParamViQpDelta(params *Parameters) error {
    return nil
}

func checkParamBFrmNum(params *Parameters) error {
    return nil
}

func checkParamBQpDelta(params *Parameters) error {
    return nil
}

