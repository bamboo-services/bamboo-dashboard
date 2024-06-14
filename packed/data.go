package packed

import "github.com/gogf/gf/v2/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/6TUsU76UBTH8VNo//8yIUlDCCAsEDY7aGLihoPGEBMGTVyRFKwUbaC4M/gAmjgaF5/DR/AJXFxJfARSA6TxnvaeW3rbxTQxn37vJfl1O1nVAB0AHo8/r4B58pCDiTW9n036lunOrh27f3mhgXK0HA26nf86+7+0YkQV88YbO7GUDvMQtUtQ5t7Q9kaW5a7N/eVooAD4fgBrsY1FDuw6s6F9J1FZJ7HUnSUOHbxL/DRVAWf2p1OJwzdjyNRXUBN9wB73hpZEdSsWTd1dEX3iVuaqG2IxdXGB9T1r7Do9T+Zyyzwncd3DGVZ1Rl0b3ZzDiVIyRpZet82zAz/t1V9662jFQMpJVAlvXUD93ZPGrFvwvM+fCCp6caK8IspbcMzIzNGBdRR4qMDWMydKLKHEZ47KWTiaqyLuW8CFFo4+dxOd+yADCRdOlFtDuS8iOzxudHALBS9i0WTJFZR8noUtd43ubaDeN7GYLLaAYn2W5kwanVhGiacqbDNpbBhvv4Iwv/3KgMyaaf82S5WHDxXgS129/QYAAP//XPPhA7AJAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
