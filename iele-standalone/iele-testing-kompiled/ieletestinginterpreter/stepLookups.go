package ieletestinginterpreter 

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-standalone/iele-testing-kompiled/ieletestingmodel"
)

func stepLookups(c m.K, config m.K, guard int) (m.K, error) {
	var result m.K
	var err error
	result, err = stepLookupRule768(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule769(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule770(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule771(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule772(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule773(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule774(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule679(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule680(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule681(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule682(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule683(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule684(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule685(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule686(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule687(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule688(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule689(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule690(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule691(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule692(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule693(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule694(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule695(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule696(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule697(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule698(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule699(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule700(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule701(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule702(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule703(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule704(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule705(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule706(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule707(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule708(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule709(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule710(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule711(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule712(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule713(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule714(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule715(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule716(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule717(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule718(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule719(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule720(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule721(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule722(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule723(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule724(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule725(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule726(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule727(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule728(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule729(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule730(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule731(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule732(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule733(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule734(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule735(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule736(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule737(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule738(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule739(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule740(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule741(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule742(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule743(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule744(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule745(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule746(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule747(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule748(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule749(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule750(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule751(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule752(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule753(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule754(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule755(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule756(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule757(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule758(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule759(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule760(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule761(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule762(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule763(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule764(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule765(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule766(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	result, err = stepLookupRule767(c, config, guard)
	if err == nil {
		return result, nil
	}
	if _, isNoStep := err.(*noStepError); !isNoStep {
		return result, err
	}
	return c, noStep
}

