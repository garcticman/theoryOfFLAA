package PDA

type DPDA struct {
	currentConfig PDAConfiguration
	acceptStates  []int32
	rulebook      PDARulebook
}

func (d DPDA) Accepting() bool {
	for _, v := range d.acceptStates {
		if v == d.GetCurrentConfig().state {
			return true
		}
	}
	return false
}
func (n DPDA) GetCurrentConfig() PDAConfiguration {
	return n.rulebook.followFreeMoves(n.currentConfig)
}
func (d *DPDA) ReadCharacter(character int32) {
	d.currentConfig = d.rulebook.NextConfig(d.GetCurrentConfig(), character)
}
func (d *DPDA) ReadString(string string) {
	for _, v := range string {
		d.ReadCharacter(v)
	}
}
