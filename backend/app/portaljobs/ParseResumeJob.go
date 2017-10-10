package portaljobs

import (
	"os/exec"

	"github.com/HackGT/SponsorshipPortal/backend/app"
	"github.com/revel/revel"
)

type ParseResume struct {
	ResumeURL     string
	ParticipantID string
	Name          string
	Email         string
}

func (p ParseResume) Run() {
	revel.INFO.Println(p.ResumeURL)
	out, err := exec.Command("ruby", "/home/brow/Documents/textextract.rb", p.ResumeURL).Output()
	if err != nil {
		revel.ERROR.Println(err)
		return
	}
	data := struct {
		Resume string
	}{
		Resume: string(out) + " " + p.Name + " " + p.Email,
	}

	// index some data
	app.BleveIndex.Index(p.ParticipantID, data)
}
