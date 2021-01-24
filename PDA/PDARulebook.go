package PDA

import "reflect"

type PDARulebook struct {
	rules []PDARule
}

func (r PDARulebook) NextConfig(configuration PDAConfiguration, character int32) PDAConfiguration {
	ruleFor := r.ruleFor(configuration, character)
	return ruleFor.Follow(configuration)
}

func (r PDARulebook) ruleFor(configuration PDAConfiguration, character int32) PDARule {
	for _, v := range r.rules {
		if v.AppliesTo(configuration, character) {
			return v
		}
	}

	return PDARule{}
}

func (d PDARulebook) AppliesTo(configuration PDAConfiguration, character int32) bool {
	return !reflect.DeepEqual(d.ruleFor(configuration, character), PDARule{})
}
func (d PDARulebook) followFreeMoves(configuration PDAConfiguration) PDAConfiguration {
	if d.AppliesTo(configuration, -1) {
		return d.followFreeMoves(d.NextConfig(configuration, -1))
	} else {
		return configuration
	}
}
